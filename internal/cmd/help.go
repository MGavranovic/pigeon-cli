package cmd

import (
	"fmt"
)

type HelpCommand struct {
	Commands map[string]Command
}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Lists all existing commands and it's descriptions"
}

func (c *HelpCommand) Execute(args []string) (bool, error) {
	if len(args) > 0 {
		return false, fmt.Errorf("help needs not arguments")
	}
	fmt.Println("Available commands:")
	for name, cmd := range c.Commands {
		fmt.Printf("%-10s - %s\n", name, cmd.Description())
	}
	return false, nil
}
