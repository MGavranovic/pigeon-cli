package cmd

import (
	"fmt"
	"os"
)

type TouchCommand struct{}

func (c *TouchCommand) Name() string {
	return "touch"
}

func (c *TouchCommand) Description() string {
	return "command for creating empty files"
}

func (c *TouchCommand) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please specify the 1 filename you are trying to create")
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting cur working dir")
	}

	files, err := os.ReadDir(wd)
	if err != nil {
		return fmt.Errorf("error reading dir")
	}
	for _, f := range files {
		if args[0] == f.Name() {
			return fmt.Errorf("file already exists")
		}
	}

	file, err := os.OpenFile(args[0], os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("issue creating a file: %s", err)
	}
	defer file.Close()

	return nil
}
