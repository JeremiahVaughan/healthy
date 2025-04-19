package views

import (
    "fmt"

    "github.com/JeremiahVaughan/healthy/ui_util"
    "github.com/JeremiahVaughan/healthy/models"
)

type Views struct {
    InfraGraph *InfraGraphView
    TemplateLoader *ui_util.TemplateLoader
    DashBoard *DashBoardView
}

func New(localMode bool, uiPath string, models *models.Models) (*Views, error) { 
    templates := []ui_util.HtmlTemplate{
        {
            Name: "dash-board",
            FileOverrides: []string{
                "dash_board.html",
            },
        },
        {
            Name: "infra-graph",
            FileOverrides: []string{
                "infra_graph.html",
            },
        },
    }
    tl, err := ui_util.NewTemplateLoader(
        uiPath + "/templates/base",
        uiPath + "/templates/overrides",
        templates,
        localMode,
    )
    if err != nil {
        return nil, fmt.Errorf("error, when creating template loader. Error: %v", err)
    }
    return &Views{
        InfraGraph: NewInfraGraphView(tl, localMode),
        DashBoard: NewDashBoardView(tl, localMode, models),
        TemplateLoader: tl,
    }, nil
}
