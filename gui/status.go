package main

import (
	"github.com/gotk3/gotk3/gtk"
)

//Status shows bot status
type Status struct {
	isOpen  bool
	onClose func()
	stop    *gtk.Button
	window  *gtk.Window
	text    *gtk.TextView
}

//Open opens status window
func (w *Status) Open() {
	w.initWidgets()
	w.isOpen = true
}

//Close closes the status window
func (w *Status) Close() {
	w.window.Destroy()
}

func (w *Status) initWidgets() {
	var err error
	w.window, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	fatal(err)
	w.window.SetTitle("Bot Status")
	w.window.SetPosition(gtk.WIN_POS_CENTER)
	w.window.SetSizeRequest(200, 200)
	w.window.SetBorderWidth(6)
	w.stop, err = gtk.ButtonNew()
	fatal(err)
	w.stop.SetLabel("Stop")
	w.stop.Connect("clicked", func() {
		w.window.Destroy()
		mv.setActive()
	})
	grid, err := gtk.GridNew()
	fatal(err)
	grid.SetColumnHomogeneous(true)
	grid.SetColumnSpacing(6)
	grid.SetRowSpacing(6)
	w.text, err = gtk.TextViewNew()
	fatal(err)
	w.text.SetVExpand(true)
	w.text.SetEditable(false)
	grid.Attach(w.text, 0, 0, 1, 1)
	grid.Attach(w.stop, 0, 1, 1, 1)

	w.window.Add(grid)
	w.window.ShowAll()
}
