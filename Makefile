prefix=/usr/local
PWD := $(shell pwd)
GOPATH=/usr/local/go/bin/go
appname = newyearsbot

all:
	GOPATH=$(GOPATH) go build -v
install:
	install -Dm755 $(appname) $(prefix)/bin/$(appname)
	install -Dm644 LICENSE "$(prefix)/share/licenses/$(appname)/LICENSE"

uninstall:
	rm "$(prefix)/bin/$(appname)"
	rm "$(prefix)/share/licenses/$(appname)/LICENSE"

clean:
	chmod -R 755 $(GOPATH)
	rm -rf $(GOPATH)
	rm $(appname)
