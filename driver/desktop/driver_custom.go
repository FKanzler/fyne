package desktop

import (
	"fyne.io/fyne/v2"
)

type DriverCustom interface {
	CreateCustomWindow(title string, decorate, transparent, onTop, centered bool) fyne.Window
}
