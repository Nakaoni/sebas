package command

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
)

type Command struct {
	Cmd  string
	Args []string
	Dir  string
}

func NewCommand(cmd string, args []string) *Command {
	return &Command{
		Cmd:  cmd,
		Args: args,
		Dir:  "",
	}
}

func (command *Command) UpdateCmd(cmd string) {
	command.Cmd = cmd
}

func (command *Command) UpdateArgs(args []string) {
	command.Args = args
}

func (command *Command) UpdateDir(dir string) {
	command.Dir = dir
}

func (command *Command) Run(ctx context.Context, output chan string) {
	cmd := exec.CommandContext(ctx, command.Cmd, command.Args...)

	if command.Dir != "" {
		cmd.Dir = command.Dir
	}

	defer close(output)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		output <- fmt.Sprint(err)
		return
	}

	if err := cmd.Start(); err != nil {
		output <- fmt.Sprint(err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		output <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		output <- fmt.Sprint(err)
		return
	}

	if err := cmd.Wait(); err != nil {
		output <- fmt.Sprint(err)
		return
	}
}
