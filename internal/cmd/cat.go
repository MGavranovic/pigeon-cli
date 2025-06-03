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

func (c *CatCommand) Execute(args []string) error {
	fmt.Printf("Args received: %s\n", args)
	if len(args) == 0 {
		fmt.Println("Pleace specify the file you want to read")
	}

	fName := args[0]
	file, err := os.ReadFile(fName)
	if err != nil {
		fmt.Printf("Error reading from file %s: %s\n", file, err)
	}
	s := string(file)
	fmt.Println(s)
	return nil
}
