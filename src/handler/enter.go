package handler

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Mau005/LaucherOpenTibia/src/configuration"
	"github.com/Mau005/LaucherOpenTibia/src/controller"
)

type ViewEnterHandler struct{}

func (veh *ViewEnterHandler) CreateView() *fyne.Container {

	api := controller.ApiController{}
	image := canvas.NewImageFromFile(configuration.API.PathLogo)
	image.SetMinSize(fyne.NewSize(0, 450))
	card := widget.NewCard(configuration.API.NameApp, configuration.API.Version, image)

	button := widget.NewButton(configuration.API.NameButton, api.RunClient)

	return container.NewPadded(container.NewVBox(card, button))

}
