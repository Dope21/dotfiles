package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunCustomScript(cmdSequence []string, isPath bool) error {
	if isPath {
		path, err := filepath.Abs(cmdSequence[0])
		if err != nil {
			return err
		}
		cmdSequence[0] = path
	}

	cmd := exec.Command(cmdSequence[0], cmdSequence[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}