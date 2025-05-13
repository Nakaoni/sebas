package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	SIDE_WIDTH = 160
)

// dividers are 3 objects in this order: top, left, bottom
type SebasLayout struct {
	top, bottom, left, content fyne.CanvasObject
	dividers                   [3]fyne.CanvasObject
}

func newSebasLayout(top, bottom, left, content fyne.CanvasObject, dividers [3]fyne.CanvasObject) fyne.Layout {
	return &SebasLayout{
		top,
		bottom,
		left,
		content,
		dividers,
	}
}

func (sebasLayout *SebasLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	topHeight := sebasLayout.top.MinSize().Height
	sebasLayout.top.Resize(fyne.NewSize(size.Width, topHeight))

	sebasLayout.left.Move(fyne.NewPos(0, topHeight))
	sebasLayout.left.Resize(fyne.NewSize(SIDE_WIDTH, size.Height-topHeight))

	bottomHeight := sebasLayout.bottom.MinSize().Height
	sebasLayout.bottom.Move(fyne.NewPos(SIDE_WIDTH, size.Height-bottomHeight))
	sebasLayout.bottom.Resize(fyne.NewSize(size.Width-SIDE_WIDTH, bottomHeight))

	dividerThickness := theme.SeparatorThicknessSize()
	// Top divider
	sebasLayout.dividers[0].Move(fyne.NewPos(0, topHeight))
	sebasLayout.dividers[0].Resize(fyne.NewSize(size.Width, dividerThickness))

	// Left divider
	sebasLayout.dividers[1].Move(fyne.NewPos(sebasLayout.left.Size().Width+dividerThickness, topHeight))
	sebasLayout.dividers[1].Resize(fyne.NewSize(1, size.Height-topHeight))

	// Bottom divider
	sebasLayout.dividers[2].Move(fyne.NewPos(sebasLayout.left.Size().Width, size.Height-bottomHeight-dividerThickness))
	sebasLayout.dividers[2].Resize(fyne.NewSize(size.Width, dividerThickness))

	// Content
	sebasLayout.content.Move(fyne.NewPos(SIDE_WIDTH, topHeight))
	sebasLayout.content.Resize(fyne.NewSize(size.Width-SIDE_WIDTH, size.Height-topHeight-bottomHeight))
}

func (sebasLayout *SebasLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(
		sebasLayout.left.MinSize().Width+sebasLayout.content.MinSize().Width,
		sebasLayout.top.MinSize().Height+sebasLayout.content.MinSize().Height+sebasLayout.bottom.MinSize().Height,
	)
}
