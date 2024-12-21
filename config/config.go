package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/0x00-ketsu/slides/config/structures"
	"github.com/0x00-ketsu/slides/internal/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//go:embed default.yaml
var DefaultConfig []byte

// Config is the main configuration for the slides
type Config struct {
	// Tagbar is the configuration for the tagbar.
	Tagbar structures.Tagbar `mapstructure:"tagbar"`

	// Keymaps is the configuration for the key maps.
	Keymaps structures.Keymap `mapstructure:"keymaps"`
}

// Load loads the configuration.
func Load() (Config, error) {
	data, err := loadConfigData()
	if err != nil {
		return Config{}, fmt.Errorf("failed to load config data: %w", err)
	}
	return parseConfig(data)
}

// LoadTagbar loads the tagbar configuration.
func LoadTagbar() structures.Tagbar {
	config, err := Load()
	if err != nil {
		return structures.Tagbar{}
	}
	return config.Tagbar
}

// LoadKeymaps loads the keymaps configuration.
func LoadKeymaps() structures.Keymap {
	config, err := Load()
	if err != nil {
		return structures.Keymap{}
	}
	return config.Keymaps
}

// Parse config data into the Config structure.
func parseConfig(data []byte) (Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	var config Config
	if err := v.ReadConfig(bytes.NewReader(data)); err != nil {
		return config, fmt.Errorf("failed to read config: %w", err)
	}
	if err := v.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return config, nil
}

// Load config data from the OS local config directory.
//
// If the local config file does not exist, it will return the default config.
func loadConfigData() ([]byte, error) {
	path, err := homedir.Expand("~/.config/slides")
	if err != nil {
		return nil, fmt.Errorf("failed to expand home directory: %w", err)
	}

	filePath := filepath.Join(path, "slides.yaml")
	if !utils.FileExists(filePath) {
		return DefaultConfig, nil
	}
	return os.ReadFile(filePath)
}
