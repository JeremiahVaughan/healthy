package controllers

import (
    "fmt"
    "net/http"

    "github.com/JeremiahVaughan/healthy/views"
    "github.com/JeremiahVaughan/healthy/models"
)

type InfraGraphController struct {
    view *views.InfraGraphView
    errModel *models.HealthStatusModel
}

func NewInfraGraphController(views *views.Views, errModel *models.HealthStatusModel) *InfraGraphController { 
    return &InfraGraphController{
        view: views.InfraGraph,
        errModel: errModel,
    }
}

func (c *InfraGraphController) Graph(w http.ResponseWriter, r *http.Request) {
    err := c.view.Render(w)
    if err != nil {
        err = fmt.Errorf("error, when rendering view for infra graph. Error: %v", err)
        c.errModel.InternalUnexpectedError(err) 
        return
    }
}
