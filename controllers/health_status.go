package controllers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "context"
    "time"

    "github.com/JeremiahVaughan/healthy/models"
    "github.com/JeremiahVaughan/healthy/views"
    "github.com/JeremiahVaughan/healthy/clients/database"

    "github.com/nats-io/nats.go"
)

type HealthStatusController struct {
    view *views.DashBoardView
    model *models.HealthStatusModel
    statusRefreshIntervalInSeconds int64
    RefreshAllKey string
}

func NewHealthStatusController(
    views *views.Views,
    models *models.Models,
    statusRefreshIntervalInSeconds int64,
) *HealthStatusController { 
    return &HealthStatusController{
        view: views.DashBoard,
        model: models.HealthStatus,
        statusRefreshIntervalInSeconds: statusRefreshIntervalInSeconds,
        RefreshAllKey: "refresh-all-health-statuses",
    }
}

func (c *HealthStatusController) Check(w http.ResponseWriter, r *http.Request) {
    healthy, err := c.model.IsHealthy()
    if err != nil {
        err = fmt.Errorf("error, when handling status cake check. Error: %v", err)
        c.model.InternalUnexpectedError(err) 
        return
    }
    if !healthy {
        w.WriteHeader(http.StatusInternalServerError)
    }
}

func (c *HealthStatusController) Dashboard(w http.ResponseWriter, r *http.Request) {
    err := c.view.Render(w)
    if err != nil {
        err = fmt.Errorf("error, when rendering dashboard. Error: %v", err)
        c.model.InternalUnexpectedError(err) 
        return
    }
}

func (c *HealthStatusController) UpdateHealthStatus(msg *nats.Msg) {
    var s database.HealthStatus
    err := json.Unmarshal(msg.Data, &s)
    if err != nil {
        err = fmt.Errorf("error, when decoding the health status for HealthStatusController.UpdateHealthStatus(). Error: %v", err)
        c.model.InternalUnexpectedError(err)
        return
    }
    err = c.model.UpdateHealthStatus(s)
}

func (c *HealthStatusController) ReportUnexpectedError(msg *nats.Msg) {
    var s database.HealthStatus
    err := json.Unmarshal(msg.Data, &s)
    if err != nil {
        err = fmt.Errorf("error, when decoding the health status for HealthStatusController.ReportUnexpectedError(). Error: %v", err)
        c.model.InternalUnexpectedError(err)
        return
    }
    c.model.ExternalUnexpectedError(s)
}

func (c *HealthStatusController) RefreshAll(ctx context.Context) {
    for {
        select {
        case <- time.After(time.Second * time.Duration(c.statusRefreshIntervalInSeconds)):
            statuses, err := c.model.GetAllStatuses()
            if err != nil {
                err = fmt.Errorf("error, when GetAllStatuses() for HealthStatusController.RefreshAll(). Error: %v", err)
                c.model.InternalUnexpectedError(err)
                continue
            }
            for _, s := range statuses {
                bytes, err := json.Marshal(s)
                if err != nil {
                    err = fmt.Errorf("error, when marshaling status for HealthStatusController.RefreshAll(). Error: %v", err)
                    c.model.InternalUnexpectedError(err)
                    continue
                }
                err = c.model.RefreshStatus(c.RefreshAllKey, bytes)
                if err != nil {
                    err = fmt.Errorf("error, when RefreshStatus() for HealthStatusController.RefreshAll(). Error: %v", err)
                    c.model.InternalUnexpectedError(err)
                    continue
                }
            }
        case <- ctx.Done():
        }
    }
}
