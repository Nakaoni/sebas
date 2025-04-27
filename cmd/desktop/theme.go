package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type SebasTheme struct {
	fyne.Theme
}

func newSebasTheme() fyne.Theme {
	return &SebasTheme{
		Theme: theme.DefaultTheme(),
	}
}

func (sebasTheme *SebasTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return sebasTheme.Theme.Size(name)
}
