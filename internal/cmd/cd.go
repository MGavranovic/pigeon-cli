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

func (c *CdCommand) Execute(args []string) (bool, error) {
	if len(args) == 0 {
		return false, fmt.Errorf("please specify the path")
	}

	path := args[0]
	err := os.Chdir(path)
	if err != nil {
		return false, fmt.Errorf("error changing dir using cd: %s", err)
	}
	return false, nil
}
