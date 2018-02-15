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
	nyb "github.com/ugjka/newyearsbot/nyb"
	"mvdan.cc/xurls"
)

//Custom flag to get irc channelsn to join
var ircChansFlag nyb.IrcChans

func init() {
	flag.Var(&ircChansFlag, "chans", "comma seperated list of irc channels to join")
}

var usage = `
New Year Eve Party Irc Bot
This bot announces new years as they happen in each timezone
You can query location using "hny" trigger for example "hny New York"

CMD Options:
[mandatory]
-chans			comma seperated list of irc channels to join eg. "#test, #test2"
-botnick		nick for the bot
-email			referrer email for Nominatim

[optional]
-ircserver		irc server to use (default: irc.freenode.net:7000)
-trigger		trigger used for queries. (default: hny)
-usetls			use tls encryption for irc. (default: true)
-nominatim		Nominatim server to use (default: http://nominatim.openstreetmap.org)

`

func main() {
	//Syncing for graceful exit
	var wait sync.WaitGroup
	//Flags
	ircServer := flag.String("ircserver", "irc.freenode.net:7000", "irc server to use")
	ircNick := flag.String("botnick", "", "irc nick for the bot")
	ircTrigger := flag.String("trigger", "hny", "trigger for queries")
	ircTLS := flag.Bool("usetls", true, "use tls for irc")
	ircEmail := flag.String("email", "", "referrer email for Nominatim")
	ircNominatim := flag.String("nominatim", "http://nominatim.openstreetmap.org", "nominatim server to use")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	//Check inputs
	if *ircNominatim == "" {
		fmt.Fprint(os.Stderr, "error: need to provide a Nominatim Server url\n")
		flag.Usage()
		return
	}
	if !xurls.Strict().MatchString(*ircNominatim) {
		fmt.Fprint(os.Stderr, "error: invalid Nominatim server url\n")
		flag.Usage()
		return
	}
	if *ircEmail == "" {
		fmt.Fprint(os.Stderr, "error: need to provide referrer email for Nominatim\n")
		flag.Usage()
		return
	}
	if err := checkmail.ValidateFormat(*ircEmail); err != nil {
		fmt.Fprint(os.Stderr, "error: invalid email address\n")
		flag.Usage()
		return
	}
	if *ircNick == "" {
		fmt.Fprintf(os.Stderr, "error: no bot nick defined\n")
		flag.Usage()
		return
	}
	nickreg := regexp.MustCompile("^\\A[a-z_\\-\\[\\]\\^{}|`][a-z0-9_\\-\\[\\]\\^{}|`]{2,15}\\z$")
	if !nickreg.MatchString(*ircNick) {
		fmt.Fprintf(os.Stderr, "error: invalid nickname\n")
		flag.Usage()
		return
	}
	if len(ircChansFlag) == 0 {
		fmt.Fprintf(os.Stderr, "error: no channels defined\n")
		flag.Usage()
		return
	}
	chanreg := regexp.MustCompile("^([#&][^\\x07\\x2C\\s]{0,200})$")
	for _, ch := range ircChansFlag {
		if !chanreg.MatchString(ch) {
			fmt.Fprintf(os.Stderr, "error: invalid channel name: %s\n", ch)
			flag.Usage()
			return
		}
	}
	if *ircServer == "" {
		fmt.Fprintf(os.Stderr, "error: no irc server defined\n")
		flag.Usage()
		return
	}
	serverreg := regexp.MustCompile("^\\S+:\\d+$")
	if !serverreg.MatchString(*ircServer) {
		fmt.Fprintf(os.Stderr, "error: invalid irc server address\n")
		flag.Usage()
		return
	}
	if *ircTrigger == "" {
		fmt.Fprintf(os.Stderr, "error: no trigger defined\n")
		flag.Usage()
		return
	}
	triggerreg := regexp.MustCompile("^\\S+$")
	if !triggerreg.MatchString(*ircTrigger) {
		fmt.Fprintf(os.Stderr, "error: trigger contains white space\n")
		return
	}
	//Catch interrupt ^C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	//New bot instance
	bot := nyb.New(*ircNick, ircChansFlag, *ircTrigger, *ircServer, *ircTLS, *ircEmail, *ircNominatim)
	//Log printer
	go func() {
		wait.Add(1)
		defer wait.Done()
		for {
			select {
			case msg, ok := <-bot.LogChan:
				if !ok {
					return
				}
				fmt.Fprintf(os.Stdout, "%s", msg)
			}
		}
	}()
	//Iterrupt catcher
	go func() {
		<-stop
		bot.Stop()
	}()
	bot.Start()
	//close and wait
	close(bot.LogChan)
	wait.Wait()
}
