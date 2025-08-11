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
	Suggestions []string
	mu          sync.RWMutex
	Commands    map[string]cmd.Command
}

func New(cmds map[string]cmd.Command) *Engine {
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

			cmdResults := make(chan []string)
			fsResults := make(chan []string)

			go func() {
				matches := []string{}
				for name := range e.Commands {
					if strings.HasPrefix(name, currPrefix) {
						matches = append(matches, name)
					}
				}
				cmdResults <- matches
			}()

			go func() {
				matches := []string{}
				entries, err := os.ReadDir(".")
				if err != nil {
					fmt.Printf("error reading current dir in autocomplete: %s", err)
				}
				for _, e := range entries {
					if strings.HasPrefix(e.Name(), currPrefix) {
						matches = append(matches, e.Name())
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

func (e *Engine) GetSuggestions() []string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return append([]string{}, e.Suggestions...)
}
