// Package command
// Date: 2023/4/6 15:35
// Author: Amu
// Description:
package command

import (
	"bytes"
	"context"
	"os/exec"
)

func RunCommand(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	return buf.String(), err
}

func RunCommandWithBlock(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Start()
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return buf.String(), err
}
