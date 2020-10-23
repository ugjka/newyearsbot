package nyb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/hako/durafmt"
	kitty "github.com/ugjka/kittybot"
	"gopkg.in/ugjka/go-tz.v2/tz"
)

var stHelp = "Query location: '%s <location>', Time in location: '%s !time <location>', Next zone: '%s !next', Last zone: '%s !last', Remaining: '%s !remaining', Print this: '%s !help'"

func (bot *Settings) addTriggers() {
	irc := bot.IrcBot
	//Log Notices
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "NOTICE"
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("[NOTICE] " + m.Content)
		},
	})

	//Trigger for !help
	stSource := "Source code: https://github.com/ugjka/newyearsbot"
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(m.Content, fmt.Sprintf("%s !help", bot.IrcTrigger)) ||
				m.Content == fmt.Sprintf("%s help", bot.IrcTrigger)
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying !help...")
			b.Reply(m, fmt.Sprintf(stHelp+", "+stSource, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger))
		},
	})
	//Trigger for !next
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(m.Content, fmt.Sprintf("%s !next", bot.IrcTrigger)) ||
				m.Content == fmt.Sprintf("%s next", bot.IrcTrigger)
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying !next...")
			dur := time.Minute * time.Duration(bot.next.Offset*60)
			if timeNow().UTC().Add(dur).After(target) {
				irc.Reply(m, fmt.Sprintf("No more next, %d is here AoE", target.Year()))
				return
			}
			humandur := durafmt.Parse(target.Sub(timeNow().UTC().Add(dur)))
			irc.Reply(m, fmt.Sprintf("Next New Year in %s in %s",
				roundDuration(humandur), bot.next))
		},
	})
	//Trigger for !last
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(m.Content, fmt.Sprintf("%s !last", bot.IrcTrigger)) ||
				m.Content == fmt.Sprintf("%s last", bot.IrcTrigger)
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying !last...")
			dur := time.Minute * time.Duration(bot.last.Offset*60)
			humandur := durafmt.Parse(timeNow().UTC().Add(dur).Sub(target))
			if bot.last.Offset == -12 {
				humandur = durafmt.Parse(timeNow().UTC().Add(dur).Sub(target.AddDate(-1, 0, 0)))
			}
			irc.Reply(m, fmt.Sprintf("Last New Year %s ago in %s",
				roundDuration(humandur), bot.last))
		},
	})
	//Trigger for !remaining
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(m.Content, fmt.Sprintf("%s !remaining", bot.IrcTrigger)) ||
				m.Content == fmt.Sprintf("%s remaining", bot.IrcTrigger)
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying !remaining...")
			ss := "s"
			if bot.remaining == 1 {
				ss = ""
			}
			irc.Reply(m, fmt.Sprintf("%s: %d timezone%s remaining", m.Name, bot.remaining, ss))
		},
	})
	//Trigger for time in location
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(m.Content, fmt.Sprintf("%s !time ", bot.IrcTrigger))
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying !time...")
			res, err := bot.getTime(m.Content[len(bot.IrcTrigger)+7:])
			if err == errNoZone || err == errNoPlace {
				b.Warn("Query error: " + err.Error())
				irc.Reply(m, fmt.Sprintf("%s: %s", m.Name, err))
				return
			}
			if err != nil {
				b.Warn("Query error: " + err.Error())
				irc.Reply(m, fmt.Sprintf("%s: Some error occurred!", m.Name))
				return
			}
			irc.Reply(m, res)
		},
	})
	//UTC
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				m.Content == fmt.Sprintf("%s !time", bot.IrcTrigger) ||
				m.Content == fmt.Sprintf("%s time", bot.IrcTrigger)
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying !time...")
			res := fmt.Sprintf("Time is %s", time.Now().UTC().Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
			irc.Reply(m, res)
		},
	})

	//Trigger for incorrect querries
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				m.Content != fmt.Sprintf("%s !time", bot.IrcTrigger) &&
				!strings.HasPrefix(m.Content, fmt.Sprintf("%s !time ", bot.IrcTrigger)) &&
				!strings.HasPrefix(m.Content, fmt.Sprintf("%s !next", bot.IrcTrigger)) &&
				!strings.HasPrefix(m.Content, fmt.Sprintf("%s !last", bot.IrcTrigger)) &&
				!strings.HasPrefix(m.Content, fmt.Sprintf("%s !help", bot.IrcTrigger)) &&
				!strings.HasPrefix(m.Content, fmt.Sprintf("%s !remaining", bot.IrcTrigger)) &&
				strings.HasPrefix(m.Content, fmt.Sprintf("%s !", bot.IrcTrigger))
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			irc.Reply(m, fmt.Sprintf("Invalid command, valid commands are: "+stHelp, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger, bot.IrcTrigger))
		},
	})

	//Trigger for location queries
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				!strings.HasPrefix(m.Content, fmt.Sprintf("%s !", bot.IrcTrigger)) &&
				m.Content != fmt.Sprintf("%s next", bot.IrcTrigger) &&
				m.Content != fmt.Sprintf("%s last", bot.IrcTrigger) &&
				m.Content != fmt.Sprintf("%s help", bot.IrcTrigger) &&
				m.Content != fmt.Sprintf("%s time", bot.IrcTrigger) &&
				m.Content != fmt.Sprintf("%s remaining", bot.IrcTrigger) &&
				strings.HasPrefix(m.Content, fmt.Sprintf("%s ", bot.IrcTrigger))
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			tz, err := bot.getNewYear(m.Content[len(bot.IrcTrigger)+1:])
			if err == errNoZone || err == errNoPlace {
				b.Warn("Query error: " + err.Error())
				irc.Reply(m, fmt.Sprintf("%s: %s", m.Name, err))
				return
			}
			if err != nil {
				b.Warn("Query error: " + err.Error())
				irc.Reply(m, fmt.Sprintf("%s: Some error occurred!", m.Name))
				return
			}
			irc.Reply(m, fmt.Sprintf("%s: %s", m.Name, tz))
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

func (bot *Settings) getTime(location string) (string, error) {
	bot.IrcBot.Info("Querying location: " + location)
	data, err := NominatimGetter(bot.getNominatimReqURL(&location))
	if err != nil {
		bot.IrcBot.Warn("Nominatim error: " + err.Error())
		return "", err
	}
	var res NominatimResults
	if err = json.Unmarshal(data, &res); err != nil {
		bot.IrcBot.Warn("Nominatim error: " + err.Error())
		return "", err
	}
	if len(res) == 0 {
		return "", errNoPlace
	}
	p := tz.Point{
		Lat: res[0].Lat,
		Lon: res[0].Lon,
	}
	tzid, err := tz.GetZone(p)
	if err != nil {
		return "", errNoZone
	}
	zone, err := time.LoadLocation(tzid[0])
	if err != nil {
		return "", errNoZone
	}
	address := res[0].DisplayName
	return fmt.Sprintf("Time in %s is %s", address, time.Now().In(zone).Format("Mon Jan 2 15:04:05 -0700 MST 2006")), nil
}

func (bot *Settings) getNewYear(location string) (string, error) {
	bot.IrcBot.Info("Querying location: " + location)
	data, err := NominatimGetter(bot.getNominatimReqURL(&location))
	if err != nil {
		bot.IrcBot.Warn("Nominatim error: " + err.Error())
		return "", err
	}
	var res NominatimResults
	if err = json.Unmarshal(data, &res); err != nil {
		bot.IrcBot.Warn("Nominatim error: " + err.Error())
		return "", err
	}
	if len(res) == 0 {
		return "", errNoPlace
	}
	p := tz.Point{
		Lat: res[0].Lat,
		Lon: res[0].Lon,
	}
	tzid, err := tz.GetZone(p)
	if err != nil {
		return "", errNoZone
	}
	zone, err := time.LoadLocation(tzid[0])
	if err != nil {
		return "", errNoZone
	}
	offset := time.Second * time.Duration(getOffset(target, zone))
	address := res[0].DisplayName

	if timeNow().UTC().Add(offset).Before(target) {
		humandur := durafmt.Parse(target.Sub(timeNow().UTC().Add(offset)))
		return fmt.Sprintf(stNewYearWillHappen, address, roundDuration(humandur)), nil
	}
	humandur := durafmt.Parse(timeNow().UTC().Add(offset).Sub(target))
	return fmt.Sprintf(stNewYearHappenned, address, roundDuration(humandur)), nil
}
