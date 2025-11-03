package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type TreeCommand struct{}

func (c *TreeCommand) Name() string {
	return "tree"
}

func (c *TreeCommand) Description() string {
	return fmt.Sprintf("displays the directory structure\n%-10s - tree <flags[-p(path), if not used, display the cwd tree]>\n", "syntax:")
}

func (c *TreeCommand) Execute(args []string) (bool, error) {
	path := ""
	if len(args) == 0 {
		return false, printTree(".", "")
	}
	if slices.Contains(args, "-p") {
		i := slices.Index(args, "-p")
		for _, p := range args[i+1:] {
			if strings.HasPrefix(p, "-") {
				break
			}
			path = p
		}
	} else {
		path = "."
	}
	return false, printTree(path, "")
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
