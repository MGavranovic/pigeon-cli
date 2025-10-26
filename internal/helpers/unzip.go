package helpers

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(files []string) error {
	for _, file := range files {
		r, err := zip.OpenReader(file)
		if err != nil {
			return err
		}

		for _, f := range r.File {
			outPath := filepath.Join("output", f.Name)
			err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
			if err != nil {
				return err
			}

			outFile, err := os.Create(outPath)
			if err != nil {
				return err
			}
			rc, err := f.Open()
			if err != nil {
				return err
			}

			io.Copy(outFile, rc)
			outFile.Close()
			rc.Close()
		}
	}
	return nil
}
