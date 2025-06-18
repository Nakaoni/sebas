//go:generate fyne bundle -o bundled.go assets

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
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

func (sebasTheme *SebasTheme) Color(c fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return sebasTheme.Theme.Color(c, theme.VariantDark)
}
