package views

import (
    "fmt"

    "github.com/JeremiahVaughan/healthy/ui_util"
)

type Views struct {
    InfraGraph *InfraGraphView
    TemplateLoader *ui_util.TemplateLoader
}

func New(localMode bool, uiPath string) (*Views, error) { 
    templates := []ui_util.HtmlTemplate{
        {
            Name: "base",
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
        TemplateLoader: tl,
    }, nil
}
