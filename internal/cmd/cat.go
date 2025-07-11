package cmd

import (
	"fmt"
	"os"
)

type CatCommand struct{}

func (c *CatCommand) Name() string {
	return "cat"
}

func (c *CatCommand) Description() string {
	return "Reads the contents of a file"
}

func (c *CatCommand) Execute(args []string) (bool, error) {
	fmt.Printf("Args received: %s\n", args)
	if len(args) == 0 {
		return false, fmt.Errorf("please specify the file you want to read")
	}

	fName := args[0]
	file, err := os.ReadFile(fName)
	if err != nil {
		return false, fmt.Errorf("error reading from file %s: %s", fName, err)
	}
	s := string(file)
	fmt.Println(s)
	return false, nil
}
