package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
)

type gui struct {
	win  fyne.Window
	file binding.String
}

func (g *gui) makeTopBar() fyne.CanvasObject {
	title := widget.NewLabel(TOP_BAR_TITLE)

	projectSelect := widget.NewSelect([]string{"Project 1", "Project 2", "Project 3"}, func(selected_value string) {})
	projectSelect.PlaceHolder = PROJECT_SELECT_PLACEHOLDER

	projectAdd := widget.NewButtonWithIcon(PROJECT_ADD_BUTTON, theme.Icon(theme.IconNameContentAdd), func() {})
	projectEdit := widget.NewButtonWithIcon(PROJECT_EDIT_BUTTON, theme.Icon(theme.IconNameDocumentCreate), func() {})
	projectDelete := widget.NewButtonWithIcon(PROJECT_DELETE_BUTTON, theme.Icon(theme.IconNameDelete), func() {})

	return container.NewGridWithColumns(2, title, container.NewHBox(projectAdd, projectEdit, projectDelete, projectSelect))
}

func (g *gui) makeUi() fyne.CanvasObject {
	top := g.makeTopBar()
	footer := widget.NewLabel("Footer")
	leftMenu := widget.NewLabel("Left Menu")

	content := widget.NewLabelWithData(g.file)

	topDivider := widget.NewSeparator()
	leftDivider := widget.NewSeparator()
	bottomDivider := widget.NewSeparator()
	dividers := [3]fyne.CanvasObject{
		topDivider,
		leftDivider,
		bottomDivider,
	}

	objs := []fyne.CanvasObject{top, footer, leftMenu, content, dividers[0], dividers[1], dividers[2]}
	return container.New(newSebasLayout(top, footer, leftMenu, content, dividers), objs...)
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
