package main

import (
  "encoding/json"
  "log"
  "os"
)

var globalConfiguration *configuration = new(configuration)

type configuration struct {
  Database database
  Port     string
}

func init() {
  loadConfiguration()
}

func main() {
  startServer(globalConfiguration.Port)
}

func loadConfiguration() {
  configFilePath := "configs/main.conf"

  log.Printf("starting %s load\n", configFilePath)
  configFile, err := os.Open(configFilePath)
  if err != nil {
    log.Println("[ERROR] ", err)
    log.Println("For your happiness an example config file is provided in the 'conf' directory in the repository.")
    os.Exit(1)
  }

  configDecoder := json.NewDecoder(configFile)
  err = configDecoder.Decode(&globalConfiguration)
  if err != nil {
    log.Println("[ERROR] ", err)
    log.Println("Please ensure that your config file is in valid JSON format.")
    os.Exit(1)
  }

  log.Printf("finished %s load\n", configFilePath)
}
