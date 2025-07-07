package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

type GrepCommand struct{}

func (c *GrepCommand) Name() string {
	return "grep"
}

func (c *GrepCommand) Description() string {
	return "command for searching for a string in a file"
}

func (c *GrepCommand) Execute(args []string) (bool, error) {
	if len(args) == 0 {
		return false, fmt.Errorf("please specify what you're searching for and the file/path")
	}

	search := args[0]
	fName := args[1]
	file, err := os.ReadFile(fName)
	if err != nil {
		return false, fmt.Errorf("error reading from file %s: %s", fName, err)
	}

	count := 0
	lines := strings.Split(string(file), "\n")
	contains := false
	for i, l := range lines {
		if strings.Contains(l, search) {
			instances := strings.Count(l, search)
			count += instances
			contains = true

			found := color.New(color.FgGreen).SprintFunc()
			arr := strings.Split(l, search)
			final := strings.Join(arr, found(search))

			fmt.Printf("%s found %dx\nLocation: line %d in %s > %s\n", search, count, i+1, fName, final)
		}
	}
	if !contains {
		return false, fmt.Errorf("found no instances of %s", search)
	}
	return false, nil
}
