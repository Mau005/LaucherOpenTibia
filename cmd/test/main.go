package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Mau005/LaucherOpenTibia/src/controller"
)

func main() {
	a := app.NewWithID("com.ainhoot.mu.theme")
	a.Settings().SetTheme(controller.NewMUTheme())

	w := a.NewWindow("AinhoOT Launcher — Theme Demo")
	w.Resize(fyne.NewSize(1000, 640))

	// Header “WELCOME” estilo crema/dorado
	title := widget.NewLabel("WELCOME TO  AinhoOT")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Simulamos panel central tipo “card”
	patch := widget.NewMultiLineEntry()
	patch.SetText("Patch Notes / News\n- Title of patch version 2  (00.02.00)\n- Title of patch          (00.01.00)")
	patch.Disable()

	server := widget.NewLabel("Server: Offline")
	progress := widget.NewProgressBar()

	btnWebsite := widget.NewButton("WEBSITE", func() {})
	btnForum := widget.NewButton("FORUM", func() {})
	btnStart := widget.NewButtonWithIcon("START", nil, func() {})
	btnStart.Importance = widget.HighImportance // usa color Primario (azul brillante)

	// Layout: columna central con “panel” y botones abajo
	center := container.NewVBox(
		title,
		container.NewVBox(
			// borde superior “dorado” simulando marco
			widget.NewSeparator(),
			patch,
			widget.NewSeparator(),
			progress,
			server,
		),
		container.NewHBox(
			container.NewMax(widget.NewButton("   ", nil)), // separador flexible si lo deseas
		),
		container.NewHBox(btnWebsite, btnForum, btnStart),
	)

	w.SetContent(container.NewPadded(center))
	w.ShowAndRun()
}
