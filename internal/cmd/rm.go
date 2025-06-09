package cmd

import (
	"fmt"
	"os"
)

type RmCommand struct{}

func (c *RmCommand) Name() string {
	return "rm"
}

func (c *RmCommand) Description() string {
	return "Command for moving the files"
}

func (c *RmCommand) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify the file you are trying to delete")
	}

	path := args[0]
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("error removing the file: %s", err)
	}
	fmt.Printf(" - %s - file deleted successfully\n", path)
	return nil
}
