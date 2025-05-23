package config

import (
    "encoding/json"
    "context"
    "fmt"
    "errors"
)

type Config struct {
    LocalMode bool `json:"localMode"`
    Nats Nats `json:"nats"`
    UiPath string `json:"uiPath"`
    Database Database `json:"database"`
    HealthStatusExpiresDurationInSeconds int64 `json:"healthStatusExpiresDurationInSeconds"`
    StatusRefreshIntervalInSeconds int64 `json:"statusRefreshIntervalInSeconds"`
}

type Nats struct {
    Host string `json:"host"`
    Port int `json:"port"`
}

type Database struct {
    DataDirectory string `json:"dataDirectory"`
    MigrationDirectory string `json:"migrationDirectory"`
}

func New(ctx context.Context) (Config, error) {
    bytes, err := fetchConfigFromS3(ctx, "healthy")
    if err != nil {
        return Config{}, fmt.Errorf("error, when fetching config file. Error: %v", err)
    }

    var c Config
    err = json.Unmarshal(bytes, &c)
    if err != nil {
        return Config{}, fmt.Errorf("error, when decoding config file. Error: %v", err)
    }

    err = c.validate()
    if err != nil {
        return Config{}, fmt.Errorf("error, configuration validation failed. Error: %v", err)
    }
    return c, nil
}


func (c *Config) validate() error {                                   
   if c.Nats.Host == "" {                                            
       return errors.New("NATS host must not be empty")              
   }                                                                 
                                                                     
   if c.Nats.Port < 1 || c.Nats.Port > 65535 {                       
       return errors.New("NATS port must be between 1 and 65535")    
   }                                                                 

   if c.UiPath == "" {
       return errors.New("UI path required") 
   }
   
   if c.Database.DataDirectory == "" {                               
       return errors.New("Database dataDirectory must not be empty") 
   }                                                                 
                                                                     
   if c.Database.MigrationDirectory == "" {                          
       return errors.New("Database migrationDirectory must not be empty")
   }                                                                 

   if c.HealthStatusExpiresDurationInSeconds == 0 {
       return errors.New("must provide a healthStatusExpiresDurationInSeconds value")
   }
                                                                     
   return nil                                                        
}                                                                     
