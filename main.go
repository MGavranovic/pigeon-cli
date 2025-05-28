package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

func main() {
	fmt.Println("Hello World!\nWelcome to pigeon-cli!")

	commands := make(map[string]cmd.Command)
	for _, c := range cmd.AllCommands() {
		commands[c.Name()] = c
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error accessing path: %s", err)
			continue
		}
		fmt.Printf("%s: ", wd)

		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())

		if line == "exit" {
			fmt.Println("See you soon!")
			os.Exit(0)
			break
		}

		tokens := strings.Fields(line)
		cmdName := tokens[0]
		args := tokens[1:]
		if c, ok := commands[cmdName]; ok {
			err := c.Execute(args)
			if err != nil {
				fmt.Printf("Error running the command %s: %s\n", cmdName, err)
			}
		} else {
			fmt.Printf("Unkwnown command %s.\n", cmdName)
		}
	}
}
