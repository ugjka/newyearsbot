//This utility prints all zones in irc channels. Useful if you want to see how it will look
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	irc "github.com/ugjka/dumbirc"
	c "github.com/ugjka/newyearsbot/common"
)

var usage = `Test utility for debugging that post all newyears on a specified channel

CMD Options:
-chans			comma seperated list of irc channels to join
-tzpath			path to tz database (../tz.json)
-ircserver		irc server to use irc.example.com:7000 (must be TLS enabled)
-botnick		nick for the bot `

//Custom flag to get irc channelsn to join
var ircChansFlag c.IrcChans

func init() {
	flag.Var(&ircChansFlag, "chans", "comma seperated list of irc channels to join")
}

const ircName = "nyebottest"

var ircChannel = []string{"#ugjkatest"}
var once sync.Once
var start = make(chan bool)

func main() {
	//flags
	tzdatapath := flag.String("tzpath", "../tz.json", "path to tz.json")
	ircServer := flag.String("ircserver", "irc.freenode.net:7000", "Irc server to use, must be TLS")
	ircNick := flag.String("botnick", "HNYbot18", "Irc Nick for the bot")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()
	if len(ircChansFlag) > 0 {
		ircChannel = ircChansFlag
	}
	//Check if tz.json exists
	if _, err := os.Stat(*tzdatapath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: file %s does not exist\n", *tzdatapath)
		os.Exit(1)
	}
	var zones c.TZS
	file, err := os.Open(*tzdatapath)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(content, &zones)
	sort.Sort(sort.Reverse(zones))

	ircobj := irc.New(*ircNick, ircName, *ircServer, true)
	ircobj.AddCallback(irc.WELCOME, func(msg irc.Message) {
		ircobj.Join(ircChannel)
		//Prevent early start
		once.Do(func() {
			start <- true
		})
	})

	ircobj.AddCallback(irc.PING, func(msg irc.Message) {
		ircobj.Pong()
	})

	ircobj.AddCallback(irc.NICKTAKEN, func(msg irc.Message) {
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
		ircobj.PrivMsg(ircChannel[0], "Next New Year in 29 minutes 57 seconds in "+k.String())
		time.Sleep(time.Second * 1)
		ircobj.PrivMsg(ircChannel[0], "Happy New Year in "+k.String())
	}

}
