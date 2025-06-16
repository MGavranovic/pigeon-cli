package cmd

import (
	"fmt"
	"os"
)

type RnCommand struct{}

func (c *RnCommand) Name() string {
	return "rn"
}

func (c *RnCommand) Description() string {
	return "Command for renaming files or dirs"
}

func (c *RnCommand) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("please specify the file/dir to rename, and it's new name")
	}
	old := args[0]
	new := args[1]
	err := os.Rename(old, new)
	if err != nil {
		return fmt.Errorf("unable to rename '%s' into '%s'", old, new)
	} else {
		fmt.Printf("Successfully renamed '%s' into '%s'", old, new)
	}
	return nil
}
