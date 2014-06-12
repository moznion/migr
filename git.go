package main

import (
	"os"
	"os/exec"
	"strings"
)

func getBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stderr = os.Stderr
	buf, err := cmd.Output()

	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(buf), "\r\n"), nil
}
