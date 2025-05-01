package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"github.com/e-felix/sebas/internal/command"
	"github.com/e-felix/sebas/internal/env"
	"github.com/e-felix/sebas/internal/project"
	"strconv"
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

	ui := &gui{win: window, file: binding.NewString(), data: data}
	window.SetContent(ui.makeUi())
	//window.SetMainMenu(ui.makeMenu())

	window.Show()
	application.Run()
}

func initData() map[string]*project.Project {
	projects := make(map[string]*project.Project)

	for i := 1; i <= 3; i++ {
		new_project := project.NewProject(fmt.Sprintf("Project_%d", i))

		for j := 0; j < 3; j++ {
			new_project.AddEnv(env.Env{
				Key:   fmt.Sprintf("ENV_%d", j),
				Value: fmt.Sprintf("VALUE_%d", j),
			})
			new_project.AddCmd(command.Command{
				Path: "echo",
				Args: []string{new_project.Name, strconv.Itoa(j)},
			})
		}

		projects[new_project.Name] = new_project
	}

	return projects
}
