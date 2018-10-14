package nyb

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/hako/durafmt"
	"github.com/ugjka/dumbirc"
	"github.com/ugjka/go-tz"
)

var nickChangeInterval = time.Second * 5

func (bot *Settings) addCallbacks() {
	irc := bot.IrcConn
	//On any message send a signal to ping timer to be ready

	irc.AddCallback(dumbirc.WELCOME, func(msg *dumbirc.Message) {
		if irc.Password != "" {
			confirmErr := fmt.Errorf("did not get identification confirmation")
			err := irc.WaitFor(func(m *dumbirc.Message) bool {
				return m.Command == dumbirc.NOTICE && strings.Contains(m.Content, "You are now identified for")
			},
				func() {},
				time.Second*30,
				confirmErr,
			)
			if err == confirmErr {
				log.Println(err)
				log.Println("trying to start the bot anyway")
			} else if err != nil {
				return
			}
		}
		irc.Join(bot.IrcChans)
		//Prevent early start
		bot.Do(func() {
			close(bot.start)
		})
	})
}

func (bot *Settings) addTriggers() {
	irc := bot.IrcConn
	//Trigger for !help
	stHelp := "%s: Query location: '%s <location>', Next zone: '%s !next', Last zone: '%s !last', Remaining: '%s !remaining', Source code: https://github.com/ugjka/newyearsbot"
	irc.AddTrigger(dumbirc.Trigger{
		Condition: func(msg *dumbirc.Message) bool {
			return msg.Command == dumbirc.PRIVMSG &&
				msg.Content == fmt.Sprintf("%s !help", bot.IrcTrigger)
		},
		Response: func(msg *dumbirc.Message) {
			log.Println("Querying !help...")
			irc.Reply(msg, fmt.Sprintf(stHelp, msg.Name, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger))
		},
	})
	//Trigger for !next
	irc.AddTrigger(dumbirc.Trigger{
		Condition: func(msg *dumbirc.Message) bool {
			return msg.Command == dumbirc.PRIVMSG &&
				msg.Content == fmt.Sprintf("%s !next", bot.IrcTrigger)
		},
		Response: func(msg *dumbirc.Message) {
			log.Println("Querying !next...")
			dur := time.Minute * time.Duration(bot.next.Offset*60)
			if timeNow().UTC().Add(dur).After(target) {
				irc.Reply(msg, fmt.Sprintf("No more next, %d is here AoE", target.Year()))
				return
			}
			humandur := durafmt.Parse(target.Sub(timeNow().UTC().Add(dur)))
			irc.Reply(msg, fmt.Sprintf("Next New Year in %s in %s",
				removeMilliseconds(humandur), bot.next))
		},
	})
	//Trigger for !last
	irc.AddTrigger(dumbirc.Trigger{
		Condition: func(msg *dumbirc.Message) bool {
			return msg.Command == dumbirc.PRIVMSG &&
				msg.Content == fmt.Sprintf("%s !last", bot.IrcTrigger)
		},
		Response: func(msg *dumbirc.Message) {
			log.Println("Querying !last...")
			dur := time.Minute * time.Duration(bot.last.Offset*60)
			humandur := durafmt.Parse(timeNow().UTC().Add(dur).Sub(target))
			if bot.last.Offset == -12 {
				humandur = durafmt.Parse(timeNow().UTC().Add(dur).Sub(target.AddDate(-1, 0, 0)))
			}
			irc.Reply(msg, fmt.Sprintf("Last New Year %s ago in %s",
				removeMilliseconds(humandur), bot.last))
		},
	})
	//Trigger for !remaining
	irc.AddTrigger(dumbirc.Trigger{
		Condition: func(msg *dumbirc.Message) bool {
			return msg.Command == dumbirc.PRIVMSG &&
				msg.Content == fmt.Sprintf("%s !remaining", bot.IrcTrigger)
		},
		Response: func(msg *dumbirc.Message) {
			log.Println("Querying !remaining...")
			ss := "s"
			if bot.remaining == 1 {
				ss = ""
			}
			irc.Reply(msg, fmt.Sprintf("%s: %d timezone%s remaining", msg.Name, bot.remaining, ss))
		},
	})
	//Trigger for location queries
	irc.AddTrigger(dumbirc.Trigger{
		Condition: func(msg *dumbirc.Message) bool {
			return msg.Command == dumbirc.PRIVMSG &&
				!strings.Contains(msg.Content, "!next") &&
				!strings.Contains(msg.Content, "!last") &&
				!strings.Contains(msg.Content, "!help") &&
				!strings.Contains(msg.Content, "!remaining") &&
				strings.HasPrefix(msg.Content, fmt.Sprintf("%s ", bot.IrcTrigger))
		},
		Response: func(msg *dumbirc.Message) {
			tz, err := bot.getNewYear(msg.Content[len(bot.IrcTrigger)+1:])
			if err == errNoZone || err == errNoPlace {
				log.Println("Query error:", err)
				irc.Reply(msg, fmt.Sprintf("%s: %s", msg.Name, err))
				return
			}
			if err != nil {
				log.Println("Query error:", err)
				irc.Reply(msg, fmt.Sprintf("%s: Some error occurred!", msg.Name))
				return
			}
			irc.Reply(msg, fmt.Sprintf("%s: %s", msg.Name, tz))
		},
	})
}

var (
	errNoZone  = errors.New("couldn't get the timezone for that location")
	errNoPlace = errors.New("Couldn't find that place")
)

func (bot *Settings) getNominatimReqURL(location *string) string {
	maps := url.Values{}
	maps.Add("q", *location)
	maps.Add("format", "json")
	maps.Add("accept-language", "en")
	maps.Add("limit", "1")
	maps.Add("email", bot.Email)
	return bot.Nominatim + NominatimEndpoint + maps.Encode()
}

var stNewYearWillHappen = "New Year in %s will happen in %s"
var stNewYearHappenned = "New Year in %s happened %s ago"

func (bot *Settings) getNewYear(location string) (string, error) {
	log.Println("Querying location:", location)
	data, err := NominatimGetter(bot.getNominatimReqURL(&location))
	if err != nil {
		log.Println(err)
		return "", err
	}
	var res NominatimResults
	if err = json.Unmarshal(data, &res); err != nil {
		log.Println(err)
		return "", err
	}
	if len(res) == 0 {
		return "", errNoPlace
	}
	p := gotz.Point{
		Lat: res[0].Lat,
		Lon: res[0].Lon,
	}
	zone, err := gotz.GetZone(p)
	if err != nil {
		return "", errNoZone
	}
	offset := time.Second * time.Duration(getOffset(target, zone))
	address := res[0].DisplayName

	if timeNow().UTC().Add(offset).Before(target) {
		humandur := durafmt.Parse(target.Sub(timeNow().UTC().Add(offset)))
		return fmt.Sprintf(stNewYearWillHappen, address, removeMilliseconds(humandur)), nil
	}
	humandur := durafmt.Parse(timeNow().UTC().Add(offset).Sub(target))
	return fmt.Sprintf(stNewYearHappenned, address, removeMilliseconds(humandur)), nil
}
