package autocomplete

import (
	"fmt"

	"github.com/fatih/color"
)

func RenderSuggestions(ac *Engine, pos int) []Suggestion {
	fmt.Print("\033[s")  // save cursor pos
	fmt.Print("\033[0J") // clear below cursor
	fmt.Print("\033[E")  // move a line below

	descColor := color.RGB(55, 69, 96)
	cmdColor := color.RGB(1, 1, 0)
	highlightDesc := color.RGB(1, 1, 0)
	highlightCmd := color.RGB(55, 69, 96)

	fileColor := color.RGB(255, 255, 255)
	dirColor := color.RGB(0, 0, 255)
	highlight := color.RGB(100, 159, 237)

	/*
		DOC:
		compare the pos to current index in for loop
	*/
	suggestions := ac.GetSuggestions()
	for i, s := range suggestions {
		gap := 10 - len(s.Cmd)
		toColorCmd := fmt.Sprintf("%s", s.Cmd)
		for range gap {
			toColorCmd += " "
		}

		if pos == i+1 {
			if s.Desc == "dir" {
				coloredDir := highlight.Sprintf("%s", s.File)
				fmt.Printf("%s\n", coloredDir)
			} else if s.File != "" {
				coloredFile := highlight.Sprintf("%s", s.File)
				fmt.Printf("%s\n", coloredFile)
			} else {
				coloredDesc := highlightDesc.AddBgRGB(1, 0, 1).Sprintf("%s", s.Desc)
				coloredCmd := highlightCmd.AddBgRGB(127, 148, 189).Sprintf("%s", toColorCmd)
				fmt.Printf("%s %*s\n", coloredCmd, gap, coloredDesc)
			}
		} else {
			if s.Desc == "dir" {
				coloredDir := dirColor.Sprintf("%s", s.File)
				fmt.Printf("%s\n", coloredDir)
			} else if s.File != "" {
				coloredFile := fileColor.Sprintf("%s", s.File)
				fmt.Printf("%s\n", coloredFile)
			} else {
				coloredDesc := descColor.AddBgRGB(127, 148, 189).Sprintf("%s", s.Desc)
				coloredCmd := cmdColor.AddBgRGB(1, 0, 1).Sprintf("%s", toColorCmd)
				fmt.Printf("%s %*s\n", coloredCmd, gap, coloredDesc) // print suggestions
			}
		}
	}

	fmt.Print("\033[u") // restore original pos
	return suggestions
}
