package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	TOP_BAR_TITLE              = "Sebas - Your Personal Butler"
	PROJECT_SELECT_PLACEHOLDER = "Select a project"
)

type gui struct {
	win fyne.Window
}

func makeTopBar() fyne.CanvasObject {
	title := widget.NewLabel(TOP_BAR_TITLE)

	projectSelect := widget.NewSelect([]string{"Project 1", "Project 2", "Project 3"}, func(selected_value string) {})
	projectSelect.PlaceHolder = PROJECT_SELECT_PLACEHOLDER

	return container.NewHBox(title, projectSelect)
}

func (g *gui) makeUi() fyne.CanvasObject {
	top := makeTopBar()
	footer := widget.NewLabel("Footer")
	leftMenu := widget.NewLabel("Left Menu")
	content := widget.NewLabel("Content")

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

func (g *gui) openDatabase() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {}, g.win)
}
