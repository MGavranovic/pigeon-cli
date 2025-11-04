package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type CatCommand struct{}

func (c *CatCommand) Name() string {
	return "cat"
}

func (c *CatCommand) Description() string {
	return "Reads the contents of a file"
}

func (c *CatCommand) Execute(args []string) (bool, error) {
	if len(args) == 0 {
		return false, fmt.Errorf("please specify the file you want to read with -f")
	}

	path := ""
	if slices.Contains(args, "-f") {
		i := slices.Index(args, "-f")
		for _, f := range args[i+1:] {
			if strings.HasPrefix(f, "-") {
				break
			}
			path = f
		}
		file, err := os.ReadFile(path)
		if err != nil {
			return false, fmt.Errorf("error reading from %s: %s", path, err)
		}

		s := string(file)
		fmt.Println(s)
	} else {
		return false, fmt.Errorf("please specify the file you want to read with -f")
	}
	return false, nil
}
