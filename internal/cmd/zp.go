package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/MGavranovic/pigeon-cli/internal/helpers"
)

type ZpCommand struct{}

func (c *ZpCommand) Name() string {
	return "zp"
}

func (c *ZpCommand) Description() string {
	return fmt.Sprintf("Zips/Compresses a file or a list of files, into a single zip file\n%-10s - zp <flags[-u(unzip), if not used, zip is default, -f(files to be zipped), -p(path where files can be found, if not used, current dir is used)]>\n", "syntax:")
}

func (c *ZpCommand) Execute(args []string) (bool, error) {
	fmt.Println("ZP with args =>", args)
	files := []string{}
	if len(args) == 0 {
		return false, fmt.Errorf("please provide arguments for the zp command")
	}
	if slices.Contains(args, "-f") {
		i := slices.Index(args, "-f")
		for _, file := range args[i+1:] {
			fmt.Println("Files => ", files)
			if strings.HasPrefix(file, "-") {
				break
			}
			files = append(files, file)
		}
	}

	// TODO: need to check the array
	fmt.Println("Files after loop completed => ", files)
	/*
		1. check if files exist
		2. check if valid files
		3. files will be looked for in ./ unless -p (path) arg is provided -p <path>
	*/
	if slices.Contains(args, "-u") {
		helpers.Unzip()
	} else {
		helpers.Zip()
	}
	return false, nil
}
