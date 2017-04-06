package main

import (
	"github.com/gotk3/gotk3/gtk"
)

//ErrorWin display error message
type ErrorWin struct {
	onClose func()
	isOpen  bool

	dialog *gtk.Dialog
}

func (w *ErrorWin) initWidgets() {
	var err error
	w.dialog, err = gtk.DialogNew()
	fatal(err)
	w.dialog.AddButton("close", gtk.RESPONSE_CLOSE)
}
