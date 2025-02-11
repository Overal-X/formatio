package services

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type ExecuteArgs struct {
	Directory      string
	Command        string
	OutputCallback func(string)
	ErrorCallback  func(string)
}

type IExecService interface {
	Execute(args ExecuteArgs) error
}

type ExecService struct{}

func (e *ExecService) Execute(args ExecuteArgs) error {
	command := strings.TrimSpace(args.Command)
	fmt.Println("> ", command)

	currentCmd := exec.Command("/bin/bash", "-c", command)

	// Create buffers for stdout and stderr
	var stdoutBuf, stderrBuf bytes.Buffer

	// Create multiwriters to write to both buffer and callback
	stdout := io.MultiWriter(&stdoutBuf, writeCallback{callback: args.OutputCallback})
	stderr := io.MultiWriter(&stderrBuf, writeCallback{callback: args.ErrorCallback})

	currentCmd.Stdout = stdout
	currentCmd.Stderr = stderr

	// Run the command and wait for completion
	if err := currentCmd.Run(); err != nil {
		if args.ErrorCallback != nil {
			args.ErrorCallback(err.Error())
		}
		return fmt.Errorf("command failed: %w", err)
	}

	return nil
}

// writeCallback implements io.Writer for streaming output
type writeCallback struct {
	callback func(string)
}

func (w writeCallback) Write(p []byte) (n int, err error) {
	if w.callback != nil {
		w.callback(string(bytes.TrimSpace(p)))
	}
	return len(p), nil
}

func NewExecService() IExecService {
	return &ExecService{}
}
