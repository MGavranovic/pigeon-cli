package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MGavranovic/pigeon-cli/internal/autocomplete"
	"github.com/MGavranovic/pigeon-cli/internal/cmd"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

/*
DOC:
pos is to be used for moving with arrow keys for selecting an autocomplete suggestion
pos = 0 - default or the original cursor position for the console line 	where you type
pos = 1 - first proposed suggestion
pos = pos + 1 - higlight the next pos
pos = pos - 1 - highlight the prev pos
pos = len(suggestsions) last position NOTE: prob have to adjust this
*/
var pos int = 0

/*
DOC:
cursor is used to move left and right in the user input string
*/
var cursor int = 0
var wdLen int

func main() {
	fmt.Println("Hello World!\nWelcome to pigeon-cli!")

	// open keyboard
	err := keyboard.Open()
	if err != nil {
		color.Red("Error opening keyboard: %s\n", err)
	}
	defer keyboard.Close()

	commands := make(map[string]cmd.Command)
	for _, c := range cmd.AllCommands() {
		commands[c.Name()] = c
	}

	help := &cmd.HelpCommand{Commands: commands}
	commands[help.Name()] = help

	historyEntries := []cmd.Entry{}
	history := &cmd.HistoryCommand{Entries: historyEntries}
	commands[history.Name()] = history

	// scanner := bufio.NewReader(os.Stdin)
	suppressPrompt := false

	// autocomplete start
	ac := autocomplete.New(commands)
	ac.Start()

	for {
		if !suppressPrompt {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error accessing path: %s", err)
				continue
			}
			fmt.Printf("%s: ", wd)
			wdLen = len(fmt.Sprint(wd))
		}
		suppressPrompt = false

		// TODO:
		buffer := ""
		var input []rune

		for {
			r, key, err := keyboard.GetKey()
			if err != nil {
				color.Red("Error getting the key: %s", err)
			}
			switch key {
			case keyboard.KeyTab:
				pos = 0
				autocomplete.RenderSuggestions(ac, pos)
			case keyboard.KeyEnter:
				fmt.Printf("\n")
				fmt.Print("\033[0J")
				goto EXECUTE
			case keyboard.KeySpace:
				fmt.Printf(string(32))
				input = append(input, ' ')
				if cursor < len(input) {
					cursor++
				}
				ac.UpdatePrefix(string(input))
			case keyboard.KeyBackspace:
				if len(input) > 0 {
					if cursor > 0 {
						cursor--
					}
					input = input[:len(input)-1]
					ac.UpdatePrefix(string(input))
					fmt.Printf("\b \b")
				}
			case keyboard.KeyArrowDown:
				if pos < len(ac.GetSuggestions()) {
					pos++
					autocomplete.RenderSuggestions(ac, pos)
				}
			case keyboard.KeyArrowUp:
				if pos != 0 {
					pos--
					autocomplete.RenderSuggestions(ac, pos)
				}
				// input = []rune{}
				// fmt.Printf("%s", historyEntries[len(historyEntries)-1])
			case keyboard.KeyArrowLeft:
				if cursor > 0 {
					cursor--
					fmt.Print("\033[D")
				}
			case keyboard.KeyArrowRight:
				if cursor < len(input) {
					cursor++
					fmt.Print("\033[C")
				}
			default:
				fmt.Printf("%s", buffer+string(r))
				cursor++
				input = append(input, r)
				ac.UpdatePrefix(string(input))
			}
		}
		// NOTE:
		// if !scanner.Scan() {
		// 	break
		// }
	EXECUTE:
		line := strings.TrimSpace(string(input))

		tokens := strings.Fields(line)
		if len(tokens) == 0 {
			continue
		}
		cmdName := tokens[0]
		args := tokens[1:]
		if c, ok := commands[cmdName]; ok {
			sp, err := c.Execute(args)
			suppressPrompt = sp

			fullCmd := strings.Join(args, " ")

			if err != nil {
				color.Red("Error running the command %s: %s\n", cmdName, err)
				history.Entries = append(history.Entries, cmd.Entry{Cmd: cmdName, Args: fullCmd, Success: false})
			} else {
				history.Entries = append(history.Entries, cmd.Entry{Cmd: cmdName, Args: fullCmd, Success: true})
			}
		} else {
			color.Red("Unkwnown command %s.\nPlease use help command to get a list of all available commands\n", cmdName)
		}
	}
}
