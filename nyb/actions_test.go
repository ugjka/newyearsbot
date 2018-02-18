package nyb

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ugjka/dumbirc"
)

func TestTriggers(t *testing.T) {
	//hny !next, hny !last
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.addTriggers()
	nye.decodeZones(&Zones)
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

func TestQuery(t *testing.T) {
	go createServer()
	time.Sleep(time.Second)
	nye := New("", []string{""}, "hny", "", false, "", "http://127.0.0.1:1234")
	nye.addTriggers()
	//Ok location
	m := dumbirc.NewMessage()
	m.Command = dumbirc.PRIVMSG
	m.Trailing = "hny ok"
	m.Name = "test"
	nye.Bot.RunTriggers(m)
	//Test cached
	nye.Bot.RunTriggers(m)
	//Not ok server status
	m = dumbirc.NewMessage()
	m.Command = dumbirc.PRIVMSG
	m.Trailing = "hny notok"
	m.Name = "test"
	nye.Bot.RunTriggers(m)
	//Test No time zone for location
	m = dumbirc.NewMessage()
	m.Command = dumbirc.PRIVMSG
	m.Trailing = "hny nozone"
	m.Name = "test"
	nye.Bot.RunTriggers(m)
	//Test no place found
	m = dumbirc.NewMessage()
	m.Command = dumbirc.PRIVMSG
	m.Trailing = "hny noplace"
	m.Name = "test"
	nye.Bot.RunTriggers(m)
	//Test malformed json
	m = dumbirc.NewMessage()
	m.Command = dumbirc.PRIVMSG
	m.Trailing = "hny borked"
	m.Name = "test"
	nye.Bot.RunTriggers(m)
	//Test past
	timeNow = func() time.Time {
		return time.Date(time.Now().Year()+1000, 0, 0, 0, 0, 0, 0, time.UTC)
	}
	m = dumbirc.NewMessage()
	m.Command = dumbirc.PRIVMSG
	m.Trailing = "hny ok"
	m.Name = "test"
	nye.Bot.RunTriggers(m)
	//BORKED SERVER
	nye.Nominatim = "//////////////"
	nye.Bot.RunTriggers(m)
	nye.Nominatim = ":"
	nye.Bot.RunTriggers(m)

}

func createServer() {
	mux := httprouter.New()
	mux.HandlerFunc("GET", "/search", fakeNominatim)
	http.ListenAndServe(":1234", mux)
	return
}

func fakeNominatim(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	//test ok
	if values.Get("q") == "ok" {
		type nom struct {
			Lat         string
			Lon         string
			DisplayName string `json:"Display_name"`
		}
		type noms []nom
		n := make(noms, 0)
		n = append(n, nom{
			DisplayName: "ok",
			Lat:         "56.946285",
			Lon:         "24.105078",
		})
		err := json.NewEncoder(w).Encode(n)
		if err != nil {
			log.Fatal(err)
		}
	}
	//test not ok
	if values.Get("q") == "notok" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//test no zone
	if values.Get("q") == "nozone" {
		type nom struct {
			Lat         string
			Lon         string
			DisplayName string `json:"Display_name"`
		}
		type noms []nom
		n := make(noms, 0)
		n = append(n, nom{
			DisplayName: "ok",
			Lat:         "0.190165906",
			Lon:         "-176.474331436",
		})
		err := json.NewEncoder(w).Encode(n)
		if err != nil {
			log.Fatal(err)
		}
	}
	//test no place
	if values.Get("q") == "noplace" {
		type nom struct {
			Lat         string
			Lon         string
			DisplayName string `json:"Display_name"`
		}
		type noms []nom
		n := make(noms, 0)
		err := json.NewEncoder(w).Encode(n)
		if err != nil {
			log.Fatal(err)
		}
	}
	//borked json
	if values.Get("q") == "borked" {
		w.Write([]byte("*eaeia"))
	}
}
