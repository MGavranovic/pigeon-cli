package helpers

import (
	// "archive/zip"
	"fmt"
	// "io"
	// "os"
)

func Zip(files []string, name string) {
	if name == "" {
		name = "archive.zip"
		fmt.Printf("Zip() => %s into > %s\n", files, name)
	} else {
		fmt.Printf("Zip() => %s into > %s\n", files, name)
	}
}
