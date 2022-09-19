package utils

import (
	"os/exec"
	"strings"
)

// Use this only when you are sure that stderr will not be populated
// For e.g., uname will always give a result in Linux
func RunCommand(cmd string, a ...string) (string, error) {
	scmd := exec.Command(cmd, a...)
	out, err := scmd.Output()

	if err != nil {
		return "", err
	}

	op := string(out)
	op = strings.Trim(op, "\n")
	op = strings.Trim(op, " ")
	return op, nil
}
