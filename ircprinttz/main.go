package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	irc "github.com/ugjka/dumbirc"
	c "github.com/ugjka/newyearsbot/common"
)

const ircNick = "HNYbotTest"
const ircName = "newyears2"
const ircServer = "irc.freenode.net:7000"

var ircChannel = []string{"#ugjkatest2"}

func main() {
	var zones c.TZS
	file, err := os.Open("../tz.json")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(content, &zones)
	sort.Sort(sort.Reverse(zones))

	ircobj := irc.New(ircNick, ircName, ircServer, true)
	ircobj.AddCallback(irc.WELCOME, func(msg irc.Message) {
		ircobj.Join(ircChannel)
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
	time.Sleep(time.Second * 30)

	for _, k := range zones {
		time.Sleep(time.Second * 2)
		ircobj.PrivMsg(ircChannel[0], "Next New Year in 29 minutes 57 seconds in "+k.String())
		time.Sleep(time.Second * 1)
		ircobj.PrivMsg(ircChannel[0], "Happy New Year in "+k.String())
	}

}
