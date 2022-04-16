package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Theme string `mapstructure:"theme" json:"theme"`
}

// Load loads config file
// @param filePath config file path
func Load(filePath string) *viper.Viper {
	v := viper.New()

	if filePath == "" {
		filePath = getDefaultFile()
	}

	v.SetConfigType("yaml")
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error: read config file failed, %v", err)
		return nil
	}

	return v
}

// get default config file and return it's file path
func getDefaultFile() string {
	return "config.yaml"
}
