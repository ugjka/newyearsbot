//Irc Bot for New Years Eve Celebration. Posts to irc when new year happens in each timezone
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"sync"

	"github.com/badoux/checkmail"
	"github.com/fatih/color"
	"github.com/ugjka/newyearsbot/nyb"
	"mvdan.cc/xurls/v2"
)

//Custom flag to get irc channelsn to join
var chans nyb.IrcChans

func init() {
	flag.Var(&chans, "chans", "comma separated list of irc channels to join")
}

var usage = `
New Year Eve Party Irc Bot
This bot announces new years as they happen in each timezone
You can query location using "hny" trigger for example "hny New York"

CMD Options:
[mandatory]
-chans			comma separated list of irc channels to join eg. "#test, #test2"
-botnick		nick for the bot
-email			referrer email for Nominatim

[optional]
-nickpass		nick password
-ircserver		irc server to use (default: irc.freenode.net:7000)
-trigger		trigger used for queries. (default: hny)
-usetls			use tls encryption for irc. (default: true)
-nominatim		Nominatim server to use (default: http://nominatim.openstreetmap.org)
-ircdebug		log irc traffic

`

func main() {
	//Syncing for graceful exit
	var wait sync.WaitGroup

	//Flags
	botnick := flag.String("botnick", "", "irc nick for the bot")
	email := flag.String("email", "", "referrer email for Nominatim")
	ircServer := flag.String("ircserver", "chat.freenode.net:6697", "irc server to use")
	nickpass := flag.String("nickpass", "", "nick password")
	trigger := flag.String("trigger", "hny", "trigger for queries")
	useTLS := flag.Bool("usetls", true, "use tls for irc")
	nominatim := flag.String("nominatim", "http://nominatim.openstreetmap.org", "nominatim server to use")
	ircdebug := flag.Bool("ircdebug", false, "log irc traffic")

	green := color.New(color.FgGreen)
	flag.Usage = func() {
		green.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	//Colorize errors
	red := color.New(color.FgHiRed)

	//Check mandatory inputs
	if len(chans) == 0 {
		red.Fprintln(os.Stderr, "error: no channels defined")
		flag.Usage()
		return
	}
	chanReg := regexp.MustCompile("^([#&][^\\x07\\x2C\\s]{0,200})$")
	for _, ch := range chans {
		if !chanReg.MatchString(ch) {
			red.Fprintf(os.Stderr, "error: invalid channel name: %s\n", ch)
			flag.Usage()
			return
		}
	}
	if *botnick == "" {
		red.Fprintln(os.Stderr, "error: no nick defined")
		flag.Usage()
		return
	}
	if len(*botnick) > 16 {
		red.Fprintln(os.Stderr, "error: nick can't be longer than 16 characters")
		flag.Usage()
		return
	}
	botnickReg := regexp.MustCompile("^\\A[a-z_\\-\\[\\]\\^{}|`][a-z0-9_\\-\\[\\]\\^{}|`]{1,15}\\z$")
	if !botnickReg.MatchString(*botnick) {
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
	if *ircServer == "" {
		red.Fprintln(os.Stderr, "error: no irc server defined")
		flag.Usage()
		return
	}
	serverReg := regexp.MustCompile("^\\S+:\\d+$")
	if !serverReg.MatchString(*ircServer) {
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
	//Catch interrupt ^C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	bot := nyb.New(*botnick, chans, *nickpass, *trigger, *ircServer, *useTLS, *email, *nominatim)
	if *ircdebug {
		bot.IrcConn.SetDebugOutput(bot.LogChan)
	}
	//Log printer
	wait.Add(1)
	go func() {
		defer wait.Done()
		for {
			select {
			case msg, ok := <-bot.LogChan:
				if !ok {
					return
				}
				green.Fprintf(os.Stdout, "%s", msg)
			}
		}
	}()
	//Iterrupt catcher
	go func() {
		<-stop
		bot.Stop()
	}()
	bot.Start()

	close(bot.LogChan)
	wait.Wait()
}
