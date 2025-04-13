package clients

import (
    "fmt"

    "github.com/JeremiahVaughan/healthy/clients/nats"
    "github.com/JeremiahVaughan/healthy/clients/database"
    "github.com/JeremiahVaughan/healthy/config"
)

type Clients struct {
    Nats *nats.Client
    Database *database.Client
}

func New(config config.Config) (*Clients, error) {
    nats, err := nats.New(config.Nats)
    if err != nil {
        return nil, fmt.Errorf("error, when creating new nats for new clients. Error: %v", err)
    }
    db, err := database.New(config.Database)
    if err != nil {
        return nil, fmt.Errorf("error, when creating database client. Error: %v", err)
    }
    return &Clients{
        Nats: nats,
        Database: db,
    }, nil
}
