package main

import (
	"fmt"

	"github.com/e-felix/sebas/internal/command"
)

func main() {
	fmt.Println("Sebas CLI")
	fmt.Println("Running `ls -la`")

	cmd := command.NewCommand("ls", []string{"-la"})

	ch := make(chan string)
	go cmd.Run(ch)

	v, ok := <-ch

	if ok {
		fmt.Println(v)
	}
}
