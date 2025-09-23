package inputpkg

import (
	"fmt"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

var historyPos int = 0

func PrintHistory(history []cmd.Entry, pos int) {
	if pos == 0 && len(history) > 0 {
		prev := history[len(history)-1].Cmd
		runeInput := []rune(prev)
		fmt.Println("this is a previous entry", prev)
		fmt.Println("this is a previous entry", runeInput)
	}
}
