package command

import (
	"context"
	"fmt"
	"log"
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

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		c <- fmt.Sprint(err)
	}

	c <- string(out)
}
