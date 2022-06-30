package glfw

import (
	"fyne.io/fyne/v2"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func (d *gLDriver) CreateCustomWindow(title string, decorate, transparent, onTop, centered bool) fyne.Window {
	win := d.createCustomWindow(title, decorate, transparent, onTop)
	if centered {
		win.SetPadded(false)
		win.CenterOnScreen()
	}
	return win
}

func (d *gLDriver) createCustomWindow(title string, decorate, transparent, onTop bool) fyne.Window {
	var ret *window
	if title == "" {
		title = defaultTitle
	}
	runOnMain(func() {
		d.initGLFW()

		ret = &window{title: title, decorate: decorate, transparent: transparent, onTop: onTop, driver: d}
		// This queue is destroyed when the window is closed.
		ret.InitEventQueue()
		go ret.RunEventQueue()

		ret.canvas = newCanvas()
		ret.canvas.context = ret
		ret.SetIcon(ret.icon)
		d.addWindow(ret)
	})
	return ret
}

func (w *window) GetGlfwCustomWindow() *glfw.Window {
	return w.view()
}

func (w *window) GetGlfwCustomMonitor() *glfw.Monitor {
	return w.getMonitorForWindow()
}

func (w *window) CreateCustomGlfwWindow() {
	w.createLock.Do(w.create)
}
