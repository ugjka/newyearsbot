package main

import (
	"encoding/json"
	"log"
	"time"

	kitty "github.com/ugjka/kittybot"
	"github.com/ugjka/newyearsbot/nyb"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	bot := kitty.NewBot(
		"testnet.ergo.chat:6697",
		//"irc.libera.chat:6697",
		"happyhappy2025v2",
		func(b *kitty.Bot) {
			b.Channels = []string{"##ugjka2"}
			b.SSL = true
		},
	)
	var zones nyb.TZS
	err := json.Unmarshal(nyb.Zones, &zones)
	if err != nil {
		log.Fatal(err)
	}
	bot.AddTrigger(kitty.Trigger{
		Condition: func(b *kitty.Bot, m *kitty.Message) bool {
			return m.Content == "!test"
		},
		Action: func(b *kitty.Bot, m *kitty.Message) {
			for _, z := range zones {
				const pre = "\x02\x0302Next New Year\x0f in \x02\x030213 seconds 323 milliseconds\x0f in "
				b.Reply(m, pre+z.Format(b.MsgMaxSize(m.To)-len(pre), true))
				time.Sleep(time.Second)
				b.Reply(m, "**************************")
				time.Sleep(time.Second * 3)
			}
		},
	})
	bot.Logger.SetHandler(log15.StdoutHandler)
	bot.Run()
}
