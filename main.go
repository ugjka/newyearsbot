//Irc Bot for New Years Eve Celebration. Posts to irc when new year happens in each timezone
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/hako/durafmt"
	irc "github.com/ugjka/dumbirc"
	c "github.com/ugjka/newyearsbot/common"
)

//Custom flag to get irc channelsn to join
var ircChansFlag c.IrcChans

func init() {
	flag.Var(&ircChansFlag, "chans", "comma seperated list of irc channels to join")
}

//Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year()-1, time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

var usage = `New Year Eve Party Irc Bot
This bot announces new years as they happen in each timezone
You can query location using "hny" trigger for example "hny New York"

CMD Options:
-chans			comma seperated list of irc channels to join
-tzpath			path to tz database (./tz.json)
-ircserver		irc server to use irc.example.com:7000 (must be TLS enabled)
-botnick		nick for the bot `

const ircName = "nyebot"

//Default channel list
var ircChannel = []string{"#ugjka", "#ugjkatest", "#ugjkatest2"}

var start = make(chan bool)
var once sync.Once
var next c.TZ

func main() {
	//flags
	tzdatapath := flag.String("tzpath", "./tz.json", "path to tz.json")
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

	//
	//Set up irc and its callbacks
	//
	ircobj := irc.New(*ircNick, ircName, *ircServer, true)
	ircobj.AddCallback(irc.WELCOME, func(msg irc.Message) {
		ircobj.Join(ircChannel)
		//Prevent early start
		once.Do(func() {
			start <- true
		})
	})
	//Reply ping messages with pong
	ircobj.AddCallback(irc.PING, func(msg irc.Message) {
		ircobj.Pong()
	})
	//Change nick if taken
	ircobj.AddCallback(irc.NICKTAKEN, func(msg irc.Message) {
		if strings.HasSuffix(ircobj.Nick, "_") {
			ircobj.Nick = ircobj.Nick[:len(ircobj.Nick)-1]
		} else {
			ircobj.Nick += "_"
		}
		ircobj.NewNick(ircobj.Nick)
	})
	//Handler for Location queries
	ircobj.AddCallback(irc.PRIVMSG, func(msg irc.Message) {
		if strings.HasPrefix(msg.Trailing, "hny !next") {
			dur, err := time.ParseDuration(next.Offset + "h")
			if err != nil {
				return
			}
			humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(dur)).String())
			if err != nil {
				return
			}
			ircobj.Reply(msg, fmt.Sprintf("Next new year in %s in %s", humandur, next.String()))
			return
		}
		if strings.HasPrefix(msg.Trailing, "hny ") {
			tz, err := getNewYear(msg.Trailing[4:])
			if err != nil {
				ircobj.Reply(msg, "Some error occurred!")
				return
			}
			ircobj.Reply(msg, tz)
			return
		}

	})
	ircobj.Start()
	//IRC pinger
	go func() {
		for {
			time.Sleep(time.Minute)
			ircobj.Ping()
		}
	}()
	//Reconnect logic
	go func() {
		for {
			log.Println(<-ircobj.Errchan)
			time.Sleep(time.Second * 30)
			ircobj.Start()
		}
	}()
	//Starts when joined, see once.Do
	<-start
	var zones c.TZS
	file, err := os.Open(*tzdatapath)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(content, &zones); err != nil {
		log.Fatal(err)
	}
	sort.Sort(sort.Reverse(zones))
	for i := 0; i < len(zones); i++ {
		dur, err := time.ParseDuration(zones[i].Offset + "h")
		if err != nil {
			log.Fatal(err)
		}
		//Check if zone is past target
		if time.Now().UTC().Add(dur).Before(target) {
			next = zones[i]
			time.Sleep(time.Second * 2)
			humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(dur)).String())
			if err != nil {
				log.Fatal(err)
			}
			msg := fmt.Sprintf("Next New Year in %s in %s", humandur, zones[i])
			ircobj.PrivMsgBulk(ircChannel, msg)
			//Wait till Target in Timezone
			time.Sleep(target.Sub(time.Now().UTC().Add(dur)))
			msg = fmt.Sprintf("Happy New Year in %s", zones[i])
			ircobj.PrivMsgBulk(ircChannel, msg)
		}
	}
	ircobj.PrivMsgBulk(ircChannel, fmt.Sprintf("That's it, year %d is here across the globe", target.Year()))

}

//Func for querying newyears in specified location
func getNewYear(loc string) (string, error) {
	maps := url.Values{}
	maps.Add("address", loc)
	maps.Add("sensor", "false")
	maps.Add("language", "en")
	data, err := c.Getter(c.Geocode + maps.Encode())
	if err != nil {
		log.Println(err)
		return "", err
	}
	var mapj c.Gmap
	if err = json.Unmarshal(data, &mapj); err != nil {
		log.Println(err)
		return "", err
	}
	if mapj.Status != "OK" {
		return "I don't know that place.", nil
	}
	adress := mapj.Results[0].FormattedAddress
	location := fmt.Sprintf("%.7f,%.7f", mapj.Results[0].Geometry.Location.Lat, mapj.Results[0].Geometry.Location.Lng)
	tmzone := url.Values{}
	tmzone.Add("location", location)
	tmzone.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	tmzone.Add("sensor", "false")
	data, err = c.Getter(c.Timezone + tmzone.Encode())
	if err != nil {
		log.Println(err)
		return "", err
	}
	var timej c.Gtime
	if err = json.Unmarshal(data, &timej); err != nil {
		log.Println(err)
		return "", err
	}
	if timej.Status != "OK" {
		return "Couldn't get timezone info.", nil
	}
	//RawOffset
	raw, err := time.ParseDuration(fmt.Sprintf("%ds", timej.RawOffset))
	if err != nil {
		log.Println(err)
		return "", err
	}
	//DstOffset
	dst, err := time.ParseDuration(fmt.Sprintf("%ds", timej.DstOffset))
	if err != nil {
		log.Println(err)
		return "", err
	}
	//Check if past target
	if time.Now().UTC().Add(raw + dst).Before(target) {
		humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(raw + dst)).String())
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf("New Year in %s will happen in %s", adress, humandur), nil
	}
	return fmt.Sprintf("New year in %s already happened.", adress), nil
}
