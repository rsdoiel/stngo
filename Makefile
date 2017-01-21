#
# Simple Timesheet Notation
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#
PROJECT = stngo

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
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then /bin/rm $(PROJECT)-$(VERSION)-release.zip; fi

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

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

dist/linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/stnparse cmds/stnparse/stnparse.go
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/stnreport cmds/stnreport/stnreport.go

dist/windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o dist/windows/stnparse.exe cmds/stnparse/stnparse.go
	env GOOS=windows GOARCH=amd64 go build -o dist/windows/stnfilter.exe cmds/stnfilter/stnfilter.go
	env GOOS=windows GOARCH=amd64 go build -o dist/windows/stnreport.exe cmds/stnreport/stnreport.go

dist/macosx-amd64:
	env GOOS=darwin	GOARCH=amd64 go build -o dist/maxosx/stnparse cmds/stnparse/stnparse.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/maxosx/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/maxosx/stnreport cmds/stnreport/stnreport.go

dist/raspbian-arm7:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/stnparse cmds/stnparse/stnparse.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/stnreport cmds/stnreport/stnreport.go

dist/raspbian-arm6:
	env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/stnparse cmds/stnparse/stnparse.go
	env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/stnfilter cmds/stnfilter/stnfilter.go
	env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/stnreport cmds/stnreport/stnreport.go

release: dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/raspbian-arm6
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/*

publish:
	./mk-website.bash
	./publish.bash

