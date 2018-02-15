package nyb

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hako/durafmt"
	"github.com/ugjka/dumbirc"
	"github.com/ugjka/go-tz"
)

func (s *Settings) addCallbacks() {
	bot := s.Bot
	//On any message send a signal to ping timer to be ready
	bot.AddCallback(dumbirc.ANYMESSAGE, func(msg dumbirc.Message) {
		pingpong(s.pp)
	})

	bot.AddCallback(dumbirc.WELCOME, func(msg dumbirc.Message) {
		bot.Join(s.IrcChans)
		//Prevent early start
		s.Do(func() {
			close(s.start)
		})
	})

	bot.AddCallback(dumbirc.PING, func(msg dumbirc.Message) {
		log.Println("PING recieved, sending PONG")
		bot.Pong()
	})

	bot.AddCallback(dumbirc.PONG, func(msg dumbirc.Message) {
		log.Println("Got PONG...")
	})
	//Change nick if taken
	//This could loop forever in some cases
	//Could be improved
	bot.AddCallback(dumbirc.NICKTAKEN, func(msg dumbirc.Message) {
		log.Println("Nick taken, changing...")
		time.Sleep(time.Second * 5)
		bot.Nick = changeNick(bot.Nick)
		log.Printf("New nick: %s", bot.Nick)
		bot.NewNick(bot.Nick)
	})
}

func (s *Settings) addTriggers() {
	bot := s.Bot
	//Trigger for !help
	stHelp := "%s: Query location: '%s <location>', Next zone: '%s !next', Last zone: '%s !last', Source code: https://github.com/ugjka/newyearsbot"
	bot.AddTrigger(dumbirc.Trigger{
		Condition: func(msg dumbirc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s !help", s.IrcTrigger))
		},
		Response: func(msg dumbirc.Message) {
			bot.Reply(msg, fmt.Sprintf(stHelp, msg.Name, s.IrcTrigger, s.IrcTrigger, s.IrcTrigger))
		},
	})
	//Trigger for !next
	bot.AddTrigger(dumbirc.Trigger{
		Condition: func(msg dumbirc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s !next", s.IrcTrigger))
		},
		Response: func(msg dumbirc.Message) {
			log.Println("Querying !next...")
			dur, err := time.ParseDuration(s.next.Offset + "h")
			if err != nil {
				return
			}
			if time.Now().UTC().Add(dur).After(target) {
				bot.Reply(msg, fmt.Sprintf("No more next, %d is here AoE", target.Year()))
				return
			}
			humandur := durafmt.Parse(target.Sub(time.Now().UTC().Add(dur)))
			bot.Reply(msg, fmt.Sprintf("Next New Year in %s in %s",
				removeMilliseconds(humandur), s.next.String()))
		},
	})
	//Trigger for !last
	bot.AddTrigger(dumbirc.Trigger{
		Condition: func(msg dumbirc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s !last", s.IrcTrigger))
		},
		Response: func(msg dumbirc.Message) {
			log.Println("Querying !last...")
			dur, err := time.ParseDuration(s.last.Offset + "h")
			if err != nil {
				return
			}
			humandur := durafmt.Parse(time.Now().UTC().Add(dur).Sub(target))
			if s.last.Offset == "-12" {
				humandur = durafmt.Parse(time.Now().UTC().Add(dur).Sub(target.AddDate(-1, 0, 0)))
			}
			bot.Reply(msg, fmt.Sprintf("Last NewYear %s ago in %s",
				removeMilliseconds(humandur), s.last.String()))
		},
	})
	//Trigger for location queries
	bot.AddTrigger(dumbirc.Trigger{
		Condition: func(msg dumbirc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				!strings.Contains(msg.Trailing, "!next") &&
				!strings.Contains(msg.Trailing, "!last") &&
				!strings.Contains(msg.Trailing, "!help") &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s ", s.IrcTrigger))
		},
		Response: func(msg dumbirc.Message) {
			tz, err := s.getNewYear(msg.Trailing[len(s.IrcTrigger)+1:])
			if err == errNoZone || err == errNoPlace {
				log.Println("Query error:", err)
				bot.Reply(msg, fmt.Sprintf("%s: %s", msg.Name, err))
				return
			}
			if err != nil {
				log.Println("Query error:", err)
				bot.Reply(msg, "Some error occurred!")
				return
			}
			bot.Reply(msg, fmt.Sprintf("%s: %s", msg.Name, tz))
		},
	})
}

var (
	errNoZone  = errors.New("couldn't get the timezone for that location")
	errNoPlace = errors.New("Couldn't find that place")
)

func (s *Settings) getNominatimReqURL(location *string) string {
	maps := url.Values{}
	maps.Add("q", *location)
	maps.Add("format", "json")
	maps.Add("accept-language", "en")
	maps.Add("limit", "1")
	maps.Add("email", s.Email)
	return s.Nominatim + NominatimGeoCode + maps.Encode()
}

var stNewYearWillHappen = "New Year in %s will happen in %s"
var stNewYearHappenned = "New Year in %s happened %s ago"

func (s *Settings) getNewYear(location string) (string, error) {
	log.Println("Querying location:", location)
	data, err := NominatimGetter(s.getNominatimReqURL(&location))
	if err != nil {
		log.Println(err)
		return "", err
	}
	if err = json.Unmarshal(data, &s.nominatimResult); err != nil {
		log.Println(err)
		return "", err
	}
	if len(s.nominatimResult) == 0 {
		return "", errNoPlace
	}
	lat, err := strconv.ParseFloat(s.nominatimResult[0].Lat, 64)
	if err != nil {
		return "", err
	}
	lng, err := strconv.ParseFloat(s.nominatimResult[0].Lon, 64)
	if err != nil {
		return "", err
	}
	p := gotz.Point{
		Lat: lat,
		Lng: lng,
	}
	zone, err := gotz.GetZone(p)
	if err != nil {
		return "", errNoZone
	}

	offset, err := time.ParseDuration(fmt.Sprintf("%ds", getOffset(target, zone)))
	if err != nil {
		log.Println(err)
		return "", err
	}
	adress := s.nominatimResult[0].DisplayName

	if time.Now().UTC().Add(offset).Before(target) {
		humandur := durafmt.Parse(target.Sub(time.Now().UTC().Add(offset)))
		return fmt.Sprintf(stNewYearWillHappen, adress, removeMilliseconds(humandur)), nil
	}
	humandur := durafmt.Parse(time.Now().UTC().Add(offset).Sub(target))
	return fmt.Sprintf(stNewYearHappenned, adress, removeMilliseconds(humandur)), nil
}
