package cmd

import (
	"fmt"
	"os"
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
	path := ""

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
	if slices.Contains(args, "-p") { // check if path provided and assign it
		i := slices.Index(args, "-p")
		for _, p := range args[i+1:] {
			if strings.HasPrefix(path, "-") {
				break
			}
			path = p
		}
	} else { // otherwise, path set to cwd
		path = "."
	}

	// reading the dir from path provided
	dir, err := os.ReadDir(path)
	fmt.Println("reading =>", dir)
	if err != nil {
		return false, fmt.Errorf("issues with reading files from %s,", path)
	}

	dirFiles := map[string]bool{}
	for _, f := range dir {
		if !f.IsDir() {
			dirFiles[f.Name()] = true
		}
	}
	fmt.Println("dirFiles => ", dirFiles)
	missingFiles := []string{}
	for _, f := range files {
		if !dirFiles[f] {
			missingFiles = append(missingFiles, f)
		}
	}

	if len(missingFiles) > 0 {
		return false, fmt.Errorf("path %s is missing the following files %s", path, missingFiles)
	}
	/*
		1. check if files exist
		2. check if valid files
		3. files will be looked for in ./ unless -p (path) arg is provided -p <path>
		4. if no -n arg, zip file will have a default name
	*/
	if slices.Contains(args, "-u") {
		helpers.Unzip(files)
	} else {
		helpers.Zip(files, "")
	}
	return false, nil
}
