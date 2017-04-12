prefix=/usr/local
PWD := $(shell pwd)
GOPATH :=$(PWD)/build
appname = newyearsbot

all:
	go get github.com/ugjka/$(appname)
	go get github.com/ugjka/$(appname)/gui

install:
	install -Dm755 $(GOPATH)/bin/$(appname) $(prefix)/bin/$(appname)
	install -Dm755 $(GOPATH)/bin/gui $(prefix)/bin/$(appname)-gui
	install -Dm644 icon.png "$(prefix)/share/icons/hicolor/256x256/apps/$(appname).png"
	install -Dm644 LICENSE "$(prefix)/share/licenses/$(appname)/LICENSE"
	install -Dm644 $(appname).desktop "$(prefix)/share/applications/$(appname).desktop"

uninstall:
	rm "$(prefix)/bin/$(appname)"
	rm "$(prefix)/bin/$(appname)-gui"
	rm "$(prefix)/share/icons/hicolor/256x256/apps/$(appname).png"
	rm "$(prefix)/share/licenses/$(appname)/LICENSE"
	rm "$(prefix)/share/applications/$(appname).desktop"

clean:
	rm -rf $(GOPATH)