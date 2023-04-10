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

func RunCommand(ctx context.Context, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	return buf.Bytes(), err
}

func RunCommandWithBlock(ctx context.Context, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Start()
	return buf.Bytes(), err
}
