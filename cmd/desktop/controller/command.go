package controller

import "github.com/e-felix/sebas/internal/command"

func RunCommand(cmd command.Command) string {
	ch := make(chan string)
	defer close(ch)

	go cmd.Run(ch)

	return <-ch
}
