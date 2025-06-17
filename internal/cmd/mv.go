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
		return fmt.Errorf("please specify the file you want to move and the destination dir")
	}
	if len(args) > 2 {
		return fmt.Errorf("too many arguments listed, please specify the file you want to move, and the destination dir")
	}
	// 1. make a note of the original/current wd
	curwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting the current working dir: %s", err)
	}

	// 2. open the file being moved
	file, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	// 3. change dir to the file destinationo
	if err := os.Chdir(args[1]); err != nil {
		file.Close()
		return fmt.Errorf("error changing to '%s' dir: %s", args[1], err)
	}

	//4. create the destination file
	newFile, err := os.Create(file.Name())
	if err != nil {
		file.Close()
		return fmt.Errorf("error creating file: %s", err)
	}
	defer newFile.Close()

	// 5. copy the contents
	_, errCopy := io.Copy(newFile, file)
	if errCopy != nil {
		file.Close()
		return fmt.Errorf("error copying contents of file: %s", errCopy)
	}

	// 6. close the original file
	if err := file.Close(); err != nil {
		return fmt.Errorf("unable to close the file: %s", err)
	}

	// 7. go back to the original wd
	if err := os.Chdir(curwd); err != nil {
		return fmt.Errorf("error going back to the original dir: %s", err)
	}

	// 8. remove the original file
	if err := os.Remove(args[0]); err != nil {
		return fmt.Errorf("unable to remove the original file: %s", err)
	}
	return nil
}
