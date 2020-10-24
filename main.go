//Irc Bot for New Years Eve Celebration. Posts to irc when new year happens in each timezone
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/badoux/checkmail"
	"github.com/fatih/color"
	"github.com/ugjka/newyearsbot/nyb"
	log "gopkg.in/inconshreveable/log15.v2"
	"mvdan.cc/xurls/v2"
)

//Custom flag to get irc channelsn to join
var channels nyb.IrcChans

func init() {
	flag.Var(&channels, "channels", "comma separated list of irc channels to join")
}

const usage = `
New Year Eve Party Irc Bot
This bot announces new years as they happen in each timezone
You can query location using "hny" trigger for example "hny New York"

CMD Options:
[mandatory]
-channels	comma separated list of irc channels to join eg. "#test, #test2"
-nick		nick for the bot
-email		referrer email for Nominatim

[optional]
-password	nick password
-server		irc server to use (default: chat.freenode.net:6697)
-trigger	trigger used for queries. (default: hny)
-ssl		use ssl encryption for irc. (default: true)
-nominatim	Nominatim server to use (default: http://nominatim.openstreetmap.org)
-debug		debug irc traffic

`

func main() {

	//Flags
	nick := flag.String("nick", "", "irc nick for the bot")
	email := flag.String("email", "", "referrer email for Nominatim")
	server := flag.String("server", "chat.freenode.net:6697", "irc server to use")
	password := flag.String("password", "", "nick password")
	trigger := flag.String("trigger", "hny", "trigger for queries")
	ssl := flag.Bool("ssl", true, "use ssl for irc")
	nominatim := flag.String("nominatim", "http://nominatim.openstreetmap.org", "nominatim server to use")
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
	chanReg := regexp.MustCompile("^([#&][^\\x07\\x2C\\s]{0,200})$")
	for _, ch := range channels {
		if !chanReg.MatchString(ch) {
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
		red.Fprintln(os.Stderr, "error: nick can't be longer than 16 characters")
		flag.Usage()
		return
	}
	botnickReg := regexp.MustCompile("^\\A[a-z_\\-\\[\\]\\^{}|`][a-z0-9_\\-\\[\\]\\^{}|`]{1,15}\\z$")
	if !botnickReg.MatchString(*nick) {
		red.Fprintln(os.Stderr, "error: invalid nickname")
		flag.Usage()
		return
	}
	if *email == "" {
		red.Fprintln(os.Stderr, "error: need to provide referrer email for Nominatim")
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
	if *trigger == "" {
		red.Fprintln(os.Stderr, "error: no trigger defined")
		flag.Usage()
		return
	}
	triggerReg := regexp.MustCompile("^\\S+$")
	if !triggerReg.MatchString(*trigger) {
		red.Fprintln(os.Stderr, "error: trigger contains white space")
		flag.Usage()
		return
	}
	if *nominatim == "" {
		red.Fprintln(os.Stderr, "error: need to provide a Nominatim Server url")
		flag.Usage()
		return
	}
	if !xurls.Strict().MatchString(*nominatim) {
		red.Fprintln(os.Stderr, "error: invalid Nominatim server url")
		flag.Usage()
		return
	}

	bot := nyb.New(*nick, channels, *password, *trigger, *server, *ssl, *email, *nominatim)
	if *debug {
		bot.LogLvl(log.LvlDebug)
	} else {
		bot.LogLvl(log.LvlInfo)
	}
	bot.Start()
}
