package cmd

import "fmt"

type CdCommand struct{}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) Description() string {
	return "Command for navigating the file system"
}

func (c *CdCommand) Execute(args []string) error {
	fmt.Println("This is a cd command")
	return nil
}
