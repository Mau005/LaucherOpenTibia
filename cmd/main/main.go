package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Mau005/LaucherOpenTibia/src/configuration"
	"github.com/Mau005/LaucherOpenTibia/src/handler"
)

func main() {
	err := configuration.Load("data/configuration.yml") //Initialized condifiguration
	if err != nil {
		log.Println(err)
		return
	}
	a := app.New()
	w := a.NewWindow(configuration.API.TitleWindow)
	w.Resize(fyne.NewSize(450, 600))
	w.SetFixedSize(true)
	//Configuration Windows
	EnterView := handler.ViewEnterHandler{}
	w.CenterOnScreen()
	w.SetContent(EnterView.CreateView())
	w.ShowAndRun()
}
