package controllers

import (
    "fmt"
    "net/http"

    "github.com/JeremiahVaughan/healthy/views"
    "github.com/JeremiahVaughan/healthy/models"
)

type InfraGraphController struct {
    view *views.InfraGraphView
    errModel *models.UnexpectedErrorsModel
}

func NewInfraGraphController(views *views.Views, errModel *models.UnexpectedErrorsModel) *InfraGraphController { 
    return &InfraGraphController{
        view: views.InfraGraph,
        errModel: errModel,
    }
}

func (i *InfraGraphController) Graph(w http.ResponseWriter, r *http.Request) {
    err := i.view.Render(w)
    if err != nil {
        err = fmt.Errorf("error, when rendering view for infra graph. Error: %v", err)
        i.errModel.Internal(err) 
        return
    }
}
