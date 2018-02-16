package nyb

import (
	"testing"

	"github.com/ugjka/dumbirc"
)

func TestTriggers(t *testing.T) {
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.addTriggers()
	nye.decodeZones()
	cases := []string{"hny !next", "hny !last", "hny !help"}
	for _, v := range cases {
		m := dumbirc.NewMessage()
		m.Command = "PRIVMSG"
		m.Trailing = v
		m.Name = "test"
		nye.Bot.RunTriggers(m)
	}
}
