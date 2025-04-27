package view

import (
	"fmt"
	"image/color"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/e-felix/sebas/cmd/desktop/controller"
	"github.com/e-felix/sebas/internal/project"
	"github.com/e-felix/sebas/internal/util"
)

type AppData struct {
	Selected_project string
	Projects         map[string]*project.Project
}

type AppView struct {
	App    fyne.App
	Window fyne.Window
	Top    *fyne.Container
	Bottom *fyne.Container
	Left   *fyne.Container
	Right  *fyne.Container
	Middle *fyne.Container
}

const WINDOW_TITLE = "Sebas - Your Personal Butler"
const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 728
const COMMANDS_SECTION_TITLE = "COMMANDS"
const COMMANDS_MENU_LABEL = "Commands"
const ENVS_MENU_LABEL = "Envs"

var app_data AppData
var app_view AppView

func Render() {
	app_view.App = app.New()
	app_view.Window = app_view.App.NewWindow(WINDOW_TITLE)
	app_view.Window.Resize(fyne.Size{Width: WINDOW_WIDTH, Height: WINDOW_HEIGHT})
	app_view.Window.CenterOnScreen()

	sebas_title := widget.NewLabel(WINDOW_TITLE)
	top_left_container := container.NewGridWithColumns(1, sebas_title)

	app_data.Projects = controller.GetProjects()
	projects_names := make([]string, 0)
	for _, p := range app_data.Projects {
		projects_names = append(projects_names, p.Name)
	}

	project_select := widget.NewSelect(projects_names, func(selected_value string) {
		app_data.Selected_project = selected_value
	})
	project_select.PlaceHolder = "Select a project"

	project_cmds_container := container.NewVBox()
	cmds_container := container.NewVBox()
	app_view.Middle = container.NewGridWithColumns(1, cmds_container)

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
	app_view.Top = container.NewVBox(top_top_section, top_bottom_section)

	cmds_menu_button := widget.NewButton(COMMANDS_MENU_LABEL, func() {
		if app_data.Selected_project == "" {
			outputWindow := app_view.App.NewWindow("Warning")
			output := widget.NewLabel("Please select a project in order to proceed.")
			outputWindow.SetContent(output)
			outputWindow.CenterOnScreen()
			outputWindow.Show()
			return
		}

		cmds_container.RemoveAll()
		project_cmds_container.RemoveAll()
		app_view.Middle.RemoveAll()

		cmds_section_label := widget.NewLabel(COMMANDS_SECTION_TITLE)
		cmds_section_divider := canvas.NewLine(color.White)

		p, ok := app_data.Projects[app_data.Selected_project]
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
					outputWindow := app_view.App.NewWindow("Result")
					output := widget.NewLabel("")
					outputWindow.SetContent(output)
					outputWindow.CenterOnScreen()
					outputWindow.Show()

					updateWidget(output, controller.RunCommand(cmd))
				})

				container := container.NewGridWithColumns(3, path_label, args_label, button)
				project_cmds_container.Add(container)
			}
		}

		cmds_container.Add(
			container.NewHBox(
				cmds_section_label,
				widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
				})))
		cmds_container.Add(cmds_section_divider)
		cmds_container.Add(project_cmds_container)
		app_view.Middle.Add(cmds_container)
		app_view.Middle.Refresh()
	})

	//envs_menu_button := widget.NewButton(ENVS_MENU_LABEL, func() {
	//	if app_data.Selected_project == "" {
	//		outputWindow := app_view.App.NewWindow("Warning")
	//		output := widget.NewLabel("Please select a project in order to proceed.")
	//		outputWindow.SetContent(output)
	//		outputWindow.CenterOnScreen()
	//		outputWindow.Show()
	//		return
	//	}
	//
	//	cmds_container.RemoveAll()
	//	cmds_container.RemoveAll()
	//	project_cmds_container.RemoveAll()
	//	app_view.Middle.RemoveAll()
	//
	//	envs_section_label := widget.NewLabel(COMMANDS_SECTION_TITLE)
	//	envs_section_divider := canvas.NewLine(color.White)
	//
	//	p, ok := app_data.Projects[app_data.Selected_project]
	//	if ok {
	//		for _, env := range p.Envs {
	//			key_label := widget.NewLabel(env.Key)
	//			value_label := widget.NewLabel(env.Value)
	//
	//			contentContainer := container.NewGridWithColumns(2, key_label, value_label)
	//			project_cmds_container.Add(contentContainer)
	//		}
	//	}
	//
	//	cmds_container.Add(envs_section_label)
	//	cmds_container.Add(envs_section_divider)
	//	cmds_container.Add(project_cmds_container)
	//	app_view.Middle.Add(cmds_container)
	//	app_view.Middle.Refresh()
	//})
	//app_view.Left = container.NewHBox(container.NewVBox(cmds_menu_button, envs_menu_button), canvas.NewLine(color.White))
	app_view.Left = container.NewHBox(container.NewVBox(cmds_menu_button), canvas.NewLine(color.White))

	version_label := widget.NewLabel(fmt.Sprintf("version %v", getCurrentVersion()))
	version_label.TextStyle.Italic = true
	app_view.Bottom = container.NewVBox(canvas.NewLine(color.White), container.NewHBox(version_label))

	content := container.NewBorder(app_view.Top, app_view.Bottom, app_view.Left, app_view.Right, app_view.Middle)

	app_view.Window.SetContent(content)
	app_view.Window.SetMaster()

	app_view.Window.Show()
	app_view.App.Run()
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
