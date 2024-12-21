package cmd

import "github.com/0x00-ketsu/slides/cmd/config"

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
}
