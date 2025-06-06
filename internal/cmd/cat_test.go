package cmd_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

func TestCatCommand(t *testing.T) {
	cat := &cmd.CatCommand{}
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("could not change directory: %v", err)
	}

	t.Run("no arguments", func(t *testing.T) {
		err := cat.Execute([]string{})
		if err == nil {
			t.Errorf("[FAIL] Expected an error when no arguments are provided, but got nil")
		} else {
			t.Logf("[PASS] Got expected error: %v", err)
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		fName := "non_existant_file.txt"
		err := cat.Execute([]string{fName})
		if err == nil {
			t.Errorf("[FAIL] Expected error for nonexistent file '%s', but got nil", fName)
		} else {
			t.Logf("[PASS] Got expected error for nonexistent file '%s': %v", fName, err)
		}
	})

	t.Run("valid file", func(t *testing.T) {
		cwd, _ := os.Getwd()
		fmt.Println("[DEBUG] Current working directory:", cwd)

		fName := "main.go"
		err := cat.Execute([]string{fName})
		if err != nil {
			t.Errorf("[FAIL] Expected no error when reding valid file '%s', but got '%s'", fName, err)
		} else {
			t.Logf("[PASS] Successfully read from valid file '%s'", fName)
		}
	})
}
