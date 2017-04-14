prefix=/usr/local
PWD := $(shell pwd)
GOPATH :=$(PWD)/build
appname = newyearsbot

all: cli
	GOPATH=$(GOPATH) go get github.com/ugjka/$(appname)/gui

cli:
	GOPATH=$(GOPATH) go get github.com/ugjka/$(appname)
install:
	install -Dm755 $(GOPATH)/bin/$(appname) $(prefix)/bin/$(appname)
	install -Dm644 LICENSE "$(prefix)/share/licenses/$(appname)/LICENSE"
	if [ -a $(GOPATH)/bin/gui ]; then \
		install -Dm755 $(GOPATH)/bin/gui $(prefix)/bin/$(appname)-gui; \
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