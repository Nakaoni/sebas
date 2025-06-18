package command

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type Command struct {
	Path string
	Args []string
}

func NewCommand(path string, args []string) *Command {
	return &Command{
		Path: path,
		Args: args,
	}
}

func (command *Command) UpdatePath(cmd string) {
	command.Path = cmd
}

func (command *Command) UpdateArgs(args []string) {
	command.Args = args
}

func (command *Command) Run(c chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, command.Path, command.Args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c <- fmt.Sprint(err)
	}

	if err := cmd.Start(); err != nil {
		c <- fmt.Sprint(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		c <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		c <- fmt.Sprint(err)
	}

	if err := cmd.Wait(); err != nil {
		c <- fmt.Sprint(err)
	}
}
