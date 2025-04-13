package config

import (
    "os"
    "encoding/json"
    "fmt"
    "errors"
)

type Config struct {
    LocalMode bool `json:"localMode"`
    Nats Nats `json:"nats"`
    Database Database 
}

type Nats struct {
    Host string `json:"host"`
    Port int `json:"port"`
}

type Database struct {
    DirectoryLocation string 
}

func New(configPath string, dbLocation string) (Config, error) {
    if dbLocation == "" {
        return Config{}, errors.New("error, must provide a database directory")
    }
    bytes, err := os.ReadFile(configPath)
    if err != nil {
        return Config{}, fmt.Errorf("error, when reading config file. Error: %v", err)
    }
    var c Config
    err = json.Unmarshal(bytes, &c)
    if err != nil {
        return Config{}, fmt.Errorf("error, when decoding config file. Error: %v", err)
    }
    c.Database = Database{ 
        DirectoryLocation: dbLocation,
    }
    return c, nil
}
