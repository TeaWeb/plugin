package apps

import (
	"errors"
	"os/exec"
	"strings"
)

func Exec(command string, arg ... string) (string, error) {
	binFile, err := exec.LookPath(command)
	if err != nil || len(binFile) == 0 {
		return "", errors.New("command '" + command + "' not found")
	}

	cmd := exec.Command(binFile, arg ...)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	dataString := strings.TrimSpace(string(data))
	if len(dataString) == 0 {
		return "", nil
	}
	return dataString, nil
}
