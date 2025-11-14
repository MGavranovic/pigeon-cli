package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

type TouchCommand struct{}

func (c *TouchCommand) Name() string {
	return "touch"
}

func (c *TouchCommand) Description() string {
	return "command for creating empty files"
}

func (c *TouchCommand) Execute(args []string) (bool, error) {
	if len(args) != 1 {
		return false, fmt.Errorf("please specify the 1 filename you are trying to create")
	}

	file := args[0]
	if _, err := os.Stat(file); err == nil {
		fmt.Println("File already exist, do you want to overwrite it?")
		for {
			fmt.Println("Y/N?")

			keyboard.Close()
			var input string
			fmt.Scan(&input)
			input = strings.ToUpper(strings.TrimSpace(input))

			keyboard.Open()
			switch input {
			case "Y":
				newFile, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
				if err != nil {
					return false, fmt.Errorf("issue creating a file: %s", err)
				}
				defer newFile.Close()
				fmt.Println("File successfully overwritten!")
				return false, nil
			case "N":
				return false, fmt.Errorf("file already exists, and you chose not to overwrite it")
			default:
				fmt.Println("Please enter Y or N")
			}
		}
	} else {
		newFile, err := os.Create(file)
		if err != nil {
			return false, fmt.Errorf("error creating a new file '%s': %s", file, err)
		}
		defer newFile.Close()
	}

	return false, nil
}
