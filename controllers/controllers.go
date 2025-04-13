package controllers

import (
    "github.com/JeremiahVaughan/healthy/views"
    "github.com/JeremiahVaughan/healthy/models"
    "github.com/JeremiahVaughan/healthy/ui_util"
)

type Controllers struct {
    InfraGraph *InfraGraphController
    TemplateLoader *ui_util.TemplateLoader
}

func New(views *views.Views, models *models.Models) *Controllers { 
    return &Controllers{
        InfraGraph: NewInfraGraphController(views, models.UnexpectedErrors),
        TemplateLoader: views.TemplateLoader,
    }
}
