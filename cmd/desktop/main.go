package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"github.com/e-felix/sebas/internal/project"
)

const (
	APP_TITLE     = "Sebas"
	WINDOW_WIDTH  = 1024
	WINDOW_HEIGHT = 728
)

func main() {
	data := initData()
	application := app.New()
	application.Settings().SetTheme(newSebasTheme())

	window := application.NewWindow(APP_TITLE)
	window.Resize(fyne.Size{Width: WINDOW_WIDTH, Height: WINDOW_HEIGHT})

	ui := &gui{
		win:                window,
		file:               binding.NewString(),
		data:               data,
		currentView:        binding.NewString(),
		currentProjectName: binding.NewString(),
	}
	window.SetContent(ui.makeUi())
	//window.SetMainMenu(ui.makeMenu())

	window.CenterOnScreen()
	window.Show()
	application.Run()
}

func initData() map[string]*project.Project {
	return initFixture()
}
