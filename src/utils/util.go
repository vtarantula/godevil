package utils

import (
	"bytes"
	"errors"
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

func RunPipedCommand(strcmd string) ([]string, error) {
	cmd := exec.Command("bash", "-c", strcmd)
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	l_mod := make([]string, 0)
	err := cmd.Run()
	if err != nil {
		return l_mod, err
	}

	outstr, errstr := stdout.String(), stderr.String()
	if len(errstr) > 0 {
		return l_mod, errors.New(errstr)
	}
	l_mod = strings.Split(outstr, "\n")
	return l_mod, nil
}
