package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TreeCommand struct{}

func (c *TreeCommand) Name() string {
	return "tree"
}

func (c *TreeCommand) Description() string {
	return "displays the directory structure"
}

func (c *TreeCommand) Execute(args []string) (bool, error) {
	if len(args) == 0 {
		return false, printTree(".", "")
	}
	if len(args) > 1 {
		return false, fmt.Errorf("please specify the path")
	}
	return false, printTree(args[0], "")
}

func printTree(path string, prefix string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading %s", err)
	}
	for i, f := range dir {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}
		isLastItem := i == len(dir)-1
		connector := "|--"
		if isLastItem {
			connector = "|__"
		}
		fmt.Printf("%s%s %s\n", prefix, connector, f.Name())

		if f.IsDir() {
			newPrefix := prefix
			if isLastItem {
				newPrefix += "   "
			} else {
				newPrefix = "|   "
			}
			err := printTree(filepath.Join(path, f.Name()), newPrefix)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
