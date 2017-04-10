//Irc Bot for New Years Eve Celebration. Posts to irc when new year happens in each timezone
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	c "github.com/ugjka/newyearsbot/common"
	nyb "github.com/ugjka/newyearsbot/nyb"
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
-tzpath			path to tz database (./tz.json)
-ircserver		irc server to use irc.example.com:7000 (must be TLS enabled)
-botnick		nick for the bot
-usetls			use tls encryption for irc
`

//Default channel list
var ircChannel = []string{}

func main() {
	//flags
	ircServer := flag.String("ircserver", "irc.freenode.net:7000", "Irc server to use, must be TLS")
	ircNick := flag.String("botnick", "", "Irc Nick for the bot")
	ircTLS := flag.Bool("usetls", true, "Use tls for irc")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()
	if len(ircChansFlag) > 0 {
		ircChannel = ircChansFlag
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", "Error: No channels defined")
		flag.Usage()
		return
	}
	if *ircNick == "" {
		fmt.Fprintf(os.Stderr, "%s\n", "Error: No nick defined")
		flag.Usage()
		return
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	bot := nyb.New(*ircNick, ircChannel, *ircServer, *ircTLS)
	go func() {
		for {
			select {
			case msg := <-bot.LogCh:
				fmt.Fprintf(os.Stdout, "%s", msg)
			}
		}
	}()
	go func() {
		<-stop
		bot.Stop()
	}()
	bot.Start()
}
