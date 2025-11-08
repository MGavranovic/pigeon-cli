package cmd

import (
	"fmt"
	"os"
	"slices"
)

type LsCommand struct{}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Description() string {
	return "Lists all the items in the current dir"
}

func (c *LsCommand) Execute(args []string) (bool, error) {
	dir, err := os.ReadDir(".")
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}
	errorCount := 0
	for _, f := range dir {
		fInfo, err := f.Info()
		if err != nil {
			return false, fmt.Errorf("unable to get file info from %s: %s ", f, err)
		}

		for i := 0; i < len(args); i++ {
			if slices.Contains(args, "-a") {
				fName := f.Name()
				fSize := fInfo.Size()
				fMode := fInfo.Mode()

				fModY, fModM, fModD := fInfo.ModTime().Local().Date()
				fModH, fModMin, fModS := fInfo.ModTime().Local().Clock()
				timeFormatted := fmt.Sprintf("%02d:%02d:%02d", fModH, fModMin, fModS)

				fmt.Printf("%-10s > %10db > %s > %d/%d/%d, %s\n", fName, fSize, fMode, fModD, fModM, fModY, timeFormatted)
			} else {
				errorCount++
				if errorCount > 1 {
					break
				} else {
					return false, fmt.Errorf("ls command doesn't recognize the flag %s", args[i])

				}
			}
		}
		if len(args) == 0 {
			fmt.Printf("%-10s\n", fInfo.Name())
		}
	}
	return false, nil
}
