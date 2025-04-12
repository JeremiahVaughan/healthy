package config

import (
    "os"
    "encoding/json"
    "fmt"
)

type Config struct {
    Nats Nats `json:"nats"`
}

type Nats struct {
    Host string `json:"host"`
    Port int `json:"port"`
}

func New(configPath string) (Config, error) {
    bytes, err := os.ReadFile(configPath)
    if err != nil {
        return Config{}, fmt.Errorf("error, when reading config file. Error: %v", err)
    }
    var c Config
    err = json.Unmarshal(bytes, &c)
    if err != nil {
        return Config{}, fmt.Errorf("error, when decoding config file. Error: %v", err)
    }
    return c, nil
}
