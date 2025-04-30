package views

import (
    "time"
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

func (v *DashBoardView) Render(w http.ResponseWriter) error {
    s, err := v.model.GetAllStatuses()
    if err != nil {
        return fmt.Errorf("error, when GetAllStatuses() for DashBoardView.Render(). Error: %v", err)
    }
    for i, status := range s {
        status.UnhealthyStartedAtDisplay = v.FormatTime(status.UnhealthyStartedAt)
        status.ExpiresAtDisplay = v.FormatTime(status.ExpiresAt)
        s[i] = status
    }
    d := DashBoard{
        LocalMode: v.localMode,
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
    err = v.tl.GetTemplateGroup("dash-board").ExecuteTemplate(w, "base", d)
    if err != nil {
        return fmt.Errorf("error, when rendering template for DashBoardView.Render(). Error: %v", err)
    }
    return nil
}

func (v *DashBoardView) FormatTime(theTime int64) string {
    if theTime == 0 {
        return "N/A"
    }
    return time.Unix(theTime, 0).Local().Format("2 Jan 2006 15:04:05 CDT")
}
