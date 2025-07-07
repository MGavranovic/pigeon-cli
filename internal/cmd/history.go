package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

type Entry struct {
	Cmd     string
	Args    string
	Success bool
}

type HistoryCommand struct {
	Entries []Entry
}

func (c *HistoryCommand) Name() string {
	return "history"
}

func (c *HistoryCommand) Description() string {
	return "shows a list of previous commands during this session"
}

func (c *HistoryCommand) Execute(args []string) (bool, error) {
	if len(args) > 0 {
		return false, fmt.Errorf("history cmd requires no args")
	}
	for _, entry := range c.Entries {
		if entry.Success {
			color.Green(" - cmd: %s %s\n", entry.Cmd, entry.Args)
		} else {
			color.Red(" - cmd: %s %s\n", entry.Cmd, entry.Args)
		}
	}
	return false, nil
}
