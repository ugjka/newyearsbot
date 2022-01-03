//New Year's Eve IRC party bot
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/badoux/checkmail"
	"github.com/fatih/color"
	"github.com/rhinosf1/newyearsbot/nyb"
	log "gopkg.in/inconshreveable/log15.v2"
	"mvdan.cc/xurls/v2"
)

//Custom flag for IRC channels
var channels nyb.Channels

func init() {
	flag.Var(&channels, "channels", "comma separated list of channels")
}

const usage = `
New Year's Eve IRC party bot
Announces new years as they happen in each timezone

CMD Options:
[mandatory]
-channels	comma separated list of channels eg. "#test, #test2"
-nick		irc nick
-email		nominatim email

[optional]
-password	irc password
-server		irc server (default: chat.freenode.net:6697)
-prefix		command prefix (default: !)
-ssl		use ssl for irc (default: true)
-nominatim	nominatim server (default: http://nominatim.openstreetmap.org)
-debug		debug irc traffic

`

func main() {

	//Flags
	nick := flag.String("nick", "NewYearBot", "irc nick")
	email := flag.String("email", "", "nominatim email")
	server := flag.String("server", "irc.snoonet.org:6697", "irc server")
	password := flag.String("password", "", "irc password")
	prefix := flag.String("prefix", "!", "command prefix")
	ssl := flag.Bool("ssl", true, "use ssl for irc")
	nominatim := flag.String("nominatim", "http://nominatim.openstreetmap.org", "nominatim server")
	debug := flag.Bool("debug", false, "debug irc traffic")

	green := color.New(color.FgGreen)
	flag.Usage = func() {
		green.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	//Colorize errors
	red := color.New(color.FgHiRed)

	//Check mandatory inputs
	if len(channels) == 0 {
		red.Fprintln(os.Stderr, "error: no channels defined")
		flag.Usage()
		return
	}
	channelReg := regexp.MustCompile("^([#&][^\\x07\\x2C\\s]{0,200})$")
	for _, ch := range channels {
		if !channelReg.MatchString(ch) {
			red.Fprintf(os.Stderr, "error: invalid channel name: %s\n", ch)
			flag.Usage()
			return
		}
	}
	if *nick == "" {
		red.Fprintln(os.Stderr, "error: no nick defined")
		flag.Usage()
		return
	}
	if len(*nick) > 16 {
		red.Fprintln(os.Stderr, "error: nick too long")
		flag.Usage()
		return
	}
	if *email == "" {
		red.Fprintln(os.Stderr, "error: no email provided")
		flag.Usage()
		return
	}
	if err := checkmail.ValidateFormat(*email); err != nil {
		red.Fprintln(os.Stderr, "error: invalid email address")
		flag.Usage()
		return
	}
	//Check optional inputs
	if *server == "" {
		red.Fprintln(os.Stderr, "error: no irc server defined")
		flag.Usage()
		return
	}
	serverReg := regexp.MustCompile("^\\S+:\\d+$")
	if !serverReg.MatchString(*server) {
		red.Fprintln(os.Stderr, "error: invalid irc server address")
		flag.Usage()
		return
	}
	if *prefix == "" {
		red.Fprintln(os.Stderr, "error: no command prefix defined")
		flag.Usage()
		return
	}
	prefixReg := regexp.MustCompile("^\\W+$")
	if !prefixReg.MatchString(*prefix) {
		red.Fprintln(os.Stderr, "error: prefix must be non-alphanumeric")
		flag.Usage()
		return
	}
	if *nominatim == "" {
		red.Fprintln(os.Stderr, "error: no nominatim server provided")
		flag.Usage()
		return
	}
	if !xurls.Strict().MatchString(*nominatim) {
		red.Fprintln(os.Stderr, "error: invalid nominatim server url")
		flag.Usage()
		return
	}

	bot := nyb.New(&nyb.Settings{
		Nick:      *nick,
		Channels:  channels,
		Server:    *server,
		SSL:       *ssl,
		Password:  *password,
		Prefix:    *prefix,
		Email:     *email,
		Nominatim: *nominatim,
	})
	if *debug {
		bot.LogLvl(log.LvlDebug)
	} else {
		bot.LogLvl(log.LvlInfo)
	}
	bot.Start()
}
