package views

import (
    "net/http"
    "fmt"

    "github.com/JeremiahVaughan/healthy/ui_util"
    "github.com/JeremiahVaughan/healthy/models"
    "github.com/JeremiahVaughan/healthy/clients/database"
)

type DashBoardView struct {
    tl *ui_util.TemplateLoader
    localMode bool
    model *models.HealthStatusModel
}

func NewDashBoardView(
    tl *ui_util.TemplateLoader,
    localMode bool,
    models *models.Models,
) *DashBoardView {
    return &DashBoardView{
        tl: tl,
        localMode: localMode,
        model: models.HealthStatus,
    }
}

type DashBoard struct {
    LocalMode bool
    TableHeaders []string
    TableRows []database.HealthStatus
}

func (i *DashBoardView) Render(w http.ResponseWriter) error {
    s, err := i.model.GetAllStatuses()
    if err != nil {
        return fmt.Errorf("error, when GetAllStatuses() for DashBoardView.Render(). Error: %v", err)
    }
    d := DashBoard{
        LocalMode: i.localMode,
        TableHeaders: []string{
            "Service",
            "Status Key",
            "Status",
            "Unhealthy Started At",
            "Unhealthy Delay In Seconds",
            "Message",
            "Expires At",
        },
        TableRows: s,
    }
    err = i.tl.GetTemplateGroup("dash-board").ExecuteTemplate(w, "base", d)
    if err != nil {
        return fmt.Errorf("error, when rendering template for DashBoardView.Render(). Error: %v", err)
    }
    return nil
}
