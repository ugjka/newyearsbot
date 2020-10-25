package nyb

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hako/durafmt"
	kitty "github.com/ugjka/kittybot"
	log "gopkg.in/inconshreveable/log15.v2"
)

//Settings for the bot
type Settings struct {
	Prefix    string
	IRC       *kitty.Bot
	Email     string
	Nominatim string
	extra
}

type extra struct {
	zones     TZS
	previous  TZ
	next      TZ
	remaining int
}

//New creates a new bot
func New(nick string, channels []string, password, prefix, server string,
	ssl bool, email, nominatim string) *Settings {
	prefix = strings.ToLower(prefix)
	return &Settings{
		prefix,
		kitty.NewBot(server, nick, func(bot *kitty.Bot) {
			bot.Channels = channels
			bot.Password = password
			bot.SSL = ssl
		}),
		email,
		nominatim,
		extra{},
	}
}

// LogLvl sets log level
func (bot *Settings) LogLvl(Lvl log.Lvl) {
	logHandler := log.LvlFilterHandler(Lvl, log.StdoutHandler)
	bot.IRC.Logger.SetHandler(logHandler)
}

//Start starts the bot
func (bot *Settings) Start() {
	irc := bot.IRC
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
	irc := bot.IRC
	for {
		irc.Run()
		irc.Info("Reconnecting...")
		time.Sleep(reconnectInterval)
	}
}

func (bot *Settings) loopTimeZones() {
	zones := bot.zones
	irc := bot.IRC
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
			humandur := durafmt.Parse(target.Sub(timeNow().UTC().Add(dur)))
			const nextYearAnnounceMsg = "Next New Year in %s in %s"
			msg := fmt.Sprintf(nextYearAnnounceMsg, roundDuration(humandur), zones[i])
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
