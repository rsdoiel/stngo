#
# Simple Timesheet Notation
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#
PROJECT = stn

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build:
	go build -o bin/stnparse cmds/stnparse/stnparse.go
	go build -o bin/stnfilter cmds/stnfilter/stnfilter.go
	go build -o bin/stnreport cmds/stnreport/stnreport.go

lint:
	gofmt -w stn.go && golint stn.go
	gofmt -w stn_test.go && golint stn_test.go
	gofmt -w report/report.go && golint report/report.go
	gofmt -w report/report_test.go && golint report/report_test.go
	gofmt -w cmds/stnparse/stnparse.go && golint cmds/stnparse/stnparse.go
	gofmt -w cmds/stnfilter/stnfilter.go && golint cmds/stnfilter/stnfilter.go
	gofmt -w cmds/stnreport/stnreport.go && golint cmds/stnreport/stnreport.go
	goimports -w stn.go && golint stn.go
	goimports -w stn_test.go && golint stn_test.go
	goimports -w report/report.go && golint report/report.go
	goimports -w report/report_test.go && golint report/report_test.go
	goimports -w cmds/stnparse/stnparse.go && golint cmds/stnparse/stnparse.go
	goimports -w cmds/stnfilter/stnfilter.go && golint cmds/stnfilter/stnfilter.go
	goimports -w cmds/stnreport/stnreport.go && golint cmds/stnreport/stnreport.go

test:
	go test
	cd report && go test

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/stnparse/stnparse.go
	env GOBIN=$(HOME)/bin go install cmds/stnfilter/stnfilter.go
	env GOBIN=$(HOME)/bin go install cmds/stnreport/stnreport.go

uninstall:
	if [ -f $(GOBIN)/stnparse ]; then /bin/rm $(GOBIN)/stnparse; fi
	if [ -f $(GOBIN)/stnfilter ]; then /bin/rm $(GOBIN)/stnfilter; fi
	if [ -f $(GOBIN)/stnreport ]; then /bin/rm $(GOBIN)/stnreport; fi

website: build
	./bin/stnparse --version
	./mk-website.bash

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/stnparse cmds/stnparse/stnparse.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/stnreport cmds/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/stnparse.exe cmds/stnparse/stnparse.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/stnfilter.exe cmds/stnfilter/stnfilter.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/stnreport.exe cmds/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/stnparse cmds/stnparse/stnparse.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/stnreport cmds/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/stnparse cmds/stnparse/stnparse.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/stnreport cmds/stnreport/stnreport.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/stnparse cmds/stnparse/stnparse.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/stnreport cmds/stnreport/stnreport.go
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

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/linux-arm64

publish:
	./mk-website.bash
	./publish.bash

