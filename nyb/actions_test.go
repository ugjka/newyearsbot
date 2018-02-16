package nyb

import (
	"testing"
	"time"

	"github.com/ugjka/dumbirc"
)

func TestTriggers(t *testing.T) {
	//hny !next, hny !last
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.addTriggers()
	nye.decodeZones(Zones)
	nye.last = nye.zones[len(nye.zones)-1]
	nye.next = nye.zones[len(nye.zones)-2]
	cases := []string{"hny !next", "hny !last"}
	offsets := []time.Duration{time.Hour * -5, time.Hour * 5, time.Hour * 24, time.Hour * -24}
	for _, v := range offsets {
		timeNow = func() time.Time {
			return target.Add(v)
		}
		for _, v := range cases {
			m := dumbirc.NewMessage()
			m.Command = dumbirc.PRIVMSG
			m.Trailing = v
			m.Name = "test"
			nye.Bot.RunTriggers(m)
		}
	}
	//test borked offsets
	nye.last.Offset = "aoeiaoi"
	nye.next.Offset = "aoeoaei"
	for _, v := range cases {
		m := dumbirc.NewMessage()
		m.Command = dumbirc.PRIVMSG
		m.Trailing = v
		m.Name = "test"
		nye.Bot.RunTriggers(m)
	}
	//hny !help
	m := dumbirc.NewMessage()
	m.Command = dumbirc.PRIVMSG
	m.Trailing = "hny !help"
	m.Name = "test"
	nye.Bot.RunTriggers(m)
}

func TestCallbacks(t *testing.T) {
	nickChangeInterval = time.Second * 0
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.addCallbacks()
	cases := []string{
		dumbirc.ANYMESSAGE, dumbirc.NICKTAKEN, dumbirc.PING,
		dumbirc.PONG, dumbirc.PRIVMSG, dumbirc.WELCOME,
	}
	for _, v := range cases {
		m := dumbirc.NewMessage()
		m.Command = v
		nye.Bot.RunCallbacks(m)
	}
}
