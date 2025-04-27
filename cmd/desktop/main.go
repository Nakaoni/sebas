package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	APP_TITLE     = "Sebas"
	WINDOW_WIDTH  = 1024
	WINDOW_HEIGHT = 728
)

func main() {
	application := app.New()
	application.Settings().SetTheme(newSebasTheme())

	window := application.NewWindow(APP_TITLE)
	window.Resize(fyne.Size{Width: WINDOW_WIDTH, Height: WINDOW_HEIGHT})

	window.SetContent(makeUi())

	window.Show()
	application.Run()
}
