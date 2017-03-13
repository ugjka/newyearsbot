package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	irc "github.com/ugjka/dumbirc"
)

type tz struct {
	Countries []struct {
		Name   string   `json:"name"`
		Cities []string `json:"cities"`
	} `json:"countries"`
	Offset string `json:"offset"`
}

func (t tz) String() (x string) {
	for i, k := range t.Countries {
		x += fmt.Sprintf("%s", k.Name)
		for i, k1 := range k.Cities {
			if k1 == "" {
				continue
			}
			if i == 0 {
				x += " ("
			}
			x += fmt.Sprintf("%s", k1)
			if i >= 0 && i < len(k.Cities)-1 {
				x += ", "
			}
			if i == len(k.Cities)-1 {
				x += ")"
			}
		}
		if i < len(t.Countries)-1 {
			x += ", "
		}
	}
	return
}

type tzs []tz

func (t tzs) Len() int {
	return len(t)
}

func (t tzs) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tzs) Less(i, j int) bool {
	x, err := strconv.ParseFloat(t[i].Offset, 64)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseFloat(t[j].Offset, 64)
	if err != nil {
		log.Fatal(err)
	}
	if x < y {
		return true
	}
	return false
}

const ircNick = "HNYbotTest"
const ircName = "newyears2"
const ircServer = "irc.freenode.net:7000"

var ircChannel = []string{"#ugjkatest2"}

func main() {
	var zones tzs
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
