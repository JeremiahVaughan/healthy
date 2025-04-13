package nats

import (
    "fmt"

    "github.com/JeremiahVaughan/healthy/config"

    "github.com/nats-io/nats.go"
)

type Client struct {
    conn *nats.Conn
}

func New(config config.Nats) (*Client, error) {
    url := fmt.Sprintf("%s:%d", config.Host, config.Port)
    var err error
    var result Client
    result.conn, err = nats.Connect(url)
    if err != nil {
        return nil, fmt.Errorf("error, when connecting to nats service for client init. Error: %v", err)
    }
    return &result, nil
}
