package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"log"

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

	return container.NewGridWithColumns(4, title, container.NewStack(), container.NewStack(), container.NewVBox(container.NewHBox(projectAdd, projectEdit, projectDelete), projectSelect))
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
			g.contentView.Add(g.makeCommandsView())
		case ENV_VIEW:
			g.contentView.Add(g.makeEnvsView())
		}
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
	return container.NewVBox(commandViewButton, envViewButton)
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
	for _, cmd := range cmds {
		label := widget.NewLabel(fmt.Sprintf("%s %s", cmd.Path, cmd.Args))

		runButton := widget.NewButton(BUTTON_RUN_LABEL, func() {
			log.Println("Run command: ", cmd.Path, cmd.Args)
		})

		editButton := widget.NewButton(BUTTON_EDIT_LABEL, func() {
			log.Println("Edit command: ", cmd.Path, cmd.Args)
		})

		deleteButton := widget.NewButton(BUTTON_DELETE_LABEL, func() {
			log.Println("Delete command: ", cmd.Path, cmd.Args)
		})

		content.Add(container.NewHBox(label, runButton, editButton, deleteButton))
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
		label := widget.NewLabel(fmt.Sprintf("%s %s", env.Key, env.Value))

		editButton := widget.NewButton(BUTTON_EDIT_LABEL, func() {
			log.Println("Edit command: ", env.Key, env.Value)
		})

		deleteButton := widget.NewButton(BUTTON_DELETE_LABEL, func() {
			log.Println("Delete command: ", env.Key, env.Value)
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
