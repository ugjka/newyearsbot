package nyb

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/ugjka/dumbirc"
)

//Settings for bot
type Settings struct {
	IrcNick    string
	IrcChans   []string
	IrcServer  string
	IrcTrigger string
	IrcUseTLS  bool
	Bot        *dumbirc.Connection
	LogChan    LogChan
	Stopper    chan bool
	Email      string
	Nominatim  string
	extra
}

type extra struct {
	zones           TZS
	last            TZ
	next            TZ
	nominatimResult NominatimResults
	//We close this when we get WELCOME msg on join in irc
	start chan bool
	//This is used to prevent sending ping before we
	//have response from previous ping (any activity on irc)
	//pingpong(pp) sends a signal to ping timer
	pp chan bool
	sync.Once
	sync.WaitGroup
}

//New creates new bot
func New(nick string, chans []string, trigger string, server string,
	tls bool, email string, nominatim string) *Settings {
	return &Settings{
		nick,
		chans,
		server,
		trigger,
		tls,
		dumbirc.New(nick, "nyebot", server, tls),
		newLogChan(),
		make(chan bool),
		email,
		nominatim,
		extra{
			start: make(chan bool),
			pp:    make(chan bool, 1),
		},
	}
}

var stFinished = "That's it, Year %d is here AoE"

//Start starts the bot
func (s *Settings) Start() {
	log.SetOutput(s.LogChan)
	log.Println("Starting the bot...")

	defer s.Wait()
	//
	//Set up irc
	//
	bot := s.Bot
	s.addCallbacks()
	s.addTriggers()

	s.Add(1)
	go s.ircControl()

	bot.Start()

	select {
	case <-s.start:
		log.Println("Got start...")
	case <-s.Stopper:
		return
	}

	s.decodeZones(Zones)
	for {
		s.loopTimeZones()
		select {
		case <-s.Stopper:
			return
		default:
		}
		bot.PrivMsgBulk(s.IrcChans, fmt.Sprintf(stFinished, target.Year()))
		log.Println("All zones finished...")
		target = target.AddDate(1, 0, 0)
		log.Printf("Wrapping the target date around to %d\n", target.Year())
	}
}

func (s *Settings) decodeZones(z string) {
	if err := json.Unmarshal([]byte(z), &s.zones); err != nil {
		log.Println("Fatal error:", err)
		close(s.Stopper)
		return
	}
	sort.Sort(sort.Reverse(s.zones))
}

//Stop stops the bot
func (s *Settings) Stop() {
	select {
	case <-s.Stopper:
		return
	default:
		close(s.Stopper)
	}
}

var reconnectInterval = time.Second * 30
var pingInterval = time.Minute * 1

func (s *Settings) ircControl() {
	bot := s.Bot
	defer s.Done()
	for {
		timer := time.NewTimer(pingInterval)
		select {
		case err := <-bot.Errchan:
			log.Println("Error:", err)
			log.Printf("Reconnecting to irc in %s...\n", reconnectInterval)
			time.AfterFunc(reconnectInterval, func() {
				select {
				case <-s.Stopper:
					return
				default:
					bot.Start()
				}
			})
		case <-s.Stopper:
			timer.Stop()
			log.Println("Stopping the bot...")
			log.Println("Disconnecting...")
			bot.Disconnect()
			return
		case <-timer.C:
			timer.Stop()
			//pingpong stuff
			select {
			case <-s.pp:
				log.Println("Sending PING...")
				bot.Ping()
			default:
				log.Println("Got no PONG...")
			}
		}

	}
}

var stNextNewYear = "Next New Year in %s in %s"
var stHappyNewYear = "Happy New Year in %s"

func (s *Settings) loopTimeZones() {
	zones := s.zones
	bot := s.Bot
	for i := 0; i < len(zones); i++ {
		dur, err := time.ParseDuration(zones[i].Offset + "h")
		if err != nil {
			log.Println("Fatal error:", err)
			close(s.Stopper)
			return
		}
		s.next = zones[i]
		if i == 0 {
			s.last = zones[len(zones)-1]
		} else {
			s.last = zones[i-1]
		}
		if timeNow().UTC().Add(dur).Before(target) {
			time.Sleep(time.Second * 2)
			log.Println("Zone pending:", zones[i].Offset)
			humandur := durafmt.Parse(target.Sub(timeNow().UTC().Add(dur)))
			msg := fmt.Sprintf(stNextNewYear, removeMilliseconds(humandur), zones[i])
			bot.PrivMsgBulk(s.IrcChans, msg)
			//Wait till Target in Timezone
			timer := NewTimer(target.Sub(timeNow().UTC().Add(dur)))

			select {
			case <-timer.C:
				timer.Stop()
				msg = fmt.Sprintf(stHappyNewYear, zones[i])
				bot.PrivMsgBulk(s.IrcChans, msg)
				log.Println("Announcing zone:", zones[i].Offset)
			case <-s.Stopper:
				timer.Stop()
				return
			}
		}
	}
}
