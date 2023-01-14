package common

import (
	"io"
	"os/exec"

	"github.com/creack/pty"
)

type Command struct {
	Cmd  *exec.Cmd
	File io.ReadWriteCloser
}

func StartCommand(name string, arg ...string) (*Command, error) {
	cmd := exec.Command(name, arg...)
	f, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}

	return &Command{
		Cmd:  cmd,
		File: f,
	}, nil
}

func (c *Command) Wait() error {
	return c.Cmd.Wait()
}
