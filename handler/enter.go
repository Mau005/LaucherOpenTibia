package handler

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	conf "github.com/Mau005/LaucherOpenTibia/configuration"
	"github.com/Mau005/LaucherOpenTibia/controller"
)

type ViewEnterHandler struct{}

func (veh *ViewEnterHandler) CreateView() *fyne.Container {

	api := controller.ApiController{}
	image := canvas.NewImageFromFile(conf.API.PathLogo)
	image.SetMinSize(fyne.NewSize(0, 450))
	card := widget.NewCard(conf.API.NameApp, conf.API.Version, image)

	button := widget.NewButton(conf.API.NameButton, api.RunClient)

	return container.NewPadded(container.NewVBox(card, button))

}
