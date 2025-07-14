package main

import (
	"context"
	"fmt"
	"github.com/e-felix/sebas/internal/command"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	cmd := command.NewCommand("cmd", []string{"/C", "dir"})
	cmd.UpdateDir("E:\\")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	output := make(chan string)
	go cmd.Run(ctx, output)

	for value := range output {
		fmt.Println(value)
	}
}
