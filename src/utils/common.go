package utils

import "github.com/kirsle/configdir"

func GetConfigDir() (string, error) {
  configPath := configdir.LocalConfig("gosora")
  err := configdir.MakePath(configPath)
  if err != nil {
    return "", err
  }
  return configPath, nil
}
