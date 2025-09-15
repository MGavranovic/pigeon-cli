package autocomplete

import (
	"fmt"

	"github.com/fatih/color"
)

func RenderSuggestions(ac *Engine, pos int) {
	fmt.Print("\033[s")  // save cursor pos
	fmt.Print("\033[0J") // clear below cursor
	fmt.Print("\033[E")  // move a line below

	descColor := color.RGB(55, 69, 96)
	cmdColor := color.RGB(1, 1, 0)
	highlightDesc := color.RGB(1, 1, 0)
	highlightCmd := color.RGB(55, 69, 96)

	/*
		DOC:
		compare the pos to current index in for loop
	*/
	for i, s := range ac.GetSuggestions() {
		gap := 10 - len(s.Cmd)
		toColorCmd := fmt.Sprintf("%s", s.Cmd)
		for range gap {
			toColorCmd += " "
		}
		// fmt.Println("pos", pos, i+1)
		if pos == i+1 {
			// fmt.Println("pos", pos, i+1)
			coloredDesc := highlightDesc.AddBgRGB(1, 0, 1).Sprintf("%s", s.Desc)
			coloredCmd := highlightCmd.AddBgRGB(127, 148, 189).Sprintf("%s", toColorCmd)
			fmt.Printf("%s %*s\n", coloredCmd, gap, coloredDesc)
		} else {
			coloredDesc := descColor.AddBgRGB(127, 148, 189).Sprintf("%s", s.Desc)
			coloredCmd := cmdColor.AddBgRGB(1, 0, 1).Sprintf("%s", toColorCmd)
			fmt.Printf("%s %*s\n", coloredCmd, gap, coloredDesc) // print suggestions
		}
	}

	fmt.Print("\033[u") // restore original pos
}
