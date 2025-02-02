// Package command
// Date: 2023/4/6 15:35
// Author: Amu
// Description:
package command

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os/exec"
	"syscall"
	"time"
)

func RunCommand(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	if err := cmd.Start(); err != nil {
		return buf.String(), err
	}
	if err := cmd.Wait(); err != nil {
		return buf.String(), err
	}
	return buf.String(), nil
}

func RunCommandWithTimeout(ctx context.Context, timeout uint32, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf

	if err := cmd.Start(); err != nil {
		return buf.String(), err
	}

	timer := time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		_ = cmd.Process.Signal(syscall.SIGTERM)
	})

	if err := cmd.Wait(); err != nil {
		return buf.String(), err

	}
	timer.Stop()
	return buf.String(), nil
}

func StreamRunCommand(callback func(output string), name string, args ...string) error {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	reader := bufio.NewReaderSize(io.MultiReader(stdout, stderr), 128)
	if err = cmd.Start(); err != nil {
		return err
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		callback(line)
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
