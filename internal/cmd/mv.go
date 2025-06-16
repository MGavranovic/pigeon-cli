package cmd

import (
	"fmt"
	"io"
	"os"
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

	file, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("error opening %s", file.Name())
	}
	defer file.Close()

	dest, err := os.Create(args[1])
	if err != nil {
		return fmt.Errorf("errror creating %s", dest.Name())
	}
	defer dest.Close()

	_, errCopy := io.Copy(dest, file)
	if errCopy != nil {
		return fmt.Errorf("error copying contents of %s", file.Name())
	}

	if err := os.Remove(file.Name()); err != nil {
		return fmt.Errorf("unable to move %s to %s", file.Name(), dest.Name())
	}
	return nil
}
