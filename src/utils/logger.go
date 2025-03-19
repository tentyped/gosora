package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
		DisableColors:   false,
		ForceColors:     true,
	})

	// Set log level
	logger.SetLevel(logrus.InfoLevel)

	return logger
}

func ClearLogs() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config directory: %w", err)
	}

	logPath := filepath.Join(configDir, "sora", "logs.txt")
	return os.Truncate(logPath, 0)
}
