package controller

import (
	"github.com/e-felix/sebas/internal/command"
)

func RunCommand(cmd command.Command) string {
	ch := make(chan string)
	go cmd.Run(ch)

	return <-ch
}
