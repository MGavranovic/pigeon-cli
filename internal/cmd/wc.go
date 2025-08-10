package cmd

import (
	"fmt"
	"os"
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
	if len(args) < 1 {
		return false, fmt.Errorf("please specify the file you are looking for")
	}
	fName := args[0]
	file, err := os.Stat(fName)
	if err != nil {
		return false, fmt.Errorf("error getting stats of %s: %s", fName, err)
	}

	if file.IsDir() {
		return false, fmt.Errorf("%s is a directory", fName)
	}

	fContent, err := os.ReadFile(fName)
	if err != nil {
		return false, fmt.Errorf("error reading content of %s", fName)
	}
	wc := 0
	words := strings.Fields(string(fContent))
	for i := 0; i < len(words); i++ {
		wc++
	}
	fmt.Printf("%s is %d bytes long and has %d words\n", file.Name(), file.Size(), wc)
	return false, nil
}
