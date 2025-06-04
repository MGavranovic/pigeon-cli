package cmd_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/MGavranovic/pigeon-cli/internal/cmd"
)

func TestExitCommand(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcess")
	cmd.Env = append(cmd.Env, "GO_TEST_EXIT=1")

	err := cmd.Run()

	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitErr.ExitCode() != 0 {
			t.Errorf("expected exit code 0, got %d", exitErr.ExitCode())
		}
	} else if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// This runs in the subprocess when the test above is invoked
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_TEST_EXIT") != "1" {
		return
	}

	c := &cmd.ExitCommand{}
	c.Execute(nil)

	t.Fatal("expected os.Exit(0) to terminate the process")
}
