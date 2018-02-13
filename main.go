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
	c "github.com/ugjka/newyearsbot/common"
	nyb "github.com/ugjka/newyearsbot/nyb"
	"mvdan.cc/xurls"
)

//Custom flag to get irc channelsn to join
var ircChansFlag c.IrcChans

func init() {
	flag.Var(&ircChansFlag, "chans", "comma seperated list of irc channels to join")
}

var usage = `New Year Eve Party Irc Bot
This bot announces new years as they happen in each timezone
You can query location using "hny" trigger for example "hny New York"

CMD Options:
-chans			comma seperated list of irc channels to join eg. "#test, #test2"
-ircserver		irc server to use irc.example.com:7000 (must be TLS enabled)
-botnick		nick for the bot
-trigger		trigger used for queries
-usetls			use tls encryption for irc
-email			Refferer Email for Nominatim
-nominatim		Nominatim server to use (Default: http://nominatim.openstreetmap.org)
`

func main() {
	//Syncing for graceful exit
	var wait sync.WaitGroup
	//Flags
	ircServer := flag.String("ircserver", "irc.freenode.net:7000", "Irc server to use, must be TLS")
	ircNick := flag.String("botnick", "", "Irc Nick for the bot")
	ircTrigger := flag.String("trigger", "hny", "trigger for queries")
	ircTLS := flag.Bool("usetls", true, "Use tls for irc")
	ircEmail := flag.String("email", "", "Refferer email for Nominatim")
	ircNominatim := flag.String("nominatim", "http://nominatim.openstreetmap.org", "Nominatim server to use")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	//Check inputs
	if *ircNominatim == "" {
		fmt.Fprint(os.Stderr, "Error: Need to provide Nominatim Server url\n")
		flag.Usage()
		return
	}
	if !xurls.Strict().MatchString(*ircNominatim) {
		fmt.Fprint(os.Stderr, "Error: Invalid Nominatim Server url\n")
		flag.Usage()
		return
	}
	if *ircEmail == "" {
		fmt.Fprint(os.Stderr, "Error: Need to provide refferer Email for Nominatim\n")
		flag.Usage()
		return
	}
	if err := checkmail.ValidateFormat(*ircEmail); err != nil {
		fmt.Fprint(os.Stderr, "Error: Invalid email address\n")
		flag.Usage()
		return
	}
	if *ircNick == "" {
		fmt.Fprintf(os.Stderr, "Error: No nick defined\n")
		flag.Usage()
		return
	}
	nickreg := regexp.MustCompile("^\\A[a-z_\\-\\[\\]\\^{}|`][a-z0-9_\\-\\[\\]\\^{}|`]{2,15}\\z$")
	if !nickreg.MatchString(*ircNick) {
		fmt.Fprintf(os.Stderr, "Error: Invalid nickname\n")
		flag.Usage()
		return
	}
	if len(ircChansFlag) <= 0 {
		fmt.Fprintf(os.Stderr, "Error: No channels defined\n")
		flag.Usage()
		return
	}
	chanreg := regexp.MustCompile("^([#&][^\\x07\\x2C\\s]{0,200})$")
	for _, ch := range ircChansFlag {
		if !chanreg.MatchString(ch) || len(ch) <= 1 {
			fmt.Fprintf(os.Stderr, "Error: Invalid channel name: %s\n", ch)
			flag.Usage()
			return
		}
	}
	if *ircServer == "" {
		fmt.Fprintf(os.Stderr, "Error: No irc server defined\n")
		flag.Usage()
		return
	}
	serverreg := regexp.MustCompile("^\\S+:\\d+$")
	if !serverreg.MatchString(*ircServer) {
		fmt.Fprintf(os.Stderr, "Error: Invalid irc server url\n")
		flag.Usage()
		return
	}
	if len(*ircTrigger) <= 0 {
		fmt.Fprintf(os.Stderr, "Error: No trigger defined\n")
		flag.Usage()
		return
	}
	triggerreg := regexp.MustCompile("^\\S+$")
	if !triggerreg.MatchString(*ircTrigger) {
		fmt.Fprintf(os.Stderr, "Error: Trigger contains white space\n")
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
			case msg, ok := <-bot.LogCh:
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
	close(bot.LogCh)
	wait.Wait()
}
