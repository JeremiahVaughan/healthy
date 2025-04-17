package main

import (
    "flag"
    "log"

    "github.com/JeremiahVaughan/healthy/config"
    "github.com/JeremiahVaughan/healthy/clients"
    "github.com/JeremiahVaughan/healthy/views"
    "github.com/JeremiahVaughan/healthy/models"
    "github.com/JeremiahVaughan/healthy/controllers"
    "github.com/JeremiahVaughan/healthy/router"
)

func main() {
    log.Println("healthy is starting")
    var configPath string
    flag.StringVar(
        &configPath,
        "c",
        "config/config.json",
        "location of config file",
    ) 

    flag.Parse()

    config, err := config.New(configPath)
    if err != nil {
        log.Fatalf("error, when creating new config for main(). Error: %v", err)
    }

    clients, err := clients.New(config)
    if err != nil {
        log.Fatalf("error, when creating clients for main(). Error: %v", err)
    }

    views, err := views.New(config.LocalMode, config.UiPath)
    if err != nil {
        log.Fatalf("error, when creating views for main(). Error: %v", err)
    }
    models := models.New(clients)
    controllers := controllers.New(views, models)
    router := router.New(controllers)

    log.Println("healthy is running")
    err = router.Run()
    if err != nil {
        log.Fatalf("error, when starting router. Error: %v", err)
    }
}

