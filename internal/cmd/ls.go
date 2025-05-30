package cmd

import (
	"fmt"
	"os"
)

type LsCommand struct{}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Description() string {
	return "Lists all the items in the current dir"
}

func (c *LsCommand) Execute(args []string) error {
	for _, arg := range args {
		fmt.Printf("These are args for the ls command %s\n", arg)
	}

	dir, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, f := range dir {
		fmt.Println(f.Name())
	}
	return nil
}
