package models

import (
    "github.com/JeremiahVaughan/healthy/clients"
)

type Models struct {
    UnexpectedErrors *UnexpectedErrorsModel
    // UpdateHealthStatus *UpdateHealthStatusModel
    // CheckHealthStatus *CheckHealthStatusModel
}

func New(clients *clients.Clients) *Models {
    // todo implement
    return &Models{
        UnexpectedErrors: NewUnexpectedErrorsModel(),
    }
}
