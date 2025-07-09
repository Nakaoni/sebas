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
}

func NewCommand(cmd string, args []string) *Command {
	return &Command{
		Cmd:  cmd,
		Args: args,
	}
}

func (command *Command) UpdateCmd(cmd string) {
	command.Cmd = cmd
}

func (command *Command) UpdateArgs(args []string) {
	command.Args = args
}

func (command *Command) Run(ctx context.Context, output chan string) {
	cmd := exec.CommandContext(ctx, command.Cmd, command.Args...)
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
