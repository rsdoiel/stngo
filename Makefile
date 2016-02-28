#
# Simple Timesheet Notation
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#
bin/build: bin/stnparse bin/stnfilter bin/stnreport

bin/stnparse: cmds/stnparse/stnparse.go stn/stn.go
	go build -o bin/stnparse cmds/stnparse/stnparse.go

bin/stnfilter: cmds/stnfilter/stnfilter.go stn/stn.go
	go build -o bin/stnfilter cmds/stnfilter/stnfilter.go

bin/stnreport: cmds/stnreport/stnreport.go report/report.go
	go build -o bin/stnreport cmds/stnreport/stnreport.go

lint:
	gofmt -w stn/stn.go && golint stn/stn.go
	gofmt -w stn/stn_test.go && golint stn/stn_test.go
	gofmt -w report/report.go && golint report/report.go
	gofmt -w report/report_test.go && golint report/report_test.go
	gofmt -w cmds/stnparse/stnparse.go && golint cmds/stnparse/stnparse.go
	gofmt -w cmds/stnfilter/stnfilter.go && golint cmds/stnfilter/stnfilter.go
	gofmt -w cmds/stnreport/stnreport.go && golint cmds/stnreport/stnreport.go

test:
	cd stn && go test
	cd report && go test

# ok test throws false Fail so is skipped
#	cd ok && go test

clean: bin/shorthand bin/stnparse bin/stnfilter bin/stnreport
	if [ -f bin/stnparse ]; then rm bin/stnparse; fi
	if [ -f bin/stnfilter ]; then rm bin/stnfilter; fi
	if [ -f bin/stnreport ]; then rm bin/stnreport; fi

build:
	go build -o bin/stnparse cmds/stnparse/stnparse.go
	go build -o bin/stnfilter cmds/stnfilter/stnfilter.go
	go build -o bin/stnreport cmds/stnreport/stnreport.go

install:
	go install cmds/stnparse/stnparse.go
	go install cmds/stnfilter/stnfilter.go
	go install cmds/stnreport/stnreport.go

uninstall:
	if [ -f $(GOBIN)/stnparse ]; then /bin/rm $(GOBIN)/stnparse; fi
	if [ -f $(GOBIN)/stnfilter ]; then /bin/rm $(GOBIN)/stnfilter; fi
	if [ -f $(GOBIN)/stnreport ]; then /bin/rm $(GOBIN)/stnreport; fi

website: build
	./bin/stnparse --version
	shorthand build.shorthand

