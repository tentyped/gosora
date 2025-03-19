package sora

import (
	"os"
	"path/filepath"
	"sora/src/modules"
	"sora/src/utils"

	"go.uber.org/zap"
)

func InitializationStageFunction() (*zap.Logger, *utils.Settings, string, *modules.ModuleManager, error) {
  // Ensure Sora config directory exists
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, nil, "", nil, err
	}
	configDir = filepath.Join(configDir, "sora")
	os.MkdirAll(configDir, 0755) // Ensure directory exists

	// Create logger with a file writer
	logFile := filepath.Join(configDir, "logs.txt")
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{logFile, "stdout"} // Log to file & stdout
	logger, err := cfg.Build()
	if err != nil {
		return nil, nil, "", nil, err
	}

	logger.Info("Logger initialized successfully")

	// Load settings
	settings := &utils.Settings{}
	err = settings.Load()
	if err != nil {
		logger.Error("Failed to load settings", zap.Error(err))
		return nil, nil, "", nil, err
	}
	logger.Info("Settings loaded successfully")

	// Initialize ModuleManager
	moduleManager := modules.NewModuleManager(configDir, logger)
	logger.Info("Module manager initialized successfully")

	return logger, settings, configDir, moduleManager, nil
}
