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
	IrcNick     string
	IrcChans    []string
	IrcServer   string
	IrcTrigger  string
	IrcUseTLS   bool
	IrcPassword string
	IrcConn     *dumbirc.Connection
	LogChan     LogChan
	Stopper     chan bool
	Email       string
	Nominatim   string
	extra
}

type extra struct {
	zones     TZS
	last      TZ
	next      TZ
	remaining int
	//We close this when we get WELCOME msg on join in irc
	start chan bool
	sync.Once
	sync.WaitGroup
}

//New creates new bot
func New(nick string, chans []string, password, trigger, server string,
	tls bool, email, nominatim string) *Settings {
	return &Settings{
		nick,
		chans,
		server,
		trigger,
		tls,
		password,
		dumbirc.New(nick, nick, server, tls),
		newLogChan(),
		make(chan bool),
		email,
		nominatim,
		extra{
			start: make(chan bool),
		},
	}
}

var stFinished = "That's it, Year %d is here AoE"

// Cleanup cleans up irc gouroutines if we are not reusing the bot
func (bot *Settings) Cleanup() {
	bot.Stop()
	bot.Wait()
	dumbirc.Destroy(bot.IrcConn)
}

//Start starts the bot
func (bot *Settings) Start() {
	log.SetOutput(bot.LogChan)
	log.Println("Starting the bot...")
	defer bot.Wait()
	bot.Add(1)
	//
	//Set up irc
	//
	bot.addCallbacks()
	bot.addTriggers()
	go bot.ircControl()
	irc := bot.IrcConn
	irc.RealN = "github.com/ugjka/newyearsbot"
	irc.HandleNickTaken()
	irc.HandlePingPong()
	irc.LogNotices()
	irc.SetLogOutput(bot.LogChan)
	if bot.IrcPassword != "" {
		irc.SetPassword(bot.IrcPassword)
	}
	irc.Start()

	select {
	case <-bot.start:
		log.Println("Got start...")
	case <-bot.Stopper:
		return
	}

	if err := bot.decodeZones(Zones); err != nil {
		log.Println("Fatal error:", err)
		bot.Stop()
		return
	}
	for {
		bot.loopTimeZones()
		select {
		case <-bot.Stopper:
			return
		default:
		}
		irc.MsgBulk(bot.IrcChans, fmt.Sprintf(stFinished, target.Year()))
		log.Println("All zones finished...")
		target = target.AddDate(1, 0, 0)
		log.Printf("Wrapping the target date around to %d\n", target.Year())
	}
}

func (bot *Settings) decodeZones(z []byte) error {
	if err := json.Unmarshal(z, &bot.zones); err != nil {
		return err
	}
	sort.Sort(sort.Reverse(bot.zones))
	return nil
}

//Stop stops the bot
func (bot *Settings) Stop() {
	select {
	case <-bot.Stopper:
		return
	default:
		close(bot.Stopper)
	}
}

var reconnectInterval = time.Second * 30
var pingInterval = time.Minute * 1

func (bot *Settings) ircControl() {
	irc := bot.IrcConn
	defer bot.Done()
	for {
		select {
		case err := <-irc.Errchan:
			log.Println("Error:", err)
			log.Printf("Reconnecting to irc in %s...\n", reconnectInterval)
			time.AfterFunc(reconnectInterval, func() {
				select {
				case <-bot.Stopper:
					return
				default:
					irc.Start()
				}
			})
		case <-bot.Stopper:
			log.Println("Stopping the bot...")
			log.Println("Disconnecting...")
			irc.Disconnect()
			return
		}

	}
}

var stNextNewYear = "Next New Year in %s in %s"
var stHappyNewYear = "Happy New Year in %s"

func (bot *Settings) loopTimeZones() {
	zones := bot.zones
	irc := bot.IrcConn
	for i := 0; i < len(zones); i++ {
		dur := time.Minute * time.Duration(zones[i].Offset*60)
		bot.next = zones[i]
		if i == 0 {
			bot.last = zones[len(zones)-1]
		} else {
			bot.last = zones[i-1]
		}
		bot.remaining = len(zones) - i
		if timeNow().UTC().Add(dur).Before(target) {
			time.Sleep(time.Second * 2)
			log.Println("Zone pending:", zones[i].Offset)
			humandur := durafmt.Parse(target.Sub(timeNow().UTC().Add(dur)))
			msg := fmt.Sprintf(stNextNewYear, removeMilliseconds(humandur), zones[i])
			irc.MsgBulk(bot.IrcChans, msg)
			//Wait till Target in Timezone
			timer := NewTimer(target.Sub(timeNow().UTC().Add(dur)))
			select {
			case <-timer.C:
				timer.Stop()
				msg = fmt.Sprintf(stHappyNewYear, zones[i])
				irc.MsgBulk(bot.IrcChans, msg)
				log.Println("Announcing zone:", zones[i].Offset)
			case <-bot.Stopper:
				timer.Stop()
				return
			}
		}
	}
}
