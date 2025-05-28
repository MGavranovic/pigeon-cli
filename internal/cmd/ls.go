package cmd

import (
	"fmt"
	"os"
)

type LsCommand struct{}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Execute(args []string) error {
	dir, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, f := range dir {
		fmt.Println(f.Name())
	}
	return nil
}
