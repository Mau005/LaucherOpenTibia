package controller

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Paleta inspirada en el mock MU (valores aproximados)
var (
	// Fondo azul-verdoso con toque “mist”
	clrBackground = color.NRGBA{R: 18, G: 26, B: 38, A: 255}
	// Paneles metálico-azulados
	clrCard = color.NRGBA{R: 31, G: 43, B: 59, A: 255}
	// Borde dorado suave
	clrGold = color.NRGBA{R: 199, G: 162, B: 77, A: 255}
	// Texto crema
	clrText = color.NRGBA{R: 241, G: 219, B: 134, A: 255}
	// Primario (botón START azul brillante)
	clrPrimary = color.NRGBA{R: 20, G: 122, B: 196, A: 255}
	// Hover primario (más claro)
	clrPrimary2 = color.NRGBA{R: 35, G: 150, B: 225, A: 255}
	// Secundario (botones Website/Forum)
	clrSecondary = color.NRGBA{R: 53, G: 69, B: 92, A: 255}
	// Campos de entrada
	clrInputBG = color.NRGBA{R: 24, G: 32, B: 45, A: 255}
	// Resaltado/Focus
	clrFocus = color.NRGBA{R: 96, G: 174, B: 255, A: 255}
)

// muTheme implementa fyne.Theme para colores y tamaños
type muTheme struct{ fyne.Theme }

func NewMUTheme() fyne.Theme { return &muTheme{Theme: theme.DarkTheme()} }

// Colores clave para widgets
func (m *muTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return clrBackground
	// case theme.ColorNameCardBackground:
	// 	return clrCard
	case theme.ColorNameButton:
		return clrSecondary
	case theme.ColorNamePrimary:
		return clrPrimary
	// case theme.ColorNamePrimaryHover:
	// 	return clrPrimary2
	case theme.ColorNameForeground:
		return clrText
	case theme.ColorNameShadow:
		return color.NRGBA{0, 0, 0, 160}
	case theme.ColorNameFocus:
		return clrFocus
	case theme.ColorNameInputBackground:
		return clrInputBG
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 170, G: 170, B: 170, A: 255}
	case theme.ColorNameSeparator:
		return clrGold
	default:
		return theme.DarkTheme().Color(name, theme.VariantDark)
	}
}

// Tamaños para una UI “más densa”
func (m *muTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameInnerPadding:
		return 10
	case theme.SizeNameText:
		return 16
	case theme.SizeNameCaptionText:
		return 13
	case theme.SizeNameInputBorder:
		return 2
	// case theme.SizeNameScrollbar:
	// 	return 8
	default:
		return theme.DarkTheme().Size(name)
	}
}

// Conserva fuentes e íconos por defecto (puedes cambiarlo si quieres tipografía custom)
func (m *muTheme) Font(style fyne.TextStyle) fyne.Resource    { return theme.DarkTheme().Font(style) }
func (m *muTheme) Icon(name fyne.ThemeIconName) fyne.Resource { return theme.DarkTheme().Icon(name) }
