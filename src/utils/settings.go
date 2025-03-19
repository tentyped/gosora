package utils

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
)

type Settings struct {
  SelectedModule uuid.UUID `toml:"selected_module"`
}

func LoadSettings() (Settings, error) {
  configDir, err := GetConfigDir()
  if err != nil {
    return Settings{}, err
  }

  file := filepath.Join(configDir, "settings.toml")

  if _, err := os.Stat(configDir); os.IsNotExist(err) {
    if err := os.MkdirAll(configDir, 0755); err != nil {
      return Settings{}, err
    }
  }

  if _, err := os.Stat(file); os.IsNotExist(err) {
    def := Settings {
      SelectedModule: uuid.Nil,
    }
    if err := SaveSettings(def); err != nil {
      return Settings{}, err
    }
  }

  data, err := os.ReadFile(file)
  if err != nil {
    return Settings{}, err
  }

  var settings Settings
  if err := toml.Unmarshal(data, &settings); err != nil {
    return Settings{}, err
  }

  return settings, nil
}

func SaveSettings(settings Settings) error {
  configDir, err := GetConfigDir()
  if err != nil {
    return err
  }

  file := filepath.Join(configDir, "settings.toml")

  data, err := toml.Marshal(settings)
  if err != nil {
    return err
  }

  return os.WriteFile(file, data, 0644)
}
