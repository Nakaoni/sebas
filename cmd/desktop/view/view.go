package view

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/e-felix/sebas/internal/util"
)

const WINDOW_TITLE = "Sebas - Your Personal Butler"
const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 728
const COMMANDS_SECTION_TITLE = "COMMANDS"
const COMMANDS_MENU_LABEL = "Commands"

func Render() {
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

	project_select := widget.NewSelect(projects_names, func(selected_value string) {
		log.Println("Project seleted: ", selected_value)
		selected_project = selected_value
	})
	project_select.PlaceHolder = "Select a project"

	project_cmds_container := container.NewVBox()
	cmds_container := container.NewVBox()
	middle_container := container.NewGridWithColumns(1, cmds_container)

	top_right_container := container.NewGridWithColumns(1, project_select)
	top_top_section := container.NewGridWithColumns(
		6,
		top_left_container,
		container.NewVBox(),
		container.NewVBox(),
		container.NewVBox(),
		container.NewVBox(),
		top_right_container,
	)
	top_bottom_section := canvas.NewLine(color.White)
	top_container := container.NewVBox(top_top_section, top_bottom_section)

	cmds_menu_button := widget.NewButton(COMMANDS_MENU_LABEL, func() {
		cmds_container.RemoveAll()
		project_cmds_container.RemoveAll()
		middle_container.RemoveAll()

		cmds_section_label := widget.NewLabel(COMMANDS_SECTION_TITLE)
		cmds_section_divider := canvas.NewLine(color.White)

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

					updateWidget(output, <-c)
				})

				container := container.NewGridWithColumns(3, path_label, args_label, button)
				project_cmds_container.Add(container)
			}
		}

		cmds_container.Add(cmds_section_label)
		cmds_container.Add(cmds_section_divider)
		cmds_container.Add(project_cmds_container)
		middle_container.Add(cmds_container)
		middle_container.Refresh()
	})
	left_container := container.NewHBox(container.NewVBox(cmds_menu_button), canvas.NewLine(color.White))

	version_label := widget.NewLabel(fmt.Sprintf("version %v", getCurrentVersion()))
	version_label.TextStyle.Italic = true
	bottom_container := container.NewVBox(canvas.NewLine(color.White), container.NewHBox(version_label))

	content := container.NewBorder(top_container, bottom_container, left_container, nil, middle_container)

	window.SetContent(content)
	window.SetMaster()

	window.Show()
	my_app.Run()
}

func updateWidget(w *widget.Label, msg string) {
	w.SetText(msg)
}

func getCurrentVersion() string {
	var version = ""

	cwd, err := os.Getwd()

	if err != nil {
		return version
	}

	version, err = util.GetFileContent(path.Join(cwd, "/../../", "VERSION"))

	return version
}
