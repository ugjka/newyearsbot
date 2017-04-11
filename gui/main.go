//
//GTK3 implementation of newyearsbot
//
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ugjka/newyearsbot/nyb"
)

func main() {
	var st Status
	st.logStopper = make(chan bool)
	var mv Window
	mv.ircServer = "irc.freenode.net:7000"
	mv.ircUseTLS = true
	mv.ircTrigger = "hny"
	bot := &nyb.Settings{}

	st.onClose = func() {
		st.logStopper <- true
		if !bot.Exited {
			bot.Stop()
		}
		st.Close()
		mv.setActive()
	}

	mv.startBot = func() {
		bot = &nyb.Settings{}
		bot.IrcChans = mv.ircChannels
		bot.IrcNick = mv.ircNick
		bot.IrcTrigger = mv.ircTrigger
		bot.IrcServer = mv.ircServer
		bot.UseTLS = mv.ircUseTLS
		bot.Stopper = make(chan bool)
		bot.LogCh = nyb.NewLogChan()
		bot.IrcObj = nyb.NewIrcObj()
		go bot.Start()
	}
	mv.onClose = func() {
		gtk.MainQuit()
	}
	mv.onHide = func() {
		if st.isOpen {
			return
		}
		mv.setInactive()
		st.Open()
		mv.startBot()
		go func(s *Status) {
			for {
				var logmsg string
				select {
				case <-s.logStopper:
					return
				case logmsg = <-bot.LogCh:
					_, err := glib.IdleAdd(s.addMessage, logmsg)
					fatal(err)
				}
			}
		}(&st)
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
	ircTrigger  string

	onClose  func()
	onHide   func()
	startBot func()

	isOpen bool

	window  *gtk.Window
	chans   *gtk.Entry
	server  *gtk.Entry
	nick    *gtk.Entry
	trigger *gtk.Entry
	tls     *gtk.CheckButton
	start   *gtk.Button
	stop    *gtk.Button
}

func (w *Window) open() {
	if w.isOpen {
		return
	}
	w.initWidgets()
	w.fillInputs()
	w.isOpen = true
}

func (w *Window) close() {
	if !w.isOpen {
		return
	}
	w.window.Destroy()
	w.isOpen = false
}

func (w *Window) fillInputs() {
	w.nick.SetText(w.ircNick)
	chans := ""
	for i, ch := range w.ircChannels {
		chans += ch
		if i != len(w.ircChannels)-1 {
			chans += ", "
		}
	}
	w.chans.SetText(chans)
	w.trigger.SetText(w.ircTrigger)
	w.server.SetText(w.ircServer)
	w.tls.SetActive(w.ircUseTLS)
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
	grid2.Attach(labelNew("Bot trigger for queries:"), 0, 2, 1, 1)
	w.trigger, err = gtk.EntryNew()
	fatal(err)
	grid2.Attach(w.trigger, 0, 3, 1, 1)
	grid2.Attach(labelNew("Irc channels (comma seperated):"), 0, 4, 1, 1)
	w.chans, err = gtk.EntryNew()
	fatal(err)
	grid2.Attach(w.chans, 0, 5, 1, 1)
	grid2.Attach(labelNew("Irc server (host:port):"), 0, 6, 1, 1)
	w.server, err = gtk.EntryNew()
	fatal(err)
	w.server.SetText(w.ircServer)
	grid2.Attach(w.server, 0, 7, 1, 1)
	grid2.Attach(labelNew("Use TLS:"), 0, 8, 1, 1)
	w.tls, err = gtk.CheckButtonNew()
	fatal(err)
	w.tls.SetActive(w.ircUseTLS)
	w.tls.SetHAlign(gtk.ALIGN_END)
	grid2.Attach(w.tls, 0, 9, 1, 1)
	config.Add(grid2)
	w.start, err = gtk.ButtonNew()
	fatal(err)
	w.start.SetLabel("Start")
	w.start.SetHAlign(gtk.ALIGN_CENTER)
	w.start.Connect("clicked", w.startClicked)
	grid.Attach(w.start, 0, 1, 1, 2)
	w.window.ShowAll()
}

func (w *Window) setInactive() {
	w.window.SetVisible(false)
}

func (w *Window) setActive() {
	w.window.SetVisible(true)
}

func (w *Window) startClicked() {
	if err := w.validateInputs(); err != nil {
		msg := gtk.MessageDialogNew(w.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR,
			gtk.BUTTONS_CLOSE, "%s", err.Error())
		_, err := msg.Connect("response", func() {
			msg.Destroy()
		})
		fatal(err)
		msg.ShowAll()
	} else {
		var err error
		w.ircNick, err = w.nick.GetText()
		fatal(err)
		chans, err := w.chans.GetText()
		fatal(err)
		w.ircChannels = make([]string, 0)
		for _, ch := range strings.Split(chans, ", ") {
			w.ircChannels = append(w.ircChannels, ch)
		}
		w.ircServer, err = w.server.GetText()
		fatal(err)
		w.ircUseTLS = w.tls.GetActive()
		fatal(err)
		w.ircTrigger, err = w.trigger.GetText()
		fatal(err)
		w.onHide()
	}

}

func (w *Window) validateInputs() error {
	nickreg := regexp.MustCompile("^\\w*$")
	nick, err := w.nick.GetText()
	fatal(err)
	if nick == "" {
		return fmt.Errorf("Empty nick")
	}
	if !nickreg.MatchString(nick) {
		return fmt.Errorf("Nick contains non alpha numeric characters")
	}
	chanreg := regexp.MustCompile("^#*\\w*$")
	chans, err := w.chans.GetText()
	fatal(err)
	for _, ch := range strings.Split(chans, ", ") {
		if !chanreg.MatchString(ch) || len(ch) <= 1 {
			return fmt.Errorf("Invalid channel name: %s", ch)
		}
	}
	serverreg := regexp.MustCompile("^\\S*:\\d*$")
	server, err := w.server.GetText()
	fatal(err)
	if !serverreg.MatchString(server) {
		return fmt.Errorf("Invalid irc server name")
	}
	triggerreg := regexp.MustCompile("^\\w*$")
	trigger, err := w.trigger.GetText()
	if len(trigger) <= 0 {
		return fmt.Errorf("Empty trigger")
	}
	fatal(err)
	if !triggerreg.MatchString(trigger) {
		return (fmt.Errorf("Trigger contains nonalphanumeric characters"))
	}
	return nil
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
