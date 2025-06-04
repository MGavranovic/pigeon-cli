package cmd_test

import (
	"testing"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

func TestCatCommand(t *testing.T) {
	cat := &cmd.CatCommand{}

	t.Run("no arguments", func(t *testing.T) {
		err := cat.Execute([]string{})
		if err == nil {
			t.Errorf("[FAIL] Expected an error when no arguments are provided, but got nil")
		} else {
			t.Logf("[PASS] Got expected error: %v", err)
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		filename := "non_existant_file.txt"
		err := cat.Execute([]string{filename})
		if err == nil {
			t.Errorf("[FAIL] Expected error for nonexistent file '%s', but got nil", filename)
		} else {
			t.Logf("[PASS] Got expected error for nonexistent file '%s': %v", filename, err)
		}
	})
}
