package views

import (
    "net/http"
    "fmt"

    "github.com/JeremiahVaughan/healthy/ui_util"
)

type InfraGraphView struct {
    tl *ui_util.TemplateLoader
    LocalMode bool
}

func NewInfraGraphView(tl *ui_util.TemplateLoader, localMode bool) *InfraGraphView {
    return &InfraGraphView{
        tl: tl,
        LocalMode: localMode,
    }
}

func (i *InfraGraphView) Render(w http.ResponseWriter) error {
    err := i.tl.GetTemplateGroup("base").ExecuteTemplate(w, "base", i)
    if err != nil {
        return fmt.Errorf("error, when rendering template for InfraGraphView.Render(). Error: %v", err)
    }

    return nil
}
