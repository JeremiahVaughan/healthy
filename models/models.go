package models

import (
    "github.com/JeremiahVaughan/healthy/clients"
    "github.com/JeremiahVaughan/healthy/config"
)

type Models struct {
    HealthStatus *HealthStatusModel
}

func New(clients *clients.Clients, config config.Config) *Models {
    return &Models{
        HealthStatus: NewHealthStatusModel(clients, config),
    }
}
