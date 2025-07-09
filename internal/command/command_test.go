package command

import (
	"context"
	"testing"

	Assert "github.com/e-felix/sebas/internal/util/assert"
)

func TestNewCommand(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")
	expected := &Command{Cmd: cmd, Args: args}

	command := NewCommand(cmd, args)

	Assert.DeepEqual(command, expected)
}

func TestCommandUpdateCmd(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")

	newCmd := "ls"
	expected := &Command{Cmd: newCmd, Args: args}

	command := &Command{Cmd: cmd, Args: args}
	command.UpdateCmd(newCmd)

	Assert.DeepEqual(command, expected)
}

func TestCommandUpdateArgs(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")

	newArgs := make([]string, 0)
	newArgs = append(newArgs, "Hi")
	expected := &Command{Cmd: cmd, Args: newArgs}

	command := &Command{Cmd: cmd, Args: args}
	command.UpdateArgs(newArgs)

	Assert.DeepEqual(command, expected)
}

func TestCommandRun(t *testing.T) {
	expected := "Hello"

	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")

	command := &Command{Cmd: cmd, Args: args}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	output := make(chan string)

	go command.Run(ctx, output)
	value := <-output

	Assert.DeepEqual(value, expected)
}
