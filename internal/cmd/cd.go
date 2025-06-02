package cmd

import (
	"fmt"
	"os"
)

type CdCommand struct{}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) Description() string {
	return "Command for navigating the file system"
}

func (c *CdCommand) Execute(args []string) error {
	if len(args) == 0 {
		fmt.Printf("Pleace specify the path\n")
		return nil
	}

	path := args[0]
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("Error changing dir using cd: %s\n", err)
	}
	return nil
}
