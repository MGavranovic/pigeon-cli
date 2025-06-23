package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type RmCommand struct{}

func (c *RmCommand) Name() string {
	return "rm"
}

func (c *RmCommand) Description() string {
	return fmt.Sprintf("Command for moving the files\n%-10s - rm <file> <flags[--force]>\n", "syntax:")
}

func (c *RmCommand) Execute(args []string) (bool, error) {
	if len(args) == 0 {
		return false, fmt.Errorf("please specify the file you are trying to delete")
	}

	path := args[0]
	if slices.Contains(args, "--force") {
		err := os.Remove(path)
		if err != nil {
			return false, fmt.Errorf("error removing the file: %s", err)
		}
		fmt.Printf(" - %s - file deleted successfully with '--force'\n", path)
		return false, nil
	} else {
		fmt.Printf("Are you sure you want to remove '%s'?\n", path)
		fmt.Println("Y/N?")

		var input string
		fmt.Scan(&input)
		input = strings.ToUpper(strings.TrimSpace(input))
		switch input {
		case "Y":
			err := os.Remove(path)
			if err != nil {
				return false, fmt.Errorf("error removing the file: %s", err)
			}
			fmt.Printf(" - %s - file deleted successfully\n", path)
			return true, nil
		case "N":
			return true, fmt.Errorf("you chose not to delete the file")
		default:
			fmt.Println("Please enter Y or N")
		}
	}

	return false, nil
}
