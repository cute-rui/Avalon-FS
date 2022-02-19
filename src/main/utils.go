package main

import (
	"bytes"
	"os/exec"
)

func RunCommand(stdout, stderr *bytes.Buffer, command string, args ...string) error {
	c := exec.Command(command, args...)
	c.Stdout = stdout
	c.Stderr = stderr

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
