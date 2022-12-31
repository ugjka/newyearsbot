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
	irc       *kitty.Bot
	extra
}

type extra struct {
	zones     TZS
	previous  TZ
	next      TZ
	remaining int
	now       time.Time
}

// New creates a new bot
func New(s *Settings) *Settings {
	s.irc = kitty.NewBot(s.Server, s.Nick,
		func(irc *kitty.Bot) {
			irc.Channels = s.Channels
			irc.Password = s.Password
			irc.SSL = s.SSL
		})
	return s
}

// LogLvl sets the log level
func (bot *Settings) LogLvl(Lvl log.Lvl) {
	logHandler := log.LvlFilterHandler(Lvl, log.StdoutHandler)
	bot.irc.Logger.SetHandler(logHandler)
}

// Start starts the bot
func (bot *Settings) Start() {
	bot.now = time.Now()
	irc := bot.irc
	irc.Info("Starting the bot...")

	bot.addTriggers()
	go bot.ircControl()

	<-irc.Joined
	irc.Info("Got start...")

	if err := bot.decodeZones(Zones); err != nil {
		irc.Crit("Decode zones error: " + err.Error())
		return
	}
	for {
		bot.loopTimeZones()
		const zonesFinishedMsg = "That's it, Year %d is here Anywhere on Earth"
		for _, ch := range irc.Channels {
			irc.Msg(ch, fmt.Sprintf(zonesFinishedMsg, target.Year()))
		}
		irc.Info("All zones finished...")
		target = target.AddDate(1, 0, 0)
		irc.Info(fmt.Sprintf("Wrapping the target date around to %d", target.Year()))
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
		if timeNow().UTC().Add(dur).Before(target) {
			time.Sleep(time.Second * 2)
			irc.Info(fmt.Sprintf("Zone pending: %.2f", zones[i].Offset))
			hdur := humanDur(target.Sub(timeNow().UTC().Add(dur)))
			const nextYearAnnounceMsg = "Next New Year in %s in %s"
			msg := fmt.Sprintf(nextYearAnnounceMsg, hdur, zones[i])
			help := fmt.Sprintf(helpMsg, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix, bot.Prefix)
			for _, ch := range irc.Channels {
				irc.Msg(ch, msg)
				irc.Msg(ch, help)
			}
			//Wait till Target in Timezone
			timer := NewTimer(target.Sub(timeNow().UTC().Add(dur)))
			select {
			case <-timer.C:
				timer.Stop()
				msg = "Happy New Year in " + zones[i].String()
				for _, ch := range irc.Channels {
					irc.Msg(ch, msg)
				}
				irc.Info(fmt.Sprintf("Announcing zone: %.2f", zones[i].Offset))
			}
		}
	}
}
