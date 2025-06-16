package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type ExitCommand struct{}

func (c *ExitCommand) Name() string {
	return "exit"
}

func (c *ExitCommand) Description() string {
	return "Command for closing the program"
}

func (c *ExitCommand) Execute(args []string) error {
	color.RGB(255, 128, 0).Println("See you soon!")
	color.Red("See you soon!")
	fmt.Println("See you soon!")

	color.Red("This is red")
	color.Green("This is green")
	color.Yellow("This is yellow")

	bold := color.New(color.FgCyan).Add(color.Bold)
	bold.Println("This is bold cyan")

	os.Exit(0)
	return nil
}
