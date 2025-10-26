package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func Zip(files []string, name string) error {
	zipName := ""
	if name == "" {
		zipName = "archive.zip"
		fmt.Printf("Zip() => %s into > %s\n", files, name)
	} else {
		zipName = name
	}
	f, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		w, err := writer.Create(file)
		if err != nil {
			return err
		}
		_, errCopy := io.Copy(w, f)
		if errCopy != nil {
			return errCopy
		}
	}
	return nil
}
