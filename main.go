package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	conf "github.com/Mau005/LaucherOpenTibia/configuration"
	"github.com/Mau005/LaucherOpenTibia/handler"
)

func main() {
	err := conf.Load("data/configuration.yml") //Initialized condifiguration
	if err != nil {
		log.Println(err)
		return
	}
	a := app.New()
	w := a.NewWindow(conf.API.TitleWindow)
	w.Resize(fyne.NewSize(450, 600))
	w.SetFixedSize(true)

	//Configuration Windows
	EnterView := handler.ViewEnterHandler{}
	w.CenterOnScreen()
	w.SetContent(EnterView.CreateView())
	w.ShowAndRun()
}
