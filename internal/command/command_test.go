package command

import (
	"testing"

	Assert "github.com/e-felix/sebas/internal/util/assert"
)

func TestNewCommand(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")
	expected := &Command{Path: cmd, Args: args}

	command := NewCommand(cmd, args)

	Assert.DeepEqual(command, expected)
}

func TestCommandUpdatePath(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")

	newCmd := "ls"
	expected := &Command{Path: newCmd, Args: args}

	command := &Command{Path: cmd, Args: args}
	command.UpdatePath(newCmd)

	Assert.DeepEqual(command, expected)
}

func TestCommandUpdateArgs(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")

	newArgs := make([]string, 0)
	newArgs = append(newArgs, "Hi")
	expected := &Command{Path: cmd, Args: newArgs}

	command := &Command{Path: cmd, Args: args}
	command.UpdateArgs(newArgs)

	Assert.DeepEqual(command, expected)
}

func TestCommandRun(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")
	expected := "Hello\n"

	ch := make(chan string)

	command := &Command{Path: cmd, Args: args}
	go command.Run(ch)

	Assert.DeepEqual(<-ch, expected)
}
