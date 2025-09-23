package inputpkg

import "fmt"

func RedrawInput(cwd string, input []rune, cursor int) {
	cwdLen := len(cwd) + 2

	fmt.Printf("\r\033[%dC", cwdLen)
	fmt.Print("\r\033[K")

	fmt.Print(cwd + ": ")
	fmt.Print(string(input))

	// move cursor back if not at end
	if cursor < len(input) {
		back := len(input) - cursor
		fmt.Printf("\033[%dD", back)
	}
}
