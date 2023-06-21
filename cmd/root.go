package cmd

import (
	"fmt"
	"os"

	"github.com/0x00-ketsu/slides/cmd/flags"
	"github.com/0x00-ketsu/slides/term/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "smooth",
	Short: "A terminal based preview tool for markdown",
	Run: func(cmd *cobra.Command, args []string) {
		// Read file
		var filename string
		if len(os.Args) > 1 {
			filename = os.Args[1]
		}

		smooth := model.Model{
			Filename: filename,
			Page:     0,
			Theme: flags.Theme,
		}
		if err := smooth.Initial(); err != nil {
			fmt.Printf("Error: %v", err.Error())
			os.Exit(0)
		}

		p := tea.NewProgram(smooth)
		if err := p.Start(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(0)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.Theme, "theme", "t", "smooth", "theme for markdown, choices: smooth|ascii|light|dark|notty")
}
