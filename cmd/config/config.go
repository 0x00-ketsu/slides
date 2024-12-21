package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ConfigCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Run: func(cmd *cobra.Command, args []string) {
		isCopyConfigFile, err := cmd.Flags().GetBool("copy")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if isCopyConfigFile {
			if err := CopyDefaultConfigFile(); err == nil {
				fmt.Println("Copy the default config file to the local config directory successfully.")
			} else {
				fmt.Printf("Error: %v\n", err)
			}
		}
	},
}

func init() {
	ConfigCmd.Flags().BoolP("copy", "c", false, "Copy the default config file to the local config directory (~/.config/slides/slides.yaml)")
}
