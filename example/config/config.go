package config

type Config struct {
    Nats Nats `json:"nats"`
}

// config struct for nats
type Nats struct {
    Host string `json:"host"`
    Port int `json:"port"`
}


