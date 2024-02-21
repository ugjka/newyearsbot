module github.com/ugjka/newyearsbot

require (
	github.com/badoux/checkmail v1.2.4
	github.com/fatih/color v1.16.0
	github.com/hako/durafmt v0.0.0-20210608085754-5c1018a4e16b
	github.com/ugjka/kittybot v0.0.62
	gopkg.in/inconshreveable/log15.v2 v2.16.0
	gopkg.in/ugjka/go-tz.v2 v2.2.0
	gopkg.in/yaml.v3 v3.0.1
	mvdan.cc/xurls/v2 v2.5.0
)

require (
	github.com/ftrvxmtrx/fd v0.0.0-20150925145434-c6d800382fff // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ugjka/ircmsg v0.0.3 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/term v0.17.0 // indirect
)

//replace github.com/ugjka/kittybot => ../kittybot

go 1.18
