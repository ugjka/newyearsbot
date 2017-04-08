package nyb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/hako/durafmt"
	irc "github.com/ugjka/dumbirc"
	c "github.com/ugjka/newyearsbot/common"
)

const logChanLen = 100

//LogChan is a channel that sends log messages
type LogChan chan string

func (l LogChan) Write(p []byte) (n int, err error) {
	if len(l) < logChanLen {
		l <- string(p)
	}
	return len(p), nil
}

//NewLogChan make new log channel
func NewLogChan() LogChan {
	return make(chan string, logChanLen)
}

//Settings for bot
type Settings struct {
	IrcNick   string
	IrcChans  []string
	IrcServer string
	UseTLS    bool
	LogCh     LogChan
	Stopper   chan bool
	IrcObj    *irc.Connection
	Exited    bool
}

//Stop stops the bot
func (s *Settings) Stop() {
	s.Stopper <- true
}

//New creates new bot
func New(nick string, chans []string, server string, tls bool) *Settings {
	return &Settings{
		nick,
		chans,
		server,
		tls,
		make(chan string, 100),
		make(chan bool),
		&irc.Connection{},
		false,
	}
}

//Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	//return time.Date(tmp.Year(), time.April, 4, 0, 0, 0, 0, time.UTC)
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

//Start starts the bot
func (s *Settings) Start() {
	//log.SetOutput(os.Stderr)
	log.SetOutput(s.LogCh)
	log.Println("Starting the bot...")
	var start = make(chan bool)
	var once sync.Once
	var next c.TZ

	//
	//Set up irc and its callbacks
	//
	s.IrcObj = irc.New(s.IrcNick, "nyebot", s.IrcServer, s.UseTLS)
	s.IrcObj.AddCallback(irc.WELCOME, func(msg irc.Message) {
		s.IrcObj.Join(s.IrcChans)
		//Prevent early start
		once.Do(func() {
			start <- true
		})
	})
	//Reply ping messages with pong
	s.IrcObj.AddCallback(irc.PING, func(msg irc.Message) {
		log.Println("PING recieved, sending PONG")
		s.IrcObj.Pong()
	})
	//Change nick if taken
	s.IrcObj.AddCallback(irc.NICKTAKEN, func(msg irc.Message) {
		if strings.HasSuffix(s.IrcObj.Nick, "_") {
			s.IrcObj.Nick = s.IrcObj.Nick[:len(s.IrcObj.Nick)-1]
		} else {
			s.IrcObj.Nick += "_"
		}
		s.IrcObj.NewNick(s.IrcObj.Nick)
	})
	//Handler for Location queries
	s.IrcObj.AddCallback(irc.PRIVMSG, func(msg irc.Message) {
		if strings.HasPrefix(msg.Trailing, "hny !next") {
			dur, err := time.ParseDuration(next.Offset + "h")
			if err != nil {
				return
			}
			humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(dur)).String())
			if err != nil {
				return
			}
			s.IrcObj.Reply(msg, fmt.Sprintf("Next new year in %s in %s", humandur, next.String()))
			return
		}
		if strings.HasPrefix(msg.Trailing, "hny ") {
			tz, err := getNewYear(msg.Trailing[4:])
			if err != nil {
				s.IrcObj.Reply(msg, "Some error occurred!")
				return
			}
			s.IrcObj.Reply(msg, tz)
			return
		}

	})
	s.IrcObj.Start()
	//Reconnect logic and Irc Pinger
	stopper := make(chan bool)
	go func() {
		var err error
		for {
			timer := time.NewTimer(time.Minute)
			select {
			case err = <-s.IrcObj.Errchan:
				log.Println(err)
				time.Sleep(time.Second * 30)
				log.Println("Restarting the bot...")
				s.IrcObj.Start()
			case <-stopper:
				s.IrcObj.Disconnect()
				return
			case <-timer.C:
				log.Println("Sending PING...")
				timer.Stop()
				s.IrcObj.Ping()
			}
		}
	}()
	//Starts when joined, see once.Do
	<-start
	var zones c.TZS
	if err := json.Unmarshal([]byte(TZ), &zones); err != nil {
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
			log.Println("Zone pending:", zones[i].Offset)
			msg := fmt.Sprintf("Next New Year in %s in %s", humandur, zones[i])
			s.IrcObj.PrivMsgBulk(s.IrcChans, msg)
			//Wait till Target in Timezone
			timer := time.NewTimer(target.Sub(time.Now().UTC().Add(dur)))

			select {
			case <-timer.C:
				log.Println("Announcing zone:", zones[i].Offset)
				msg = fmt.Sprintf("Happy New Year in %s", zones[i])
				s.IrcObj.PrivMsgBulk(s.IrcChans, msg)
			case <-s.Stopper:
				log.Println("Stopping the bot...")
				timer.Stop()
				stopper <- true
				log.Println("Disconnecting...")
				s.IrcObj.Disconnect()
				return
			}
		}
	}
	s.IrcObj.PrivMsgBulk(s.IrcChans, fmt.Sprintf("That's it, year %d is here across the globe", target.Year()))
	log.Println("All zones finished...")
	stopper <- true
	s.Exited = true
}

//Func for querying newyears in specified location
func getNewYear(loc string) (string, error) {
	log.Println("Querying location:", loc)
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
