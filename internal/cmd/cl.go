package cmd

import "fmt"

type ClCommand struct{}

func (c *ClCommand) Name() string {
	return "cl"
}

func (c *ClCommand) Description() string {
	return "clears the console window"
}

func (c *ClCommand) Execute(args []string) (bool, error) {
	if len(args) > 0 {
		return false, fmt.Errorf("cl requires no args")
	}
	fmt.Print("\033[H\033[2J")
	return false, nil
}
