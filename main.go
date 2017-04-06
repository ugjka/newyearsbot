//Irc Bot for New Years Eve Celebration. Posts to irc when new year happens in each timezone
package main

import (
	"flag"
	"fmt"
	"os"

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
-botnick		nick for the bot `

//Default channel list
var ircChannel = []string{"#ugjka", "#ugjkatest", "#ugjkatest2"}

func main() {
	//flags
	ircServer := flag.String("ircserver", "irc.freenode.net:7000", "Irc server to use, must be TLS")
	ircNick := flag.String("botnick", "HNYbot18", "Irc Nick for the bot")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()
	if len(ircChansFlag) > 0 {
		ircChannel = ircChansFlag
	}
	bot := nyb.New(*ircNick, ircChannel, *ircServer, true)
	bot.Start()
}
