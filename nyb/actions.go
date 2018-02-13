package nyb

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hako/durafmt"
	irc "github.com/ugjka/dumbirc"
)

func (s *Settings) addCallbacks() {
	bot := s.IrcObj
	//On any message send a signal to ping timer to be ready
	bot.AddCallback(irc.ANYMESSAGE, func(msg irc.Message) {
		pingpong(s.extra.pp)
	})

	//Join channels on WELCOME
	bot.AddCallback(irc.WELCOME, func(msg irc.Message) {
		bot.Join(s.IrcChans)
		//Prevent early start
		s.extra.once.Do(func() {
			close(s.extra.start)
		})
	})
	//Reply ping messages with pong
	bot.AddCallback(irc.PING, func(msg irc.Message) {
		log.Println("PING recieved, sending PONG")
		bot.Pong()
	})
	//Log pongs
	bot.AddCallback(irc.PONG, func(msg irc.Message) {
		log.Println("Got PONG...")
	})
	//Change nick if taken
	bot.AddCallback(irc.NICKTAKEN, func(msg irc.Message) {
		log.Println("Nick taken, changing...")
		if strings.HasSuffix(bot.Nick, "_") {
			bot.Nick = bot.Nick[:len(bot.Nick)-1]
		} else {
			bot.Nick += "_"
		}
		bot.NewNick(bot.Nick)
	})
}

func (s *Settings) addTriggers() {
	bot := s.IrcObj
	//Trigger for !help
	bot.AddTrigger(irc.Trigger{
		Condition: func(msg irc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s !help", s.IrcTrigger))
		},
		Response: func(msg irc.Message) {
			bot.Reply(msg, fmt.Sprintf("%s: Query location: '%s <location>', Next zone: '%s !next', Last zone: '%s !last', Source code: https://github.com/ugjka/newyearsbot",
				msg.Name, s.IrcTrigger, s.IrcTrigger, s.IrcTrigger))
		},
	})
	//Trigger for !next
	bot.AddTrigger(irc.Trigger{
		Condition: func(msg irc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s !next", s.IrcTrigger))
		},
		Response: func(msg irc.Message) {
			log.Println("Querying !next...")
			dur, err := time.ParseDuration(s.extra.next.Offset + "h")
			if err != nil {
				return
			}
			if time.Now().UTC().Add(dur).After(target) {
				bot.Reply(msg, fmt.Sprintf("No more next, %d is here AoE", target.Year()))
				return
			}
			humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(dur)).String())
			if err != nil {
				return
			}
			bot.Reply(msg, fmt.Sprintf("Next New Year in %s in %s",
				removeMilliseconds(humandur.String()), s.extra.next.String()))
		},
	})
	//Trigger for !last
	bot.AddTrigger(irc.Trigger{
		Condition: func(msg irc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s !last", s.IrcTrigger))
		},
		Response: func(msg irc.Message) {
			log.Println("Querying !last...")
			dur, err := time.ParseDuration(s.extra.last.Offset + "h")
			if err != nil {
				return
			}
			humandur, err := durafmt.ParseString(time.Now().UTC().Add(dur).Sub(target).String())
			if err != nil {
				return
			}
			if s.extra.last.Offset == "-12" {
				humandur, err = durafmt.ParseString(time.Now().UTC().Add(dur).Sub(target.AddDate(-1, 0, 0)).String())
				if err != nil {
					return
				}
			}
			bot.Reply(msg, fmt.Sprintf("Last NewYear %s ago in %s",
				removeMilliseconds(humandur.String()), s.extra.last.String()))
		},
	})
	//Trigger for location queries
	bot.AddTrigger(irc.Trigger{
		Condition: func(msg irc.Message) bool {
			return msg.Command == "PRIVMSG" &&
				!strings.Contains(msg.Trailing, "!next") &&
				!strings.Contains(msg.Trailing, "!last") &&
				!strings.Contains(msg.Trailing, "!help") &&
				strings.HasPrefix(msg.Trailing, fmt.Sprintf("%s ", s.IrcTrigger))
		},
		Response: func(msg irc.Message) {
			tz, err := getNewYear(msg.Trailing[len(s.IrcTrigger)+1:], s.Email, s.Nominatim)
			if err != nil {
				log.Println("Query error:", err)
				bot.Reply(msg, "Some error occurred!")
				return
			}
			bot.Reply(msg, fmt.Sprintf("%s: %s", msg.Name, tz))
		},
	})
}
