// New Year's Eve IRC party bot
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/badoux/checkmail"
	"github.com/fatih/color"
	"github.com/ugjka/newyearsbot/nyb"
	log "gopkg.in/inconshreveable/log15.v2"
	"gopkg.in/yaml.v3"
)

const usage = `
New Year's Eve IRC party bot
Announces new years as they happen in each timezone

CMD Options:
[mandatory]
-channels	comma separated list of channels eg. "#test, #test2"
		channel key can be specifed after ":" e.g #channelname:channelkey
-nick		irc nick
-email		nominatim email

[optional]
-password	irc server password
-saslnick	sasl username
-saslpass	sasl password
-server		irc server (default: irc.libera.chat:6697)
-prefix		command prefix (default: !)
-nossl		disable ssl for irc
-nominatim	nominatim server (default: http://nominatim.openstreetmap.org)
-nolimit	disable flood kick protection
-colors		enable irc colors
-debug		debug irc traffic
-yaml		yaml config file

`
const SET_NOMINATIM_SERVER = "https://nominatim.openstreetmap.org"
const SET_LIBERA_SERVER = "irc.libera.chat:6697"
const SET_PREFIX = "!"

func main() {
	var channels nyb.Channels

	// Mandatory
	flag.Var(&channels, "channels", "comma separated list of channels")
	nick := flag.String("nick", "", "irc nick")
	email := flag.String("email", "", "nominatim email")
	// Optional
	password := flag.String("password", "", "irc server password")
	saslNick := flag.String("saslnick", "", "sasl username")
	saslPass := flag.String("saslpass", "", "sasl password")
	server := flag.String("server", SET_LIBERA_SERVER, "irc server")
	prefix := flag.String("prefix", SET_PREFIX, "command prefix")
	nossl := flag.Bool("nossl", false, "disable ssl for irc")
	nominatim := flag.String("nominatim", SET_NOMINATIM_SERVER, "nominatim server")
	nolimit := flag.Bool("nolimit", false, "disable limit bot replies.")
	colors := flag.Bool("colors", false, "enable irc colors")
	debug := flag.Bool("debug", false, "debug irc traffic")
	bind := flag.String("bind", "", "bind to host address")
	configYAML := flag.String("yaml", "", "use yaml settings file")

	green := color.New(color.FgGreen)
	flag.Usage = func() {
		green.Fprint(os.Stderr, usage)
	}
	flag.Parse()

	c := config{
		{
			Nick:      *nick,
			Channels:  channels,
			Server:    *server,
			NoSSL:     *nossl,
			Password:  *password,
			SaslNick:  *saslNick,
			SaslPass:  *saslPass,
			Prefix:    *prefix,
			Email:     *email,
			Nominatim: *nominatim,
			NoLimit:   *nolimit,
			Colors:    *colors,
			Debug:     *debug,
			Bind:      *bind,
		},
	}

	red := color.New(color.FgHiRed)

	if *configYAML != "" {
		data, err := os.ReadFile(*configYAML)
		if err != nil {
			red.Fprintln(os.Stderr, "yaml file: ", err)
			os.Exit(1)
		}
		c = config{}
		err = yaml.Unmarshal(data, &c)
		if err != nil {
			red.Fprintln(os.Stderr, "yaml file: ", err)
			os.Exit(1)
		}
		for i := range c {
			if c[i].Nominatim == "" {
				c[i].Nominatim = SET_NOMINATIM_SERVER
			}
			if c[i].Server == "" {
				c[i].Server = SET_LIBERA_SERVER
			}
			if c[i].Prefix == "" {
				c[i].Prefix = SET_PREFIX
			}
		}
	}

	err := check(c)
	if err != nil {
		red.Fprintln(os.Stderr, err)
		if *configYAML == "" {
			flag.Usage()
		}
		os.Exit(1)
	}
	var bots []*nyb.Settings
	for _, c := range c {
		var err error
		var customTLSDial nyb.TLSDialFunc
		var customDial nyb.DialFunc
		if c.NoSSL && c.Bind != "" {
			customDial, err = func() (nyb.DialFunc, error) {
				if c.Bind == "" {
					return nil, nil
				}
				localAddr, err := net.ResolveIPAddr("ip", c.Bind)
				if err != nil {
					return nil, err
				}

				localTCPAddr := net.TCPAddr{
					IP: localAddr.IP,
				}
				remoteAddr, err := net.ResolveTCPAddr("tcp", c.Server)
				if err != nil {
					return nil, err
				}
				dialer := func(network string, addr string) (net.Conn, error) {
					return net.DialTCP(network, &localTCPAddr, remoteAddr)
				}
				return dialer, nil
			}()
			if err != nil {
				fmt.Fprintln(os.Stderr, "BINDHOST:", err)
				os.Exit(1)
			}
		}
		if !c.NoSSL && c.Bind != "" {
			customTLSDial, err = func() (nyb.TLSDialFunc, error) {
				if c.Bind == "" {
					return nil, nil
				}

				localAddr, err := net.ResolveIPAddr("ip", c.Bind)
				if err != nil {
					return nil, err
				}

				localTCPAddr := net.TCPAddr{
					IP: localAddr.IP,
				}

				dialer := &net.Dialer{
					LocalAddr: &localTCPAddr,
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}

				tlsdialer := func(network string, addr string, tlsConf *tls.Config) (*tls.Conn, error) {
					return tls.DialWithDialer(dialer, network, addr, &tls.Config{})
				}
				return tlsdialer, nil
			}()
			if err != nil {
				fmt.Fprintln(os.Stderr, "BINDHOST:", err)
				os.Exit(1)
			}
		}
		bots = append(bots,
			nyb.New(
				&nyb.Settings{
					Nick:      c.Nick,
					Channels:  c.Channels,
					Server:    c.Server,
					SSL:       !c.NoSSL,
					Password:  c.Password,
					SaslNick:  c.SaslNick,
					SaslPass:  c.SaslPass,
					Prefix:    c.Prefix,
					Email:     c.Email,
					Nominatim: c.Nominatim,
					Limit:     !c.NoLimit,
					Colors:    c.Colors,
					Dial:      customDial,
					TLSDial:   customTLSDial,
				},
			),
		)
	}

	for i, bot := range bots {
		if c[i].Debug {
			bot.LogLvl(log.LvlDebug)
		} else {
			bot.LogLvl(log.LvlInfo)
		}
		go bot.Start()
	}
	select {}
}

type config []struct {
	Nick      string
	Channels  []string
	Server    string
	NoSSL     bool
	Password  string
	SaslNick  string
	SaslPass  string
	Prefix    string
	Email     string
	Nominatim string
	NoLimit   bool
	Colors    bool
	Debug     bool
	Bind      string
}

func check(c config) error {
	if len(c) == 0 {
		return fmt.Errorf("empty or misconfigured yaml")
	}

	for _, c := range c {

		// Check mandatory inputs
		if len(c.Channels) == 0 {
			return fmt.Errorf("error: no channels defined")
		}
		channelReg := regexp.MustCompile(`^([#&][^\x07\x2C\s]{0,200})$`)
		for _, ch := range c.Channels {
			if !channelReg.MatchString(ch) {
				return fmt.Errorf("error: invalid channel name: %s", ch)
			}
		}
		if c.Nick == "" {
			return fmt.Errorf("error: no nick defined")
		}
		if len(c.Nick) > 16 {
			return fmt.Errorf("error: nick too long")
		}
		if c.Email == "" {
			return fmt.Errorf("error: no email provided")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return fmt.Errorf("error: invalid email address")
		}
		// Check optional inputs
		if c.Server == "" {
			return fmt.Errorf("error: no irc server defined")
		}
		serverReg := regexp.MustCompile(`^\S+:\d+$`)
		if !serverReg.MatchString(c.Server) {
			return fmt.Errorf("error: invalid irc server address")
		}
		if c.Prefix == "" {
			return fmt.Errorf("error: no command prefix defined")
		}
		prefixReg := regexp.MustCompile(`^\W+$`)
		if !prefixReg.MatchString(c.Prefix) {
			return fmt.Errorf("error: prefix must be non-alphanumeric")
		}
		if c.Nominatim == "" {
			return fmt.Errorf("error: no nominatim server provided")
		}
	}
	return nil
}
