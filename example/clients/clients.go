package clients

import (
    "fmt"

    "github.com/JeremiahVaughan/healthy/example/clients/healthy"
    "github.com/JeremiahVaughan/healthy/example/config"
)

type Clients struct {
    Healthy *healthy.Client
}

func New(config config.Config, serviceName string, healthyRefresh func(status healthy.HealthStatus) error) (*Clients, error) {
    healthy, err := healthy.New(config.Nats, serviceName, healthyRefresh)
    if err != nil {
        return nil, fmt.Errorf("error, when creating new healthy client. Error: %v", err)
    }
    return &Clients{
        Healthy: healthy,
    }, nil
}
