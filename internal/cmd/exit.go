package cmd

import (
	"fmt"
	"os"
)

type ExitCommand struct{}

func (c *ExitCommand) Name() string {
	return "exit"
}

func (c *ExitCommand) Description() string {
	return "Command for closing the program"
}

func (c *ExitCommand) Execute(args []string) error {
	fmt.Println("See you soon!")
	os.Exit(0)
	return nil
}
