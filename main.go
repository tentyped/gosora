package main

import (
	"fmt"
	"os"
	"sora/src"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func main() {
  logger, _, configDir, moduleManager, err := sora.InitializationStageFunction()
  if err != nil {
    fmt.Println("Failed to initialize Sora", err)
    os.Exit(1)
  }
  defer logger.Sync()

  if len(os.Args) < 2 {
    fmt.Println("Usage: sora <command> [args]")
    os.Exit(1)
  }

  command := os.Args[1]

  switch command {
	case "add_module":
		if len(os.Args) < 3 {
			fmt.Println("Usage: sora add_module <metadata_url>")
			os.Exit(1)
		}
		metadataURL := os.Args[2]
		module, err := moduleManager.AddModule(metadataURL, configDir)
		if err != nil {
			logger.Error("Failed to add module", zap.String("URL", metadataURL), zap.Error(err))
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		logger.Info("Module added", zap.String("Name", module.Metadata.SourceName))
		fmt.Println("Added module:", module.Metadata.SourceName)

	case "delete_module":
		if len(os.Args) < 3 {
			fmt.Println("Usage: sora delete_module <module_id>")
			os.Exit(1)
		}
		moduleID, err := uuid.Parse(os.Args[2])
		if err != nil {
			fmt.Println("Invalid module ID format")
			os.Exit(1)
		}
		err = moduleManager.DeleteModule(moduleID, configDir)
		if err != nil {
			logger.Error("Failed to delete module", zap.String("ID", moduleID.String()), zap.Error(err))
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		logger.Info("Module deleted", zap.String("ID", moduleID.String()))
		fmt.Println("Deleted module:", moduleID.String())

	case "refresh_modules":
		moduleManager.RefreshModules(configDir)
		logger.Info("Modules refreshed")
		fmt.Println("Modules refreshed.")

  case "get_modules":
    modules := moduleManager.GetModules()
    for _, mod := range modules {
      fmt.Printf("ID: %s, Name: %s\n", mod["id"], mod["name"])
    }

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add_module, delete_module, refresh_modules")
		os.Exit(1)
	}
}
