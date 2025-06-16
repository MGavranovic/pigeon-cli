package cmd

import (
	"fmt"
	// "os"
)

type MvCommand struct{}

func (c *MvCommand) Name() string {
	return "mv"
}

func (c *MvCommand) Description() string {
	return "Command for moving the files"
}

func (c *MvCommand) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("pleace specify the file you want to move and the destination dir")
	}

	return nil
}
