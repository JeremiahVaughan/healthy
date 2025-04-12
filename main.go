package main

import (
    "flag"
    "log"

    "github.com/JeremiahVaughan/healthy/config"
)

func main() {
    log.Println("healthy is starting")
    var configPath string
    flag.StringVar(
        &configPath,
        "configPath",
        "config/config.json",
        "location of config file",
    ) 

    _, err := config.New(configPath)
    if err != nil {
        log.Fatalf("error, when creating new config in main(). Error: %v", err)
    }


    log.Println("healthy is running")
}

