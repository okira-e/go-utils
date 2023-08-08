package go_utils

import (
	"os"
	"os/exec"
)

// SpawnProcess runs a process and waits for it to finish.
func SpawnProcess(dir string, silent bool, doWait bool, name string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, args...)

	if !silent {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	// Set the working directory of the child processes.
	cmd.Dir = dir

	var err error
	if doWait {
		// Wait for the command to finish.
		err = cmd.Run()
	} else {
		// Start the command but don't wait for it to finish.
		err = cmd.Start()
	}

	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// ExecCommand executes a command and returns the output.
// It returns an error if the command fails to start or doesn't complete.
func ExecCommand(name string, args ...string) (string, error) {
	// An error is returned if the command fails to start or doesn't complete
	out, err := exec.Command(name, args...).Output()

	return string(out), err
}
