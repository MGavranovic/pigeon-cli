package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type RnCommand struct{}

func (c *RnCommand) Name() string {
	return "rn"
}

func (c *RnCommand) Description() string {
	return "Command for renaming files or dirs"
}

func (c *RnCommand) Execute(args []string) (bool, error) {
	if len(args) < 2 {
		return false, fmt.Errorf("please specify the file/dir to rename, and it's new name")
	}
	old := args[0]
	new := args[1]
	err := os.Rename(old, new)
	if err != nil {
		return false, fmt.Errorf("unable to rename '%s' into '%s'", old, new)
	} else {
		styleOld := color.New(color.FgRed).SprintfFunc()
		styleNew := color.New(color.FgGreen).SprintfFunc()
		fmt.Printf("Successfully renamed '%s' into '%s'\n", styleOld(old), styleNew(new))
	}
	return false, nil
}
