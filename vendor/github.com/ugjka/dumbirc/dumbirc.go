package dumbirc

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ugjka/dumbirc/messenger"

	irc "gopkg.in/sorcix/irc.v2"
)

var replacer *strings.Replacer

func init() {
	replacer = strings.NewReplacer("\n", " ", "\t", " ", "\a", " ", "\b", " ", "\f", " ", "\r", " ", "\v", " ")
}

//Map event codes
const (
	PRIVMSG   = irc.PRIVMSG
	PING      = irc.PING
	PONG      = irc.PONG
	WELCOME   = irc.RPL_WELCOME
	NICKTAKEN = irc.ERR_NICKNAMEINUSE
	JOIN      = irc.JOIN
	KICK      = irc.KICK
	NOTICE    = irc.NOTICE
	//Useful if you wanna check for activity
	ANYMESSAGE = "ANY"
)

//Connection Settings
type Connection struct {
	Nick         string
	User         string
	RealN        string
	Server       string
	TLS          bool
	Password     string
	Throttle     time.Duration
	connectedSet chan bool
	connectedGet chan bool
	//Fake Connected status
	DebugFakeConn bool
	conn          *irc.Conn
	callbacks     map[string][]func(*Message)
	triggers      []Trigger
	Log           *log.Logger
	Debug         *log.Logger
	Errchan       chan error
	Send          chan string
	prefix        *irc.Prefix
	messenger     *messenger.Messenger
	prefixlenGet  chan int
	prefixlenSet  chan []string
	destroy       chan struct{}
	sync.WaitGroup
}

//New creates a new irc object
func New(nick, user, server string, tls bool) *Connection {
	conn := &Connection{
		Nick:         nick,
		User:         user,
		Server:       server,
		TLS:          tls,
		Throttle:     time.Millisecond * 500,
		conn:         &irc.Conn{},
		callbacks:    make(map[string][]func(*Message)),
		triggers:     make([]Trigger, 0),
		Log:          log.New(&devNull{}, "", log.Ldate|log.Ltime),
		Debug:        log.New(&devNull{}, "debug", log.Ltime),
		Errchan:      make(chan error),
		WaitGroup:    sync.WaitGroup{},
		prefix:       new(irc.Prefix),
		connectedGet: make(chan bool),
		connectedSet: make(chan bool),
		prefixlenGet: make(chan int),
		prefixlenSet: make(chan []string),
		destroy:      make(chan struct{}),
	}
	conn.getPrefix()
	conn.prefix.Name = nick
	go connStatusMon(conn)
	go prefixMonitor(conn)
	return conn
}

// Destroy terminates monitor goroutines created by New()
func Destroy(c *Connection) {
	close(c.destroy)
}

func connStatusMon(c *Connection) {
	connected := false
	for {
		select {
		case connected = <-c.connectedSet:
		case c.connectedGet <- connected:
		case <-c.destroy:
			return
		}
	}
}

func prefixMonitor(c *Connection) {
	for {
		select {
		case args := <-c.prefixlenSet:
			if args[0] != "" {
				c.prefix.Name = args[0]
			}
			if args[1] != "" {
				c.prefix.User = args[1]
			}
			if args[2] != "" {
				c.prefix.Host = args[2]
			}
		case c.prefixlenGet <- c.prefix.Len():
		case <-c.destroy:
			return
		}
	}
}

//Message event
type Message struct {
	*irc.Message
	Content   string
	TimeStamp time.Time
	To        string
}

//ParseMessage converts irc.Message to Message
func ParseMessage(raw *irc.Message) (m *Message) {
	m = new(Message)
	m.Message = raw
	m.Content = m.Trailing()
	if len(m.Params) > 0 {
		m.To = m.Params[0]
	} else if m.Command == JOIN {
		m.To = m.Trailing()
	}
	m.TimeStamp = time.Now()
	return m
}

//NewMessage returns an empty message
func NewMessage() *Message {
	msg := new(irc.Message)
	msg.Prefix = new(irc.Prefix)
	msg.Params = make([]string, 0)
	return &Message{Message: msg, TimeStamp: time.Now()}
}

type devNull struct {
}

func (d *devNull) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// WaitFor will block until a message matching the given filter is received
func (c *Connection) WaitFor(filter func(*Message) bool, cmd func(), timeout time.Duration, timeoutErr error) (err error) {
	notConnErr := fmt.Errorf("WaitFor: exiting, not connected")
	if !c.IsConnected() {
		return notConnErr
	}
	cmd()
	timer := time.NewTimer(timeout)
	client, err := c.messenger.Sub()
	if err != nil {
		return notConnErr
	}
	defer func() {
		c.messenger.Unsub(client)
	}()
	for {
		select {
		case mes, ok := <-client:
			if !ok {
				timer.Stop()
				return notConnErr
			}
			if filter(mes.(*Message)) {
				timer.Stop()
				return nil
			}
		case <-timer.C:
			if !c.IsConnected() {
				return notConnErr
			}
			return timeoutErr
		}
	}
}

//SetThrottle sets post delay
func (c *Connection) SetThrottle(d time.Duration) {
	c.Throttle = d
}

//SetPassword sets the irc password
func (c *Connection) SetPassword(pass string) {
	c.Password = pass
}

//SetLogOutput sets logger writer
func (c *Connection) SetLogOutput(w io.Writer) {
	c.Log.SetOutput(w)
}

//SetDebugOutput sets debug logger writer
func (c *Connection) SetDebugOutput(w io.Writer) {
	c.Debug.SetOutput(w)
}

//IsConnected returns connection status
func (c *Connection) IsConnected() bool {
	return <-c.connectedGet
}

//AddCallback Adds callback to an event
func (c *Connection) AddCallback(event string, callback func(*Message)) {
	c.callbacks[event] = append(c.callbacks[event], callback)
}

//Trigger scheme
type Trigger struct {
	Condition func(*Message) bool
	Response  func(*Message)
}

//AddTrigger adds triggers
func (c *Connection) AddTrigger(t Trigger) {
	c.triggers = append(c.triggers, t)
}

//RunTriggers ...
func (c *Connection) RunTriggers(m *Message) {
	for _, v := range c.triggers {
		if v.Condition(m) {
			v.Response(m)
		}
	}
}

//RunCallbacks ...
func (c *Connection) RunCallbacks(m *Message) {
	if v, ok := c.callbacks[ANYMESSAGE]; ok {
		for _, v := range v {
			v(m)
		}
	}
	if v, ok := c.callbacks[m.Command]; ok {
		for _, v := range v {
			v(m)
		}
	}
}

func (c *Connection) send(msg string) {
	if !c.IsConnected() {
		return
	}
	c.Send <- msg
}

//Join channels
func (c *Connection) Join(channels []string) {
	for _, v := range channels {
		c.send(irc.JOIN + " " + v)
	}
}

// ChMode is used to change users modes in a channel
// operator = "+o" deop = "-o"
// ban = "+b"
func (c *Connection) ChMode(user, channel, mode string) {
	c.send("MODE " + channel + " " + mode + " " + user)
}

// Topic sets the channel 'channel' topic (requires bot has proper permissions)
func (c *Connection) Topic(channel, topic string) {
	msg := fmt.Sprintf("TOPIC %s :%s", channel, topic)
	c.send(msg)
}

// Action sends an action to 'dest' (user or channel)
func (c *Connection) Action(dest, msg string) {
	msg = fmt.Sprintf("\u0001ACTION %s\u0001", msg)
	c.Msg(dest, msg)
}

// Notice sends a NOTICE message to 'dest' (user or channel)
func (c *Connection) Notice(dest, msg string) {
	msg = replacer.Replace(msg)
	prefLen := 2 + <-c.prefixlenGet + len("NOTICE "+dest+" :")
	for prefLen+len(msg) > 510 {
		c.send("NOTICE " + dest + " :" + msg[:510-prefLen])
		msg = msg[510-prefLen:]
	}
	c.send("NOTICE " + dest + " :" + msg)
}

//Pong sends pong
func (c *Connection) Pong() {
	c.send(irc.PONG)
}

//Ping sends ping
func (c *Connection) Ping() {
	c.send(irc.PING + " " + c.Server)
}

//Cmd sends command
func (c *Connection) Cmd(command string) {
	c.send(command)
}

//Msg sends privmessage
func (c *Connection) Msg(dest, msg string) {
	msg = replacer.Replace(msg)
	prefLen := 2 + <-c.prefixlenGet + len(irc.PRIVMSG+" "+dest+" :")
	for prefLen+len(msg) > 510 {
		c.send(irc.PRIVMSG + " " + dest + " :" + msg[:510-prefLen])
		msg = msg[510-prefLen:]
	}
	c.send(irc.PRIVMSG + " " + dest + " :" + msg)
}

//MsgBulk sends message to many
func (c *Connection) MsgBulk(dest []string, msg string) {
	for _, k := range dest {
		c.Msg(k, msg)
	}
}

//NewNick Changes nick
func (c *Connection) NewNick(nick string) {
	c.send(irc.NICK + " " + nick)
	c.prefixlenSet <- []string{nick, "", ""}
}

//Reply replies to a message
func (c *Connection) Reply(m *Message, reply string) {
	if m.To == c.Nick {
		c.Msg(m.Name, reply)
	} else {
		c.Msg(m.To, reply)
	}
}

//Disconnect disconnects from irc
func (c *Connection) Disconnect() {
	if !c.IsConnected() {
		return
	}
	c.connectedSet <- false
	c.conn.Close()
	c.messenger.Kill()
	for {
		select {
		case <-c.Send:
		default:
			close(c.Send)
			return
		}
	}
}

func changeNick(nick string) string {
	if len(nick) < 16 {
		nick += "_"
		return nick
	}
	nick = strings.TrimRight(nick, "_")
	if len(nick) > 12 {
		nick = nick[:12] + "_"
	}
	return nick
}

//LogNotices logs notice messages
func (c *Connection) LogNotices() {
	c.AddCallback(NOTICE, func(m *Message) {
		c.Log.Printf("NOTICE %s %s", m.To, m.Content)
	})
}

//HandleNickTaken changes nick when nick taken
func (c *Connection) HandleNickTaken() {
	c.AddCallback(NICKTAKEN, func(msg *Message) {
		if c.Password != "" {
			rand.Seed(time.Now().UnixNano())
			tmp := ""
			for i := 0; i < 4; i++ {
				tmp += fmt.Sprintf("%d", rand.Intn(9))
			}
			if len(c.Nick) > 12 {
				c.NewNick(c.Nick[:12] + tmp)
			} else {
				c.NewNick(c.Nick + tmp)
			}
			ghostErr := fmt.Errorf("ghosting of %s timed out", c.Nick)
			err := c.WaitFor(func(m *Message) bool {
				return m.Command == NOTICE &&
					strings.Contains(m.Content, "has been ghosted")
			},
				func() {
					c.Log.Println("nick taken, GHOSTING " + c.Nick)
					c.Msg("NickServ", "GHOST "+c.Nick+" "+c.Password)
				},
				time.Second*30,
				ghostErr,
			)
			if err == ghostErr {
				c.Log.Println(err)
				return
			} else if err != nil {
				return
			}
			identifyErr := fmt.Errorf("nickServ identify for %s timed out", c.Nick)
			err = c.WaitFor(func(m *Message) bool {
				return m.Command == NOTICE &&
					strings.Contains(m.Content, "You are now identified")
			},
				func() {
					c.NewNick(c.Nick)
					c.Msg("NickServ", "identify "+c.Nick+" "+c.Password)
				},
				time.Second*30,
				identifyErr,
			)
			if err == identifyErr {
				c.Log.Println(err)
			}
			return
		}
		c.Log.Printf("nick %s taken, changing nick", c.Nick)
		c.Nick = changeNick(c.Nick)
		c.NewNick(c.Nick)
	})
}

func pingpong(c chan bool) {
	select {
	case c <- true:
	default:
		return
	}
}

//HandlePingPong replies to and sends pings
func (c *Connection) HandlePingPong() {
	c.AddCallback(PING, func(msg *Message) {
		c.Log.Println("got ping sending pong")
		c.Pong()
	})
	pp := make(chan bool, 1)
	c.AddCallback(ANYMESSAGE, func(msg *Message) {
		pingpong(pp)
	})
	pingTick := time.NewTicker(time.Minute * 1)
	go func(tick *time.Ticker) {
		for range tick.C {
			select {
			case <-pp:
				c.Ping()
			default:
				c.Log.Println("got no pong")
			}
		}
	}(pingTick)
}

//HandleJoin joins channels on welcome
func (c *Connection) HandleJoin(chans []string) {
	c.AddCallback(WELCOME, func(msg *Message) {
		if c.Password != "" {
			idConfirmErr := fmt.Errorf("identification confirmation for %s timed out", c.Nick)
			err := c.WaitFor(func(m *Message) bool {
				return m.Command == NOTICE && strings.Contains(m.Content, "You are now identified for")
			},
				func() {},
				time.Second*30,
				idConfirmErr,
			)
			if err == idConfirmErr {
				c.Log.Println(err)
				c.Log.Println("trying to join the channels anyway")
			} else if err != nil {
				return
			}
		}
		c.Log.Println("joining channels")
		c.Join(chans)
	})
}

func (c *Connection) getPrefix() {
	c.AddTrigger(Trigger{
		Condition: func(m *Message) bool {
			return m.Command == JOIN && m.Name == c.Nick
		},
		Response: func(m *Message) {
			c.prefixlenSet <- []string{m.Name, m.User, m.Host}
		},
	})
}

// Start the bot
func (c *Connection) Start() {
	c.Wait()
	if c.IsConnected() || c.DebugFakeConn {
		return
	}
	err := dial(c)
	if err != nil {
		c.Errchan <- err
		return
	}
	c.Send = make(chan string)
	c.connectedSet <- true
	c.messenger = messenger.New()
	err = identify(c)
	if err != nil {
		c.Disconnect()
		c.Errchan <- err
		return
	}
	c.Add(2)
	go readLoop(c)
	go writeLoop(c)
}

func dial(c *Connection) (err error) {
	if c.TLS {
		tls, err := tls.Dial("tcp", c.Server, &tls.Config{})
		if err != nil {
			return err
		}
		c.conn = irc.NewConn(tls)
	} else {
		var err error
		c.conn, err = irc.Dial(c.Server)
		if err != nil {
			return err
		}
	}
	return nil
}

func identify(c *Connection) (err error) {
	if c.Password != "" {
		out := "PASS " + c.Password
		c.Debug.Printf("→ %s", out)
		_, err := fmt.Fprintf(c.conn, "%s%s", out, "\r\n")
		if err != nil {
			return err
		}
	}
	if c.RealN == "" {
		c.RealN = c.User
	}
	out := "USER " + c.User + " +iw * :" + c.RealN
	c.Debug.Printf("→ %s", out)
	_, err = fmt.Fprintf(c.conn, "%s%s", out, "\r\n")
	if err != nil {
		return err
	}
	out = irc.NICK + " " + c.Nick
	c.Debug.Printf("→ %s", out)
	_, err = fmt.Fprintf(c.conn, "%s%s", out, "\r\n")
	if err != nil {
		return err
	}
	return nil
}

func readLoop(c *Connection) {
	defer c.Done()
	for {
		if !c.IsConnected() {
			return
		}
		raw, err := c.conn.Decode()
		if err != nil {
			c.Disconnect()
			c.Errchan <- err
			return
		}
		c.Debug.Printf("← %s", raw)
		msg := ParseMessage(raw)
		go c.messenger.Broadcast(msg)
		go c.RunCallbacks(msg)
		go c.RunTriggers(msg)
	}
}

func writeLoop(c *Connection) {
	defer c.Done()
	for {
		if !c.IsConnected() {
			return
		}
		v, ok := <-c.Send
		if !ok {
			return
		}
		c.Debug.Printf("→ %s", v)
		_, err := fmt.Fprintf(c.conn, "%s%s", v, "\r\n")
		if err != nil {
			c.Disconnect()
			c.Errchan <- err
			return
		}
		time.Sleep(c.Throttle)
	}
}
