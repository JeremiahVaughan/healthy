package controllers

import (
    "github.com/JeremiahVaughan/healthy/views"
    "github.com/JeremiahVaughan/healthy/models"
    "github.com/JeremiahVaughan/healthy/ui_util"
)

type Controllers struct {
    InfraGraph *InfraGraphController
    TemplateLoader *ui_util.TemplateLoader
    Health *HealthController
    HealthStatus *HealthStatusController
}

func New(views *views.Views, models *models.Models, statusRefreshIntervalInSeconds int64) *Controllers { 
    return &Controllers{
        InfraGraph: NewInfraGraphController(views, models.HealthStatus),
        TemplateLoader: views.TemplateLoader,
        Health: NewHealthController(),
        HealthStatus: NewHealthStatusController(views, models, statusRefreshIntervalInSeconds),
    }
}
