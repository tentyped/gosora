package main

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tentyped/gosora/src/utils"
)

func main() {
  cliLogger := logrus.New()

  config, err := utils.LoadSettings()
  if err != nil {
    cliLogger.Error("Error loading settings: ", err)
    return
  }

  cliLogger.Println("Initial Config:")
  configJson, _ := json.MarshalIndent(config, "", " ")
  cliLogger.Println(string(configJson))

  newUUID := uuid.New()
  config.SelectedModule = newUUID

  if err := utils.SaveSettings(config); err != nil {
    cliLogger.Error("Error saving config: ", err)
    return
  }

  updatedConfig, err := utils.LoadSettings()
  if err != nil {
    cliLogger.Error("Error loading updated config: ", err)
    return
  }

  cliLogger.Println("\nUpdated Config:")
  updatedConfigJson, _ := json.MarshalIndent(updatedConfig, "", "  ")
  cliLogger.Println(string(updatedConfigJson))
}
