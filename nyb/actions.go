package nyb

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	kitty "github.com/ugjka/kittybot"
	"gopkg.in/ugjka/go-tz.v2/tz"
)

const helpMsg = "COMMANDS: '%shny <location>', '%stime <location>', '%snext', '%sprevious', '%sremaining', '%shelp', '%ssource'"

func (bot *Settings) addTriggers() {
	irc := bot.irc

	//Log Notices
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "NOTICE"
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("[NOTICE] " + m.Content)
		},
	})

	//Trigger for !source
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(normalize(m.Content), bot.Prefix+"source")
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Reply(m, "https://github.com/ugjka/newyearsbot")
		},
	})

	//Trigger for !help
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(normalize(m.Content), bot.Prefix+"help") ||
				normalize(m.Content) == bot.Prefix+"hny"
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying help...")
			b.Reply(m, fmt.Sprintf(helpMsg, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix))
		},
	})

	//Trigger for !next
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(normalize(m.Content), bot.Prefix+"next")
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying next...")
			dur := time.Minute * time.Duration(bot.next.Offset*60)
			if now().UTC().Add(dur).After(target) {
				b.Reply(m, fmt.Sprintf("No more next, %d is here AoE", target.Year()))
				return
			}
			hdur := humanDur(target.Sub(now().UTC().Add(dur)))
			const next = "Next New Year in "
			max := b.ReplyMaxSize(m)
			max -= len(next)
			max -= len(hdur)
			max -= 4
			b.Reply(m, next+hdur+" in "+bot.next.Split(max))
		},
	})

	//Trigger for !previous
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(normalize(m.Content), bot.Prefix+"prev")
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying previous...")
			dur := time.Minute * time.Duration(bot.previous.Offset*60)
			hdur := humanDur(now().UTC().Add(dur).Sub(target))
			if bot.previous.Offset == -12 {
				hdur = humanDur(now().UTC().Add(dur).Sub(target.AddDate(-1, 0, 0)))
			}
			const prev = "Previous New Year "
			max := b.ReplyMaxSize(m)
			max -= len(prev)
			max -= len(hdur)
			max -= 8
			b.Reply(m, prev+hdur+" ago in "+bot.previous.Split(max))
		},
	})

	//Trigger for !remaining
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(normalize(m.Content), bot.Prefix+"remaining")
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying remaining...")
			plural := "s"
			if bot.remaining == 1 {
				plural = ""
			}
			pct := ((float64(len(bot.zones)) - float64(bot.remaining)) / float64(len(bot.zones)) * 100)
			b.Reply(m, fmt.Sprintf("%d timezone%s remaining. %.2f%% are in the new year", bot.remaining, plural, pct))
		},
	})

	//Trigger for time in location
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(normalize(m.Content), bot.Prefix+"time ")
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying time...")
			result, err := bot.time(normalize(m.Content)[len(bot.Prefix)+len("time")+1:])
			if err == errNoZone || err == errNoPlace {
				b.Warn("Query error: " + err.Error())
				b.Reply(m, err.Error())
				return
			}
			if err != nil {
				b.Warn("Query error: " + err.Error())
				b.Reply(m, "Some error occurred!")
				return
			}
			b.Reply(m, result)
		},
	})

	//Trigger for UTC time
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				normalize(m.Content) == bot.Prefix+"time"
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			b.Info("Querying time...")
			result := "Time is " + now().UTC().Format("Mon Jan 2 15:04:05 -0700 MST 2006")
			b.Reply(m, result)
		},
	})

	//Trigger for new year in location
	irc.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Command == "PRIVMSG" &&
				strings.HasPrefix(normalize(m.Content), bot.Prefix+"hny ")
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			result, err := bot.newYear(normalize(m.Content)[len(bot.Prefix)+len("hny")+1:])
			if err == errNoZone || err == errNoPlace {
				b.Warn("Query error: " + err.Error())
				b.Reply(m, err.Error())
				return
			}
			if err != nil {
				b.Warn("Query error: " + err.Error())
				b.Reply(m, "Some error occurred!")
				return
			}

			b.Reply(m, result)
		},
	})
}

var (
	errNoZone  = errors.New("couldn't get timezone for that location")
	errNoPlace = errors.New("couldn't find that place")
)

func (bot *Settings) time(location string) (string, error) {
	bot.irc.Info("Querying location: " + location)
	data, err := NominatimFetcher(&bot.Email, &bot.Nominatim, &location)
	if err != nil {
		bot.irc.Warn("Nominatim error: " + err.Error())
		return "", err
	}
	var res NominatimResults
	if err = json.Unmarshal(data, &res); err != nil {
		bot.irc.Warn("Nominatim JSON error: " + err.Error())
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
	msg := fmt.Sprintf("Time in %s is %s", address, now().In(zone).Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	return msg, nil
}

func (bot *Settings) newYear(location string) (string, error) {
	bot.irc.Info("Querying location: " + location)
	data, err := NominatimFetcher(&bot.Email, &bot.Nominatim, &location)
	if err != nil {
		bot.irc.Warn("Nominatim error: " + err.Error())
		return "", err
	}
	var res NominatimResults
	if err = json.Unmarshal(data, &res); err != nil {
		bot.irc.Warn("Nominatim JSON error: " + err.Error())
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
	offset := zoneOffset(target, zone)
	address := res[0].DisplayName
	if now().UTC().Add(offset).Before(target) {
		hdur := humanDur(target.Sub(now().UTC().Add(offset)))
		const newYearFutureMsg = "New Year in %s will happen in %s"
		return fmt.Sprintf(newYearFutureMsg, address, hdur), nil
	}
	hdur := humanDur(now().UTC().Add(offset).Sub(target))
	const newYearPastMsg = "New Year in %s happened %s ago"
	return fmt.Sprintf(newYearPastMsg, address, hdur), nil
}
