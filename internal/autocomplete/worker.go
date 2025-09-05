package autocomplete

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

type Engine struct {
	prefix      string
	Suggestions []Suggestion
	mu          sync.RWMutex
	Commands    map[string]cmd.Command
}

type Suggestion struct {
	Cmd  string
	Arg  string
	File string
	Desc string
}

func New(cmds map[string]cmd.Command) *Engine {
	fmt.Println(&Engine{Commands: cmds})
	return &Engine{Commands: cmds}
}

func (e *Engine) Start() {
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			e.mu.RLock()
			currPrefix := e.prefix
			e.mu.RUnlock()

			if currPrefix == "" {
				continue
			}

			cmdResults := make(chan []Suggestion)
			fsResults := make(chan []Suggestion)

			go func() {
				matches := []Suggestion{}
				for name := range e.Commands {
					desc := e.Commands[name].Description()
					if strings.HasPrefix(name, currPrefix) {
						matches = append(matches, Suggestion{name, "", "", desc})
					}
				}
				cmdResults <- matches
			}()

			go func() {
				matches := []Suggestion{}
				entries, err := os.ReadDir(".")
				if err != nil {
					fmt.Printf("error reading current dir in autocomplete: %s", err)
				}
				for _, e := range entries {
					if strings.HasPrefix(e.Name(), currPrefix) {
						matches = append(matches, Suggestion{"", "", e.Name(), ""})
					}
				}
				fsResults <- matches
			}()

			merged := append(<-cmdResults, <-fsResults...)
			e.mu.Lock()
			e.Suggestions = merged
			e.mu.Unlock()
		}
	}()
}

func (e *Engine) UpdatePrefix(newPrefix string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.prefix = newPrefix
}

func (e *Engine) GetSuggestions() []Suggestion {
	e.mu.RLock()
	defer e.mu.RUnlock()

	sortedSuggestions := []Suggestion{}
	for _, s := range e.Suggestions {
		if s.Cmd != "" {
			sortedSuggestions = append(sortedSuggestions, s)
		}
	}
	for _, s := range e.Suggestions {
		if s.File != "" {
			sortedSuggestions = append(sortedSuggestions, s)
		}
	}
	fmt.Println(sortedSuggestions)
	return append([]Suggestion{}, sortedSuggestions...)
}
