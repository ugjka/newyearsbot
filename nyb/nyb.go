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
	IrcConn    *dumbirc.Connection
	LogChan    LogChan
	Stopper    chan bool
	Email      string
	Nominatim  string
	extra      extra
}

type extra struct {
	zones TZS
	last  TZ
	next  TZ
	start chan bool
	once  sync.Once
	//This is used to prevent sending ping before we
	//have response from previous ping (any activity on irc)
	//pingpong(pp) sends a signal to ping timer
	pp   chan bool
	wait sync.WaitGroup
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

//Start starts the bot
func (s *Settings) Start() {
	log.SetOutput(s.LogChan)
	log.Println("Starting the bot...")

	//To exit gracefully we need to wait
	defer s.extra.wait.Wait()
	//
	//Set up irc
	//
	bot := s.IrcConn
	//Add Callbacs
	s.addCallbacks()
	//Add Triggers
	s.addTriggers()

	//Reconnect logic and Irc Pinger
	s.extra.wait.Add(1)
	go s.ircControl()
	//Start irc
	bot.Start()

	//Starts when joined, see once.Do
	select {
	case <-s.extra.start:
		log.Println("Got start...")
	case <-s.Stopper:
		return
	}
	//Load timezones
	if err := json.Unmarshal([]byte(Zones), &s.extra.zones); err != nil {
		log.Fatal(err)
	}
	//Sort them
	sort.Sort(sort.Reverse(s.extra.zones))

	//Zone Looper
	for {
		s.loopTimeZones()
		select {
		case <-s.Stopper:
			return
		default:
		}
		bot.PrivMsgBulk(s.IrcChans, fmt.Sprintf("That's it, Year %d is here AoE", target.Year()))
		log.Println("All zones finished...")
		target = target.AddDate(1, 0, 0)
		log.Printf("Wrapping target date around to %d\n", target.Year())
	}
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

func (s *Settings) ircControl() {
	bot := s.IrcConn
	var err error
	defer s.extra.wait.Done()
	for {
		timer := time.NewTimer(time.Minute * 1)
		select {
		case err = <-bot.Errchan:
			log.Println("Error:", err)
			log.Println("Recconecting to irc in 30secs...")
			time.AfterFunc(time.Second*30, func() {
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
		//ping timer
		case <-timer.C:
			timer.Stop()
			//pingpong stuff
			select {
			case <-s.extra.pp:
				log.Println("Sending PING...")
				bot.Ping()
			default:
				log.Println("Got no PONG...")
			}
		}

	}
}

func (s *Settings) loopTimeZones() {
	zones := s.extra.zones
	bot := s.IrcConn
	for i := 0; i < len(zones); i++ {
		dur, err := time.ParseDuration(zones[i].Offset + "h")
		if err != nil {
			log.Fatal(err)
		}
		//Check if zone is past target
		s.extra.next = zones[i]
		if i == 0 {
			s.extra.last = zones[len(zones)-1]
		} else {
			s.extra.last = zones[i-1]
		}
		if time.Now().UTC().Add(dur).Before(target) {
			time.Sleep(time.Second * 2)
			log.Println("Zone pending:", zones[i].Offset)
			humandur := durafmt.Parse(target.Sub(time.Now().UTC().Add(dur)))
			msg := fmt.Sprintf("Next New Year in %s in %s", removeMilliseconds(humandur), zones[i])
			bot.PrivMsgBulk(s.IrcChans, msg)
			//Wait till Target in Timezone
			timer := NewTimer(target.Sub(time.Now().UTC().Add(dur)))

			select {
			case <-timer.C:
				timer.Stop()
				msg = fmt.Sprintf("Happy New Year in %s", zones[i])
				bot.PrivMsgBulk(s.IrcChans, msg)
				log.Println("Announcing zone:", zones[i].Offset)
			case <-s.Stopper:
				timer.Stop()
				return
			}
		}
	}
}
