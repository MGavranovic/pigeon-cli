package cmd

import (
	// "fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type ExitCommand struct{}

func (c *ExitCommand) Name() string {
	return "exit"
}

func (c *ExitCommand) Description() string {
	return "Command for closing the program"
}

func (c *ExitCommand) Execute(args []string) (bool, error) {
	style := color.RGB(35, 82, 38)
	style.Add(color.BgGreen)

	style.Println("See you soon! 👋")
	time.Sleep(10 * time.Second)
	os.Exit(0)
	return false, nil
}
