package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type WcCommand struct{}

func (c *WcCommand) Name() string {
	return "wc"
}

func (c *WcCommand) Description() string {
	return "counts words in a file"
}

func (c *WcCommand) Execute(args []string) (bool, error) {
	if len(args) == 0 {
		return false, fmt.Errorf("please specify the file you are looking for using -f arg")
	}

	path := ""
	wc := 0
	if slices.Contains(args, "-f") {
		i := slices.Index(args, "-f")
		for _, f := range args[i+1:] {
			if strings.HasPrefix(f, "-") {
				break
			}
			path = f
		}
		file, err := os.Stat(path)
		if err != nil {
			return false, fmt.Errorf("error getting stats of %s: %s", path, err)
		}

		if file.IsDir() {
			return false, fmt.Errorf("%s is a directory", path)
		}

		fContent, err := os.ReadFile(path)
		if err != nil {
			return false, fmt.Errorf("error reading content of %s", path)
		}
		words := strings.Fields(string(fContent))
		for i := 0; i < len(words); i++ {
			wc++
		}
		fmt.Printf("%s is %d bytes long and has %d words\n", file.Name(), file.Size(), wc)
	} else {
		return false, fmt.Errorf("please specify the file you are looking for using -f arg")
	}

	return false, nil
}
