package cmd_test

import (
	"os"
	"testing"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

func TestCdCommand(t *testing.T) {
	cd := &cmd.CdCommand{}
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("could not change directory: %v", err)
	}

	t.Run("no args", func(t *testing.T) {
		_, err := cd.Execute([]string{})
		if err == nil {
			t.Errorf("[FAIL] Expected and error when no arguments are provided, but got nil")
		} else {
			t.Logf("[PASS] Got expected error for no args: %v", err)
		}
	})

	t.Run("valid path", func(t *testing.T) {
		args := []string{".", "..", "C:\\Users"}
		for _, arg := range args {
			_, err := cd.Execute([]string{arg})
			if err != nil {
				t.Errorf("[FAIL] Expected no error when '%s' arg is passed, but got '%s'", arg, err)
			} else {
				t.Logf("[PASS] Successfully ran cd with '%s' arg", arg)
			}
		}
	})

	t.Run("invalid path", func(t *testing.T) {
		args := []string{"invalid path", "C:\\dasd\\dasa"}
		for _, arg := range args {
			_, err := cd.Execute([]string{arg})
			if err == nil {
				t.Errorf("[FAIL] Expected an error when '%s' invalid path is provided, but got nil", arg)
			} else {
				t.Logf("[PASS] Got expected error for invalid path '%s'", err)
			}
		}

	})
}
