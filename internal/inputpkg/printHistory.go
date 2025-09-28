package inputpkg

import (
	// "fmt"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

func PrintHistory(history []cmd.Entry, pos int, historyPos int, cwd string) {
	if pos == 0 && len(history) > 0 {
		entry := history[len(history)-historyPos]
		line := entry.Cmd

		if entry.Args != "" {
			line += " " + entry.Args
		}
		runeInput := []rune(line)
		cursor := len(runeInput)
		RedrawInput(cwd, runeInput, cursor)
	}
}
