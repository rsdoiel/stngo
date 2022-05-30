#
# Simple Timesheet Notation
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#
PROJECT = stn

VERSION = $(shell grep -m1 'Version = `' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

#PREFIX = /usr/local
PREFIX = $(HOME)


build:
	go build -o bin/stnparse cmd/stnparse/stnparse.go
	go build -o bin/stnfilter cmd/stnfilter/stnfilter.go
	go build -o bin/stnreport cmd/stnreport/stnreport.go

lint:
	gofmt -w stn.go && golint stn.go
	gofmt -w stn_test.go && golint stn_test.go
	gofmt -w report/report.go && golint report/report.go
	gofmt -w report/report_test.go && golint report/report_test.go
	gofmt -w cmd/stnparse/stnparse.go && golint cmd/stnparse/stnparse.go
	gofmt -w cmd/stnfilter/stnfilter.go && golint cmd/stnfilter/stnfilter.go
	gofmt -w cmd/stnreport/stnreport.go && golint cmd/stnreport/stnreport.go
	goimports -w stn.go && golint stn.go
	goimports -w stn_test.go && golint stn_test.go
	goimports -w report/report.go && golint report/report.go
	goimports -w report/report_test.go && golint report/report_test.go
	goimports -w cmd/stnparse/stnparse.go && golint cmd/stnparse/stnparse.go
	goimports -w cmd/stnfilter/stnfilter.go && golint cmd/stnfilter/stnfilter.go
	goimports -w cmd/stnreport/stnreport.go && golint cmd/stnreport/stnreport.go

test:
	go test
	cd report && go test

man: build
	mkdir -p man/man1
	bin/stnparse -generate-manpage | nroff -Tutf8 -man > man/man1/stnparse.1
	bin/stnfilter -generate-manpage | nroff -Tutf8 -man > man/man1/stnfilter.1
	bin/stnreport -generate-manpage | nroff -Tutf8 -man > man/man1/stnreport.1

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

install:
	env GOBIN=$(PREFIX)/bin go install cmd/stnparse/stnparse.go
	env GOBIN=$(PREFIX)/bin go install cmd/stnfilter/stnfilter.go
	env GOBIN=$(PREFIX)/bin go install cmd/stnreport/stnreport.go

uninstall:
	if [ -f $(PREFIX)/bin/stnparse ]; then /bin/rm $(PREFIX)/bin/stnparse; fi
	if [ -f $(PREFIX)/bin/stnfilter ]; then /bin/rm $(PREFIX)/bin/stnfilter; fi
	if [ -f $(PREFIX)/bin/stnreport ]; then /bin/rm $(PREFIX)/bin/stnreport; fi

website: build
	./bin/stnparse --version
	./mk_website.py

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/stnparse cmd/stnparse/stnparse.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/stnfilter cmd/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/stnreport cmd/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/stnparse.exe cmd/stnparse/stnparse.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/stnfilter.exe cmd/stnfilter/stnfilter.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/stnreport.exe cmd/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macos-amd64:
	mkdir -p dist/bin
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/stnparse cmd/stnparse/stnparse.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/stnfilter cmd/stnfilter/stnfilter.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/stnreport cmd/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macos-arm64:
	mkdir -p dist/bin
	env GOOS=darwin	GOARCH=arm64 go build -o dist/bin/stnparse cmd/stnparse/stnparse.go
	env GOOS=darwin	GOARCH=arm64 go build -o dist/bin/stnfilter cmd/stnfilter/stnfilter.go
	env GOOS=darwin	GOARCH=arm64 go build -o dist/bin/stnreport cmd/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-arm64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/stnparse cmd/stnparse/stnparse.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/stnfilter cmd/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/stnreport cmd/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/stnparse cmd/stnparse/stnparse.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/stnfilter cmd/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/stnreport cmd/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v stnparse.md dist/
	cp -v stnfilter.md dist/
	cp -v stnreport.md dist/

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macos-amd64 dist/macos-arm64 dist/raspbian-arm7 dist/linux-arm64

publish:
	./mk_website.py
	./publish.bash

