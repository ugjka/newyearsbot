package nyb

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	kitty "github.com/ugjka/kittybot"
	log "gopkg.in/inconshreveable/log15.v2"
)

// Settings for the bot
type Settings struct {
	Nick      string
	Channels  []string
	Server    string
	SSL       bool
	Password  string
	Prefix    string
	Email     string
	Nominatim string
	Limit     bool
	Colors    bool
	irc       *kitty.Bot
	extra
}

type extra struct {
	zones     TZS
	previous  TZ
	next      TZ
	remaining int
	first     bool
	target    time.Time
}

// New creates a new bot
func New(s *Settings) *Settings {
	s.irc = kitty.NewBot(s.Server, s.Nick,
		func(irc *kitty.Bot) {
			irc.Channels = s.Channels
			irc.Password = s.Password
			irc.SSL = s.SSL
			irc.LimitReplies = s.Limit
		})
	return s
}

// LogLvl sets the log level
func (bot *Settings) LogLvl(Lvl log.Lvl) {
	logHandler := log.LvlFilterHandler(Lvl, log.StderrHandler)
	bot.irc.Logger.SetHandler(logHandler)
}

// Start starts the bot
func (bot *Settings) Start() {
	bot.target = target
	irc := bot.irc
	irc.Info("Starting the bot...")

	bot.addTriggers()
	go bot.ircControl()

	<-irc.Joined
	// Neet to wait a bit for prefix
	time.Sleep(time.Second * 5)
	irc.Info("Got start...")

	if err := bot.decodeZones(Zones); err != nil {
		irc.Crit("Decode zones error: " + err.Error())
		return
	}
	for {
		bot.loopTimeZones()
		var zonesFinishedMsg = bot.col("That's it") + ", Year " +
			bot.col("%d") + " is here " +
			bot.col("Anywhere on Earth")
		for _, ch := range irc.Channels {
			irc.Msg(ch, fmt.Sprintf(zonesFinishedMsg, bot.target.Year()))
		}
		irc.Info("All zones finished...")
		bot.target = bot.target.AddDate(1, 0, 0)
		irc.Info(fmt.Sprintf("Wrapping the target date around to %d", bot.target.Year()))
	}
}

func (bot *Settings) decodeZones(z []byte) error {
	if err := json.Unmarshal(z, &bot.zones); err != nil {
		return err
	}
	sort.Sort(sort.Reverse(bot.zones))
	return nil
}

const reconnectInterval = time.Second * 30

func (bot *Settings) ircControl() {
	irc := bot.irc
	for {
		irc.Run()
		irc.Info("Reconnecting...")
		time.Sleep(reconnectInterval)
	}
}

func (bot *Settings) loopTimeZones() {
	zones := bot.zones
	irc := bot.irc
	for i := 0; i < len(zones); i++ {
		dur := time.Minute * time.Duration(zones[i].Offset*60)
		bot.next = zones[i]
		if i == 0 {
			bot.previous = zones[len(zones)-1]
		} else {
			bot.previous = zones[i-1]
		}
		bot.remaining = len(zones) - i
		if now().UTC().Add(dur).Before(bot.target) {
			time.Sleep(time.Second * 2)
			irc.Info(fmt.Sprintf("Zone pending: %.2f", zones[i].Offset))
			hdur := humanDur(bot.target.Sub(now().UTC().Add(dur)))
			hdur = bot.col(hdur)
			next := bot.col("Next New Year") + " in "
			if i == 0 && !(now().Month() == time.January && now().Day() < 2) {
				next = bot.col("First New Year") + " in "
			}
			if i == len(zones)-1 {
				next = bot.col("Final New Year") + " in "
			}
			help := fmt.Sprintf(helpMsg, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix)
			for _, ch := range irc.Channels {
				max := irc.MsgMaxSize(ch)
				max -= len(next)
				max -= len(hdur)
				max -= 4
				if !bot.first {
					irc.Msg(ch, next+hdur+" in "+zones[i].Format(max, bot.Colors))
					irc.Msg(ch, help)
					bot.first = true
				} else {
					irc.Msg(ch, next+hdur+". "+
						fmt.Sprintf("See %snext or %shelp.", bot.Prefix, bot.Prefix))
				}
			}
			//Wait till Target in Timezone
			timer := NewTimer(bot.target.Sub(now().UTC().Add(dur)))
			<-timer.C
			timer.Stop()
			var happy = bot.col("Happy New Year") + " in "
			for _, ch := range irc.Channels {
				max := irc.MsgMaxSize(ch)
				max -= len(happy)
				irc.Msg(ch, happy+zones[i].Format(max, bot.Colors))
			}
			irc.Info(fmt.Sprintf("Announcing zone: %.2f", zones[i].Offset))
		}
	}
}

// https://modern.ircdocs.horse/formatting.html
func (bot *Settings) col(s string) string {
	if bot.Colors {
		s = "\x02\x0302" + s + "\x0f"
	}
	return s
}
