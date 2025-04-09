package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/e-felix/sebas/internal/command"
	"github.com/e-felix/sebas/internal/env"
	"github.com/e-felix/sebas/internal/project"
)

const WINDOW_TITLE = "Sebas - Your Personal Butler"
const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600

func main() {
	my_app := app.New()
	window := my_app.NewWindow(WINDOW_TITLE)
	window.Resize(fyne.Size{Width: WINDOW_WIDTH, Height: WINDOW_HEIGHT})
	window.CenterOnScreen()

	sebas_title := widget.NewLabel(WINDOW_TITLE)
	top_left_container := container.NewGridWithColumns(1, sebas_title)

	var selected_project string

	projects := getProjects()
	projects_names := make([]string, 0)
	for _, p := range projects {
		projects_names = append(projects_names, p.Name)
	}

	middle_container := container.NewGridWithColumns(1)
	project_select := widget.NewSelect(projects_names, func(selected_value string) {
		log.Println("Project seleted: ", selected_value)
		selected_project = selected_value
		middle_container.RemoveAll()

		p, ok := projects[selected_project]
		if ok {
			for _, cmd := range p.Cmds {
				path_label := widget.NewLabel(cmd.Path)

				var args_builder strings.Builder
				for _, a := range cmd.Args {
					args_builder.WriteString(a)
					args_builder.WriteString(" ")
				}

				args_label := widget.NewLabel(args_builder.String())

				button := widget.NewButton("Run", func() {
					outputWindow := my_app.NewWindow("Result")
					output := widget.NewLabel("")
					outputWindow.SetContent(output)
					outputWindow.Show()

					c := make(chan string)
					go cmd.Run(c)

					updateTime(output, <-c)
				})

				container := container.NewAdaptiveGrid(3, path_label, args_label, button)
				middle_container.Add(container)
			}
		}
	})
	project_select.PlaceHolder = "Select a project"

	load_context_button := widget.NewButtonWithIcon("Load", theme.MediaReplayIcon(), func() {
		log.Println("Load context button clicked")
	})

	top_right_container := container.NewGridWithColumns(2, project_select, load_context_button)
	top_container := container.NewGridWithColumns(2, top_left_container, top_right_container)

	// label := widget.NewLabel("Hello World")
	//
	// button := widget.NewButton("Print `ls -la`", func() {
	// 	clockWindow := my_app.NewWindow("")
	// 	clock := widget.NewLabel("")
	// 	clockWindow.SetContent(clock)
	// 	clockWindow.Show()
	//
	// 	c := make(chan string)
	// 	cmd := command.NewCommand("ls", []string{"-la"})
	// 	go cmd.Run(c)
	//
	// 	updateTime(clock, <-c)
	//
	// })
	//
	// text_area := widget.NewTextGrid()

	content := container.NewBorder(top_container, nil, nil, nil, middle_container)
	window.SetContent(content)
	window.SetMaster()

	window.Show()
	my_app.Run()
}

func getProjects() map[string]*project.Project {
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
				Args: []string{strconv.Itoa(j)},
			})
		}

		projects[new_project.Name] = new_project
	}

	return projects
}

func updateTime(clock *widget.Label, msg string) {
	clock.SetText(msg)
}
