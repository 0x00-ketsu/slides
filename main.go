package main

import (
	"fmt"
	"os"

	"github.com/0x00-ketsu/smooth/config"
	"github.com/0x00-ketsu/smooth/term/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Load config file
	v := config.Load("")
	if v == nil {
		fmt.Println("Error: load config file failed")
		os.Exit(0)
	}

	var conf config.Config
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Printf("Error: Viper unmarshal failed, %v", err)
		os.Exit(0)
	}

	// Read file
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	smooth := model.Model{
		Filename: filename,
		Page:     0,
		Config:   conf,
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
}
