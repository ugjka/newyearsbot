package nyb

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/hako/durafmt"
	irc "github.com/ugjka/dumbirc"
	c "github.com/ugjka/newyearsbot/common"
)

//Settings for bot
type Settings struct {
	IrcNick    string
	IrcChans   []string
	IrcServer  string
	IrcTrigger string
	UseTLS     bool
	LogCh      LogChan
	Stopper    chan bool
	IrcObj     *irc.Connection
	Email      string
	Nominatim  string
	extra      extra
}

type extra struct {
	zones c.TZS
	last  c.TZ
	next  c.TZ
	start chan bool
	once  sync.Once
	//This is used to prevent sending ping before we
	//have response from previous ping (any activity on irc)
	//pingpong(pp) sends a signal to ping timer
	pp   chan bool
	wait sync.WaitGroup
}

//New creates new bot
func New(nick string, chans []string, trigger string, server string, tls bool, email string, nominatim string) *Settings {
	return &Settings{
		nick,
		chans,
		server,
		trigger,
		tls,
		NewLogChan(),
		make(chan bool),
		&irc.Connection{},
		email,
		nominatim,
		extra{start: make(chan bool),
			pp: make(chan bool, 1),
		},
	}
}

//NewIrcObj return empty irc connection
func NewIrcObj() *irc.Connection {
	return &irc.Connection{}
}

//Start starts the bot
func (s *Settings) Start() {
	log.SetOutput(s.LogCh)
	log.Println("Starting the bot...")

	//To exit gracefully we need to wait
	defer s.extra.wait.Wait()
	//
	//Set up irc
	//
	s.IrcObj = irc.New(s.IrcNick, "nyebot", s.IrcServer, s.UseTLS)

	//Add Callbacs
	s.addCallbacks()
	//Add Triggers
	s.addTriggers()

	//Reconnect logic and Irc Pinger
	s.extra.wait.Add(1)
	go s.ircControl()
	//Start irc
	s.IrcObj.Start()

	//Starts when joined, see once.Do
	select {
	case <-s.extra.start:
		log.Println("Got start...")
	case <-s.Stopper:
		return
	}
	//Load timezones
	if err := json.Unmarshal([]byte(TZ), &s.extra.zones); err != nil {
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
		s.IrcObj.PrivMsgBulk(s.IrcChans, fmt.Sprintf("That's it, Year %d is here AoE", target.Year()))
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
	var err error
	defer s.extra.wait.Done()
	for {
		timer := time.NewTimer(time.Minute * 1)
		select {
		case err = <-s.IrcObj.Errchan:
			log.Println("Error:", err)
			log.Println("Restarting the bot...")
			time.AfterFunc(time.Second*30, func() {
				select {
				case <-s.Stopper:
					return
				default:
					s.IrcObj.Start()
				}
			})
		case <-s.Stopper:
			timer.Stop()
			log.Println("Stopping the bot...")
			log.Println("Disconnecting...")
			s.IrcObj.Disconnect()
			return
		//ping timer
		case <-timer.C:
			timer.Stop()
			//pingpong stuff
			select {
			case <-s.extra.pp:
				log.Println("Sending PING...")
				s.IrcObj.Ping()
			default:
				log.Println("Got no Response...")
			}
		}

	}
}

func (s *Settings) loopTimeZones() {
	for i := 0; i < len(s.extra.zones); i++ {
		dur, err := time.ParseDuration(s.extra.zones[i].Offset + "h")
		if err != nil {
			log.Fatal(err)
		}
		//Check if zone is past target
		s.extra.next = s.extra.zones[i]
		if i == 0 {
			s.extra.last = s.extra.zones[len(s.extra.zones)-1]
		} else {
			s.extra.last = s.extra.zones[i-1]
		}
		if time.Now().UTC().Add(dur).Before(target) {
			time.Sleep(time.Second * 2)
			log.Println("Zone pending:", s.extra.zones[i].Offset)
			humandur, err := durafmt.ParseString(target.Sub(time.Now().UTC().Add(dur)).String())
			if err != nil {
				log.Fatal(err)
			}
			msg := fmt.Sprintf("Next New Year in %s in %s", removeMilliseconds(humandur.String()), s.extra.zones[i])
			s.IrcObj.PrivMsgBulk(s.IrcChans, msg)
			//Wait till Target in Timezone
			timer := c.NewTimer(target.Sub(time.Now().UTC().Add(dur)))

			select {
			case <-timer.C:
				timer.Stop()
				msg = fmt.Sprintf("Happy New Year in %s", s.extra.zones[i])
				s.IrcObj.PrivMsgBulk(s.IrcChans, msg)
				log.Println("Announcing zone:", s.extra.zones[i].Offset)
			case <-s.Stopper:
				timer.Stop()
				return
			}
		}
	}
}
