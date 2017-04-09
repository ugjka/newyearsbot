package main

import (
	"github.com/gotk3/gotk3/gtk"
)

//Status shows bot status
type Status struct {
	logStopper chan bool
	isOpen     bool
	onClose    func()
	stop       *gtk.Button
	window     *gtk.Window
	text       *gtk.TextView
	buffer     *gtk.TextBuffer
	scroll     *gtk.ScrolledWindow
	iter       *gtk.TextIter
}

//Open opens status window
func (w *Status) Open() {
	w.initWidgets()
	w.isOpen = true
}

//Close closes the status window
func (w *Status) Close() {
	w.window.Destroy()
	w.isOpen = false
}

func (w *Status) initWidgets() {
	var err error
	w.window, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	fatal(err)
	w.window.SetTitle("Bot Status")
	w.window.SetPosition(gtk.WIN_POS_CENTER)
	w.window.SetSizeRequest(400, 200)
	w.window.SetBorderWidth(6)
	_, err = w.window.Connect("destroy", w.onClose)
	fatal(err)
	w.stop, err = gtk.ButtonNew()
	fatal(err)
	w.stop.SetLabel("Stop")
	w.stop.Connect("clicked", func() {
		w.window.Destroy()
	})
	grid, err := gtk.GridNew()
	fatal(err)
	grid.SetColumnHomogeneous(true)
	grid.SetColumnSpacing(6)
	grid.SetRowSpacing(6)
	w.text, err = gtk.TextViewNew()
	fatal(err)
	w.buffer, err = w.text.GetBuffer()
	fatal(err)
	w.iter = w.buffer.GetEndIter()
	w.text.SetVExpand(true)
	w.text.SetEditable(false)
	w.text.SetCursorVisible(false)
	w.text.Connect("size-allocate", w.toEnd)
	w.scroll, err = gtk.ScrolledWindowNew(nil, nil)
	fatal(err)
	w.scroll.SetVExpand(true)
	w.scroll.Add(w.text)
	grid.Attach(w.scroll, 0, 0, 1, 1)
	grid.Attach(w.stop, 0, 1, 1, 1)
	w.window.Add(grid)
	w.window.ShowAll()
}

func (w *Status) toEnd() {
	w.text.ScrollToIter(w.buffer.GetEndIter(), 0, false, 0, 0)
}

func (w *Status) addMessage(msg string) bool {
	w.buffer.Insert(w.iter, msg)
	if !w.iter.IsEnd() {
		w.iter.ForwardToEnd()
	}
	return false
}
