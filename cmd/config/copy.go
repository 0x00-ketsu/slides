package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/0x00-ketsu/slides/config"
	"github.com/0x00-ketsu/slides/internal/utils"
	"github.com/mitchellh/go-homedir"
)

// CopyDefaultConfigFile copies the default configuration file to the user's home directory.
func CopyDefaultConfigFile() error {
	destFilePath, err := homedir.Expand("~/.config/slides/slides.yaml")
	if err != nil {
		return fmt.Errorf("failed to expand the path: %w", err)
	}

	if utils.FileExists(destFilePath) {
		return fmt.Errorf("file already exists: %s", destFilePath)
	}

	// Create the directory if it does not exist
	destDir := filepath.Dir(destFilePath)
	if utils.MkdirIfNotExist(destDir) != nil {
		return fmt.Errorf("failed to create the directory: %s", destDir)
	}

	data := config.DefaultConfig
	if err := os.WriteFile(destFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write the file: %w", err)
	}
	return nil
}
