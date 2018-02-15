//This utility prints all zones in irc channels. Useful if you want to see how it will look
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/ugjka/dumbirc"
	"github.com/ugjka/newyearsbot/nyb"
)

var usage = `Test utility for debugging that post all newyears on a specified channel

CMD Options:
-chans			comma seperated list of irc channels to join
-ircserver		irc server to use irc.example.com:7000 (must be TLS enabled)
-botnick		nick for the bot 
`

//Custom flag to get irc channelsn to join
var ircChansFlag nyb.IrcChans

func init() {
	flag.Var(&ircChansFlag, "chans", "comma seperated list of irc channels to join")
}

const ircName = "nyebottest"

var once sync.Once
var start = make(chan bool)

func main() {
	//flags
	ircServer := flag.String("ircserver", "irc.freenode.net:7000", "Irc server to use, must be TLS")
	ircNick := flag.String("botnick", "", "Irc Nick for the bot")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()
	if *ircNick == "" {
		fmt.Fprintln(os.Stderr, "Error: empty nick")
		flag.Usage()
		return
	}
	if len(ircChansFlag) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no channel defined")
		flag.Usage()
		return
	}
	var zones nyb.TZS
	json.Unmarshal([]byte(nyb.Zones), &zones)
	sort.Sort(sort.Reverse(zones))

	ircobj := dumbirc.New(*ircNick, ircName, *ircServer, true)
	ircobj.AddCallback(dumbirc.WELCOME, func(msg dumbirc.Message) {
		ircobj.Join(ircChansFlag)
		//Prevent early start
		once.Do(func() {
			start <- true
		})
	})

	ircobj.AddCallback(dumbirc.PING, func(msg dumbirc.Message) {
		ircobj.Pong()
	})

	ircobj.AddCallback(dumbirc.NICKTAKEN, func(msg dumbirc.Message) {
		ircobj.Nick += "_"
		ircobj.NewNick(ircobj.Nick)
	})
	ircobj.Start()
	go func() {
		for {
			time.Sleep(time.Minute)
			ircobj.Ping()
		}
	}()
	go func() {
		log.Println(<-ircobj.Errchan)
		os.Exit(1)
	}()
	<-start

	for _, k := range zones {
		time.Sleep(time.Second * 2)
		ircobj.PrivMsg(ircChansFlag[0], "Next New Year in 29 minutes 57 seconds in "+k.String())
		time.Sleep(time.Second * 1)
		ircobj.PrivMsg(ircChansFlag[0], "Happy New Year in "+k.String())
	}

}
