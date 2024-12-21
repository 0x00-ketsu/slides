package cmd

import (
	"fmt"
	"os"

	"github.com/0x00-ketsu/slides/cmd/flags"
	"github.com/0x00-ketsu/slides/config"
	"github.com/0x00-ketsu/slides/internal/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slides",
	Short: "A terminal based preview tool for markdown",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := config.Load(); err != nil {
			fmt.Printf("Error loading config: %+v", err)
			os.Exit(0)
		}

		var filename string
		if len(os.Args) > 1 {
			filename = os.Args[1]
		}
		tagbarConfig := config.LoadTagbar()
		slides := model.Model{
			Filename: filename,
			Page:     0,
			Theme:    flags.Theme,
			TagBar: &model.TagBar{
				Visible: false,
				Width:   tagbarConfig.Width,
			},
		}
		if err := slides.Initial(); err != nil {
			fmt.Printf("Error: %+v", err)
			os.Exit(0)
		}

		if _, err := tea.NewProgram(slides).Run(); err != nil {
			fmt.Printf("Error: %+v", err)
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %+v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.Theme, "theme", "t", "slides", "theme for markdown, options: slides, light, dark")
}
