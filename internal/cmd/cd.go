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
		return fmt.Errorf("please specify the path")
	}

	path := args[0]
	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("error changing dir using cd: %s", err)
	}
	return nil
}
