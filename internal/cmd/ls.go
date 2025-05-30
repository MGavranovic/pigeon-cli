package cmd

import (
	"fmt"
	"os"
	// "slices"
)

type LsCommand struct{}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Description() string {
	return "Lists all the items in the current dir"
}

func (c *LsCommand) Execute(args []string) error {
	// allowedFlags := []string{"-a"}
	// for i := 0; i < len(args); i++ {
	// 	if slices.Contains(allowedFlags, args[i]) {
	// 		fmt.Printf("Contains the flag: %s\n", args[i])
	// 	} else {
	// 		fmt.Printf("ls command doesn't recognize the flag %s\n", args[i])
	// 	}
	// }
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
