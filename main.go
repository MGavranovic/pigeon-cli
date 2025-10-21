package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MGavranovic/pigeon-cli/internal/autocomplete"
	"github.com/MGavranovic/pigeon-cli/internal/cmd"
	"github.com/MGavranovic/pigeon-cli/internal/inputpkg"
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

/*
DOC:
for global state
*/
var cwd string
var historyPos int = 0

/*
DOC:
- used for changing the mode to be able to better navigate the terminal
history enables up and down arrows when pos = 0
- input enables user typing, tab key, backspace and enter for suggestions and running the cmds
- ac, initiated on tab, enables up and down
*/
var mode string = "input" // "input", "autocomplete", "history"

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
			cwd = wd
		}
		suppressPrompt = false

		var input []rune
		for {
			r, key, err := keyboard.GetKey()
			if err != nil {
				color.Red("Error getting the key: %s", err)
			}
			switch key {
			case keyboard.KeyTab:
				pos = 0
				mode = "autocomplete"
				autocomplete.RenderSuggestions(ac, pos)
			case keyboard.KeyEnter:
				fmt.Printf("\n")
				fmt.Print("\033[0J")
				mode = "input"
				goto EXECUTE
			case keyboard.KeySpace:
				fmt.Printf(string(32))
				input = append(input, ' ')
				if cursor < len(input) {
					cursor++
				}
				ac.UpdatePrefix(string(input))
			case keyboard.KeyBackspace:
				if len(input) > 0 && cursor > 0 {
					input = append(input[:cursor-1], input[cursor:]...)
					cursor--
					inputpkg.RedrawInput(cwd, input, cursor)
					ac.UpdatePrefix(string(input))
				}
			case keyboard.KeyArrowDown:
				if mode == "autocomplete" {
					if pos < len(ac.GetSuggestions()) {
						pos++
						autocomplete.RenderSuggestions(ac, pos)
					}
				} else {
					if historyPos > 1 {
						historyPos--
						input = inputpkg.PrintHistory(history.Entries, pos, historyPos, cwd)
					} else {
						historyPos = 0
						input = []rune{}
						cursor = 0
						inputpkg.RedrawInput(cwd, input, cursor)
						mode = "input"
					}
				}
			case keyboard.KeyArrowUp:
				if mode == "autocomplete" {
					if pos > 0 {
						pos--
						autocomplete.RenderSuggestions(ac, pos)
					}
				} else {
					if historyPos < len(history.Entries) {
						historyPos++
						input = inputpkg.PrintHistory(history.Entries, pos, historyPos, cwd)
						mode = "history"
					}
				}
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
				mode = "input"
				if cursor < 0 {
					cursor = 0
				}
				if cursor > len(input) {
					cursor = len(input)
				}
				input = append(input[:cursor], append([]rune{r}, input[cursor:]...)...)
				cursor++
				inputpkg.RedrawInput(cwd, input, cursor)
				ac.UpdatePrefix(string(input))
			}
		}
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
