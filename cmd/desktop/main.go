package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/e-felix/sebas/internal/command"
)

func main() {
	my_app := app.New()
	window := my_app.NewWindow("Main")

	label := widget.NewLabel("Hello World")

	button := widget.NewButton("Print `ls -la`", func() {
		clockWindow := my_app.NewWindow("")
		clock := widget.NewLabel("")
		clockWindow.SetContent(clock)
		clockWindow.Show()

		c := make(chan string)
		cmd := command.NewCommand("ls", []string{"-la"})
		go cmd.Run(c)

		updateTime(clock, <-c)

	})

	base_layout := layout.NewGridLayoutWithRows(2)
	content := container.New(base_layout, label, button)
	window.SetContent(content)
	window.SetMaster()

	window.Show()
	my_app.Run()
}

func updateTime(clock *widget.Label, msg string) {
	clock.SetText(msg)
}
