package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type transparentEntryTheme struct {
	base fyne.Theme
}

func (t transparentEntryTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if n == theme.ColorNameInputBackground {
		return color.NRGBA{0, 0, 0, 160}
	}
	return t.base.Color(n, v)
}

func (t transparentEntryTheme) Icon(n fyne.ThemeIconName) fyne.Resource { return t.base.Icon(n) }
func (t transparentEntryTheme) Font(s fyne.TextStyle) fyne.Resource     { return t.base.Font(s) }
func (t transparentEntryTheme) Size(n fyne.ThemeSizeName) float32       { return t.base.Size(n) }

func LaucherUI(containerInternal *fyne.Container) *fyne.Container {
	wallpaper := canvas.NewImageFromFile("./assets/wallpaper.png")
	wallpaper.FillMode = canvas.ImageFillStretch
	cont := container.New(layout.NewStackLayout(), wallpaper, containerInternal)

	return cont
}

func ContentUI() *fyne.Container {
	loading := widget.NewProgressBarInfinite()
	containerButton := container.NewHBox(
		widget.NewButton("Website", nil),
		widget.NewButton("Forum", nil),
		widget.NewButton("Start", nil),
	)

	cont := container.NewVBox()

	for i := 0; i < 40; i++ {
		but := widget.NewButton(fmt.Sprintf("Number: %d", i), nil)
		cont.Add(but)
	}
	border := container.NewBorder(nil, container.NewCenter(containerButton), nil, nil, container.NewBorder(
		container.NewCenter(widget.NewLabel("Welcome to AinhoOT")),
		loading,
		nil,
		nil,
		container.NewAdaptiveGrid(2,
			container.NewVScroll(cont),
			widget.NewMultiLineEntry(),
		),
	))
	return container.NewPadded(border)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(transparentEntryTheme{base: theme.DarkTheme()})
	w := a.NewWindow("Laucher AinhoOT")
	w.Resize(fyne.NewSize(800, 600))

	InitLaucher := LaucherUI(ContentUI())
	w.SetContent(InitLaucher)
	w.ShowAndRun()
}
