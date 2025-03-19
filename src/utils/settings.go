package utils

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Settings struct {
	TestValue string `toml:"test_value"`
}

func (s *Settings) Load() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, "sora", "settings.toml")

	if _, err := os.Stat(filepath.Dir(configPath)); os.IsNotExist(err) {
		// Create directory if it doesn't exist
		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		s.InitalizeDefaults()
		return s.Save()
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	return toml.Unmarshal(data, s)
}

func (s *Settings) Save() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(configDir, "sora", "settings.toml")

	err = os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		return err
	}

	data, err := toml.Marshal(s)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func (s *Settings) InitalizeDefaults() error {
	s.TestValue = "red"
	return nil
}
