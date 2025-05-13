package controller

import (
	"github.com/e-felix/sebas/internal/command"
	"github.com/e-felix/sebas/internal/project"
)

func RunCommand(cmd command.Command) string {
	ch := make(chan string)
	defer close(ch)

	go cmd.Run(ch)

	return <-ch
}

func EditCommand(p project.Project, cmd command.Command) error {
	return p.EditCmd(cmd)
}
