package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/hako/durafmt"
	irc "github.com/ugjka/dumbirc"
	c "github.com/ugjka/newyearsbot/common"
)

var target = time.Date(2017, time.March, 14, 0, 0, 0, 0, time.UTC)

const ircNick = "HNYbot18"
const ircName = "newyears"
const ircServer = "irc.freenode.net:7000"

var ircChannel = []string{"#ugjka", "#ugjkatest", "#ugjkatest2"}

func main() {
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
	ircobj.AddCallback(irc.PRIVMSG, func(msg irc.Message) {
		post := msg.Trailing
		if strings.HasPrefix(post, "hny ") {
			tz, err := getTimeZone(post[4:])
			if err != nil {
				ircobj.PrivMsg(ircChannel[0], err.Error())
			}
			ircobj.PrivMsg(ircChannel[0], tz)
		}
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
	var zones c.TZS
	file, err := os.Open("./tz.json")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(content, &zones)
	sort.Sort(sort.Reverse(zones))
	for i := 0; i < len(zones); i++ {
		dur, err := time.ParseDuration(zones[i].Offset + "h")
		if err != nil {
			log.Fatal(err)
		}
		if time.Now().UTC().Add(dur).Before(target) {
			time.Sleep(time.Second * 2)
			humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(dur)).String())
			if err != nil {
				log.Fatal(err)
			}
			tmp := fmt.Sprint("Next New Year in ", humandur, " in ", zones[i])
			for _, k := range ircChannel {
				ircobj.PrivMsg(k, tmp)
			}
			time.Sleep(target.Sub(time.Now().UTC().Add(dur)))
			tmp = fmt.Sprint("Happy New Year in ", zones[i])
			for _, k := range ircChannel {
				ircobj.PrivMsg(k, tmp)
			}
		}
	}
	for _, k := range ircChannel {
		ircobj.PrivMsg(k, "That's it, the New Year is here across the globe!")
	}
}

func getTimeZone(loc string) (string, error) {
	client := http.Client{}
	maps := url.Values{}
	maps.Add("address", loc)
	maps.Add("sensor", "false")
	maps.Add("language", "en")
	req, err := http.NewRequest("GET", c.Geocode+maps.Encode(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla")
	get, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer get.Body.Close()
	text, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return "", err
	}
	var mapj c.Gmap
	json.Unmarshal(text, &mapj)
	log.Println(mapj)
	if mapj.Status != "OK" {
		return "", errors.New("I don't know that place.")
	}
	adress := mapj.Results[0].FormattedAddress
	tmzone := url.Values{}
	location := fmt.Sprintf("%.6f,%.6f", mapj.Results[0].Geometry.Location.Lat, mapj.Results[0].Geometry.Location.Lng)
	tmzone.Add("location", location)
	tmzone.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	tmzone.Add("sensor", "false")

	req2, err := http.NewRequest("GET", c.Timezone+tmzone.Encode(), nil)
	if err != nil {
		return "", err
	}
	req2.Header.Set("User-Agent", "Mozilla")
	get2, err := client.Do(req2)
	if err != nil {
		return "", err
	}
	defer get2.Body.Close()
	text, err = ioutil.ReadAll(get2.Body)
	if err != nil {
		return "", err
	}
	var timej c.Gtime
	json.Unmarshal(text, &timej)
	if timej.Status != "OK" {
		return "", errors.New("Couldn't get timezone info.")
	}
	log.Println(timej)
	raw, err := time.ParseDuration(fmt.Sprintf("%ds", timej.RawOffset))
	if err != nil {
		log.Fatal(err)
	}
	dst, err := time.ParseDuration(fmt.Sprintf("%ds", timej.DstOffset))
	if err != nil {
		log.Fatal(err)
	}
	if time.Now().UTC().Add(raw + dst).Before(target) {
		time.Sleep(time.Second * 2)
		humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(raw + dst)).String())
		if err != nil {
			log.Fatal(err)
		}
		return fmt.Sprint("New Year in ", adress, " will happen in ", humandur), nil
	}
	return fmt.Sprintf("New year in %s already happened.", adress), nil
}
