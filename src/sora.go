package sora

import (
	"os"
	"path/filepath"
	"sora/src/modules"
	"sora/src/utils"

	"github.com/sirupsen/logrus"
)

func InitializationStageFunction() (*logrus.Logger, *utils.Settings, string, *modules.ModuleManager, error) {
  // Ensure Sora config directory exists
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, nil, "", nil, err
	}
	configDir = filepath.Join(configDir, "sora")
	os.MkdirAll(configDir, 0755) // Ensure directory exists

	// Create logger with a file writer
  logger := utils.InitLogger()

	logger.Info("Logger initialized successfully")

	// Load settings
	settings := &utils.Settings{}
	err = settings.Load()
	if err != nil {
		logger.Error("Failed to load settings", err)
		return nil, nil, "", nil, err
	}
	logger.Info("Settings loaded successfully")

	// Initialize ModuleManager
	moduleManager := modules.NewModuleManager(configDir, logger)
	logger.Info("Module manager initialized successfully")

	return logger, settings, configDir, moduleManager, nil
}
