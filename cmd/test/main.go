package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Layout 3 columnas por proporción (ej: 0.1, 0.8, 0.1)
type ThreeRatioLayout struct {
	R1, R2, R3 float64 // deben sumar 1.0 (o cercano)
}

func (l *ThreeRatioLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	// Esperamos exactamente 3 objetos: img, loading, botón
	if len(objs) == 0 {
		return
	}
	// Asegura índices seguros
	var img, load, btn fyne.CanvasObject
	if len(objs) > 0 {
		img = objs[0]
	}
	if len(objs) > 1 {
		load = objs[1]
	}
	if len(objs) > 2 {
		btn = objs[2]
	}

	w1 := float32(l.R1) * size.Width
	w2 := float32(l.R2) * size.Width
	w3 := size.Width - w1 - w2

	x := float32(0)
	if img != nil {
		img.Resize(fyne.NewSize(w1, size.Height))
		img.Move(fyne.NewPos(x, 0))
		x += w1
	}
	if load != nil {
		load.Resize(fyne.NewSize(w2, size.Height))
		load.Move(fyne.NewPos(x, 0))
		x += w2
	}
	if btn != nil {
		btn.Resize(fyne.NewSize(w3, size.Height))
		btn.Move(fyne.NewPos(x, 0))
	}
}

func (l *ThreeRatioLayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	// El mínimo: suma de mínimos en ancho y la altura máxima
	var w, h float32
	for _, o := range objs {
		if o == nil {
			continue
		}
		ms := o.MinSize()
		w += ms.Width
		if ms.Height > h {
			h = ms.Height
		}
	}
	return fyne.NewSize(w, h)
}

// Componente: 10% imagen, 80% loading, 10% botón
func LoadingRowWithImageAndButton(imgPath string, buttonText string) *fyne.Container {
	// 10%: imagen (puedes ajustar FillMode y MinSize)
	img := canvas.NewImageFromFile(imgPath)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(50, 50)) // mínimo razonable

	// 80%: loading (progress bar). Infinite si quieres “cargando…”
	// 80% loading (alto aumentado)
	p := widget.NewProgressBarInfinite()
	p.Start()

	centerLoad := container.NewStack(p) // que ocupe todo el alto disponible

	// (Opcional) Si prefieres una barra normal:
	// pb := widget.NewProgressBar()
	// go func(){ for i:=0.0;i<=1.0;i+=0.01{ time.Sleep(30*time.Millisecond); pb.SetValue(i)} }()

	// 10%: botón
	btn := widget.NewButton(buttonText, func() {
		// acción del botón
	})

	// Alto del renglón: si quieres algo más “compacto”, puedes envolver cada uno con Center
	centerImg := container.NewCenter(img)
	centerBtn := container.NewCenter(btn)

	return container.New(&ThreeRatioLayout{R1: 0.10, R2: 0.80, R3: 0.10}, centerImg, centerLoad, centerBtn)
}

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
	wallpaper := canvas.NewImageFromFile("./assets/background2.png")
	wallpaper.FillMode = canvas.ImageFillStretch
	cont := container.New(layout.NewStackLayout(), wallpaper, containerInternal)

	return cont
}

func NewsShort(iconID uint8, newsNotice string) *fyne.Container {
	image := canvas.NewImageFromFile("./assets/newsicon_community_small.png")
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(50, 50))
	// entry := widget.NewMultiLineEntry()
	// entry.SetText(newsNotice)
	// entry.Disable()
	entry := widget.NewRichTextFromMarkdown(newsNotice)
	split := container.NewHSplit(image, entry)
	split.SetOffset(0.15)
	return container.NewPadded(split)
}

func ContentUI() *fyne.Container {
	containerButton := container.NewHBox(
		widget.NewButton("Website", nil),
		widget.NewButton("Forum", nil),
		widget.NewButton("Start", nil),
	)

	cont := container.NewVBox()

	for i := 0; i < 30; i++ {
		but := NewsShort(1, `

# This project is being streamed in development on
- https://www.youtube.com/maugame
- https://www.twitch.tv/kraynodev

## Donate
If you enjoy the project and like the work that has been done, feel free to donate. Donations are not necessary, but anything is appreciated. Thank you!

Paypal: https://paypal.me/Mau2?country.x=CL&locale.x=es_XC `)
		cont.Add(but)
	}
	title := widget.NewRichTextFromMarkdown("# Welcome To AinhoOT")
	title.Theme().Font(fyne.TextStyle{Bold: true})
	border := container.NewBorder(nil, container.NewCenter(containerButton), nil, nil, container.NewBorder(
		container.NewCenter(title),
		LoadingRowWithImageAndButton("./assets/icon.png", "Start"),
		nil,
		nil,
		container.NewAdaptiveGrid(2,
			container.NewVScroll(cont),
			canvas.NewImageFromFile("./assets/wallpaper2.png"),
		),
	))
	return container.NewPadded(border)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(transparentEntryTheme{base: theme.DarkTheme()})
	w := a.NewWindow("Laucher AinhoOT")
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(false)

	InitLaucher := LaucherUI(ContentUI())
	w.SetContent(InitLaucher)
	w.ShowAndRun()
}
