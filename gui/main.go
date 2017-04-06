//GTK3 implementation of newyearsbot
package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	var mv Window
	mv.onClose = func() {
		gtk.MainQuit()
	}
	gtk.Init(nil)
	mv.open()
	gtk.Main()
}

//Window contains top level window
type Window struct {
	ircChannels []string
	ircServer   string
	ircUseTLS   bool
	ircNick     string

	onClose func()

	isOpen bool

	window *gtk.Window
	chans  *gtk.Entry
	server *gtk.Entry
	nick   *gtk.Entry
	tls    *gtk.CheckButton
	start  *gtk.Button
	stop   *gtk.Button
}

func (w *Window) open() {
	if w.isOpen {
		return
	}
	w.initWidgets()
	w.isOpen = true
}

func (w *Window) close() {
	if !w.isOpen {
		return
	}
	w.window.Destroy()
}

func (w *Window) initWidgets() {
	var err error
	w.window, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	fatal(err)
	w.window.SetTitle("New Year Irc Party Bot")
	w.window.SetPosition(gtk.WIN_POS_CENTER)
	w.window.SetSizeRequest(640, 320)
	w.window.SetBorderWidth(6)
	_, err = w.window.Connect("destroy", w.windowDestroyed)
	fatal(err)
	grid, err := gtk.GridNew()
	fatal(err)
	grid.SetColumnHomogeneous(true)
	grid.SetColumnSpacing(6)
	grid.SetRowSpacing(6)
	w.window.Add(grid)
	config, err := gtk.FrameNew("Configuration:")
	fatal(err)
	config.SetBorderWidth(6)
	grid.Attach(config, 0, 0, 1, 1)
	grid2, err := gtk.GridNew()
	fatal(err)
	grid2.SetColumnHomogeneous(true)
	grid2.SetColumnSpacing(6)
	grid2.SetRowSpacing(6)
	grid2.SetBorderWidth(6)
	grid2.Attach(labelNew("Irc nick:"), 0, 0, 1, 1)
	w.nick, err = gtk.EntryNew()
	grid2.Attach(w.nick, 0, 1, 1, 1)
	grid2.Attach(labelNew("Irc channels (comma seperated):"), 0, 2, 1, 1)
	w.chans, err = gtk.EntryNew()
	fatal(err)
	grid2.Attach(w.chans, 0, 3, 1, 1)
	grid2.Attach(labelNew("Irc server (host:port):"), 0, 4, 1, 1)
	w.server, err = gtk.EntryNew()
	fatal(err)
	w.server.SetText("irc.freenode.net:7000")
	grid2.Attach(w.server, 0, 5, 1, 1)
	grid2.Attach(labelNew("Use TLS:"), 0, 6, 1, 1)
	w.tls, err = gtk.CheckButtonNew()
	fatal(err)
	w.tls.SetActive(true)
	w.tls.SetHAlign(gtk.ALIGN_END)
	grid2.Attach(w.tls, 0, 6, 1, 1)
	config.Add(grid2)
	w.start, err = gtk.ButtonNew()
	fatal(err)
	w.start.SetLabel("Start")
	w.start.SetHAlign(gtk.ALIGN_CENTER)
	w.start.Connect("clicked", w.startClicked)
	grid.Attach(w.start, 0, 1, 1, 2)
	w.window.ShowAll()
}

func (w *Window) startClicked() {
	msg := gtk.MessageDialogNew(w.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, "%s", "I gift you an error")
	_, err := msg.Connect("response", func() {
		msg.Destroy()
	})
	fatal(err)
	msg.ShowAll()
}

func (w *Window) windowDestroyed() {
	w.onClose()
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func labelNew(s string) *gtk.Label {
	l, err := gtk.LabelNew(s)
	fatal(err)
	l.SetHAlign(gtk.ALIGN_START)
	return l
}
