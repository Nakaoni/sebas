package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2/canvas"
	"github.com/e-felix/sebas/cmd/desktop/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/e-felix/sebas/internal/project"
)

const (
	TOP_BAR_TITLE              = "Sebas - Your Personal Butler"
	PROJECT_SELECT_PLACEHOLDER = "Select a project"

	// Menu
	FILE_MENU      = "File"
	FILE_MENU_OPEN = "Open File..."

	// Project
	PROJECT_ADD_BUTTON    = "Add new"
	PROJECT_EDIT_BUTTON   = "Edit"
	PROJECT_DELETE_BUTTON = "Delete"

	// Views
	COMMAND_VIEW              = "command-view"
	COMMAND_VIEW_TITLE        = "List of Commands"
	COMMAND_VIEW_BUTTON_LABEL = "Commands"
	ENV_VIEW                  = "env-view"
	ENV_VIEW_TITLE            = "List of Envs"
	ENV_VIEW_BUTTON_LABEL     = "Envs"

	// Button
	BUTTON_RUN_LABEL    = "Run"
	BUTTON_EDIT_LABEL   = "Edit"
	BUTTON_SAVE_LABEL   = "Save"
	BUTTON_DELETE_LABEL = "Delete"
)

type gui struct {
	win            fyne.Window
	file           binding.String
	data           map[string]*project.Project
	currentProject *project.Project
	currentView    binding.String
	contentView    *fyne.Container
}

func (g *gui) makeUi() fyne.CanvasObject {
	top := g.makeTopBar()
	footer := widget.NewLabel("Footer")
	leftMenu := g.makeLeftMenu()

	g.contentView = container.NewStack()

	topDivider := widget.NewSeparator()
	leftDivider := widget.NewSeparator()
	bottomDivider := widget.NewSeparator()
	dividers := [3]fyne.CanvasObject{
		topDivider,
		leftDivider,
		bottomDivider,
	}

	g.setUpViewListener()

	objs := []fyne.CanvasObject{top, footer, leftMenu, g.contentView, dividers[0], dividers[1], dividers[2]}
	return container.New(newSebasLayout(top, footer, leftMenu, g.contentView, dividers), objs...)
}

func (g *gui) makeTopBar() fyne.CanvasObject {
	title := widget.NewLabel(TOP_BAR_TITLE)

	projectsNameList := make([]string, 0)
	for k := range g.data {
		projectsNameList = append(projectsNameList, k)
	}

	projectSelect := widget.NewSelect(projectsNameList, func(selected_value string) {
		if g.data[selected_value] == nil {
			dialog.ShowError(errors.New(fmt.Sprintf("Project %s not found\n", selected_value)), g.win)
			return
		}

		g.currentProject = g.data[selected_value]
	})
	projectSelect.PlaceHolder = PROJECT_SELECT_PLACEHOLDER

	projectAdd := widget.NewButtonWithIcon(PROJECT_ADD_BUTTON, theme.Icon(theme.IconNameContentAdd), func() {})
	projectEdit := widget.NewButtonWithIcon(PROJECT_EDIT_BUTTON, theme.Icon(theme.IconNameDocumentCreate), func() {})
	projectDelete := widget.NewButtonWithIcon(PROJECT_DELETE_BUTTON, theme.Icon(theme.IconNameDelete), func() {})

	return container.NewPadded(container.NewBorder(nil, nil, title, container.NewHBox(projectAdd, projectEdit, projectDelete, projectSelect)))
}

func (g *gui) setUpViewListener() {
	g.currentView.AddListener(binding.NewDataListener(func() {
		if g.currentProject == nil {
			return
		}

		view, err := g.currentView.Get()
		if err != nil {
			return
		}

		g.contentView.RemoveAll()

		switch view {
		case COMMAND_VIEW:
			g.contentView.Add(container.NewPadded(g.makeCommandsView()))
		case ENV_VIEW:
			g.contentView.Add(g.makeEnvsView())
		}

		g.contentView.Refresh()
	}))
}

func (g *gui) makeLeftMenu() fyne.CanvasObject {
	commandViewButton := widget.NewButton(COMMAND_VIEW_BUTTON_LABEL, func() {
		err := g.currentView.Set(COMMAND_VIEW)
		if err != nil {
			return
		}
	})
	envViewButton := widget.NewButton(ENV_VIEW_BUTTON_LABEL, func() {
		err := g.currentView.Set(ENV_VIEW)
		if err != nil {
			return
		}
	})
	return container.NewPadded(container.NewVBox(commandViewButton, envViewButton))
}

func (g *gui) makeCommandsView() fyne.CanvasObject {
	content := container.NewVBox()

	title := canvas.NewText(COMMAND_VIEW_TITLE, theme.Color(theme.ColorNameForeground))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	content.Add(title)
	content.Add(widget.NewSeparator())

	cmds := g.currentProject.Cmds
	for i, cmd := range cmds {
		currentCmd := cmd

		currentCmdIndex := i

		currentCmdValue := binding.NewString()
		_ = currentCmdValue.Set(fmt.Sprintf("%s", currentCmd.Path))

		currentArgsValue := binding.NewString()
		_ = currentArgsValue.Set(fmt.Sprintf("%s", strings.Join(currentCmd.Args, " ")))

		inputCmd := widget.NewEntryWithData(currentCmdValue)
		inputArgs := widget.NewEntryWithData(currentArgsValue)

		inputCmd.Disable()
		inputArgs.Disable()

		runButton := widget.NewButton(BUTTON_RUN_LABEL, func() {
			log.Println(controller.RunCommand(currentCmd))
		})

		var editButton *widget.Button
		var saveButton *widget.Button

		editButton = widget.NewButton(BUTTON_EDIT_LABEL, func() {
			inputCmd.Enable()
			inputArgs.Enable()

			saveButton.Show()
			editButton.Hide()
		})

		saveButton = widget.NewButton(BUTTON_SAVE_LABEL, func() {
			currentCmd.Path = inputCmd.Text
			currentCmd.Args = strings.Split(inputArgs.Text, " ")

			cmds[currentCmdIndex].Path = currentCmd.Path
			cmds[currentCmdIndex].Args = currentCmd.Args

			err := controller.EditCommand(*g.currentProject, cmds[currentCmdIndex])
			if err != nil {
				log.Println(err)
			}
			
			log.Println(g.data, cmds)

			inputCmd.Disable()
			inputArgs.Disable()

			saveButton.Hide()
			editButton.Show()
		})

		saveButton.Hide()
		saveAndEditStack := container.NewStack(editButton, saveButton)

		deleteButton := widget.NewButton(BUTTON_DELETE_LABEL, func() {
			log.Println("Delete command: ", currentCmd.Path, currentCmd.Args)

			cmds = append(cmds[:currentCmdIndex], cmds[currentCmdIndex+1:]...)
			g.currentProject.Cmds = cmds
			g.setUpViewListener()
		})

		content.Add(container.NewAdaptiveGrid(3, inputCmd, inputArgs, container.NewHBox(runButton, saveAndEditStack, deleteButton)))
	}

	return content
}

func (g *gui) makeEnvsView() fyne.CanvasObject {
	content := container.NewVBox()

	title := canvas.NewText(ENV_VIEW_TITLE, theme.Color(theme.ColorNameForeground))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	content.Add(title)
	content.Add(widget.NewSeparator())

	envs := g.currentProject.Envs
	for _, env := range envs {
		currentEnv := env
		label := widget.NewLabel(fmt.Sprintf("%s %s", currentEnv.Key, currentEnv.Value))

		editButton := widget.NewButton(BUTTON_EDIT_LABEL, func() {
			log.Println("Edit command: ", currentEnv.Key, currentEnv.Value)
		})

		deleteButton := widget.NewButton(BUTTON_DELETE_LABEL, func() {
			log.Println("Delete command: ", currentEnv.Key, currentEnv.Value)
		})

		content.Add(container.NewHBox(label, editButton, deleteButton))
	}

	return content
}

func (g *gui) makeMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu(
			FILE_MENU,
			fyne.NewMenuItem(
				FILE_MENU_OPEN,
				func() {
					g.openFileDialog()
				},
			),
		),
	)
}

func (g *gui) openFileDialog() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, g.win)
			return
		}

		if reader == nil {
			return
		}

		g.openFile(reader)
	}, g.win)

}

func (g *gui) openFile(reader fyne.URIReadCloser) {
	filename := reader.URI().Name()

	err := g.file.Set(filename)

	if err != nil {
		dialog.ShowError(err, g.win)
		return
	}
}
