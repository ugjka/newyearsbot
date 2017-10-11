prefix=/usr/local
PWD := $(shell pwd)
GOPATH :=$(PWD)/deps
appname = newyearsbot

all: cli
	GOPATH=$(GOPATH) go get -d github.com/ugjka/$(appname)/gui
	GOPATH=$(GOPATH) go build -v -ldflags="-X main.icon=$(prefix)/share/icons/hicolor/256x256/apps/$(appname).png" -o ./newyearsbot-gui gui/*

cli:
	GOPATH=$(GOPATH) go get -d github.com/ugjka/$(appname)
	GOPATH=$(GOPATH) go build -v
install:
	install -Dm755 $(appname) $(prefix)/bin/$(appname)
	install -Dm644 LICENSE "$(prefix)/share/licenses/$(appname)/LICENSE"
	if [ -a $(appname)-gui ]; then \
		install -Dm755 $(appname)-gui $(prefix)/bin/$(appname)-gui; \
		install -Dm644 icon.png "$(prefix)/share/icons/hicolor/256x256/apps/$(appname).png"; \
		install -Dm644 $(appname).desktop "$(prefix)/share/applications/$(appname).desktop"; \
	fi

uninstall:
	rm "$(prefix)/bin/$(appname)"
	rm "$(prefix)/share/licenses/$(appname)/LICENSE"
	if [ -a $(prefix)/bin/$(appname)-gui ]; then \
		rm "$(prefix)/bin/$(appname)-gui"; \
		rm "$(prefix)/share/icons/hicolor/256x256/apps/$(appname).png"; \
		rm "$(prefix)/share/applications/$(appname).desktop"; \
	fi

clean:
	rm -rf $(GOPATH)
	rm $(appname)
	if [ -a $(appname)-gui ]; then \
		rm $(appname)-gui; \
	fi