#
# Simple Timesheet Notation
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#
bin/stnparse:
	go build -o bin/stnparse cmds/stnparse/stnparse.go

bin/stnfilter:
	go build -o bin/stnfilter cmds/stnfilter/stnfilter.go

bin/stnreport:
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

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

build:
	go build -o bin/stnparse cmds/stnparse/stnparse.go
	go build -o bin/stnfilter cmds/stnfilter/stnfilter.go
	go build -o bin/stnreport cmds/stnreport/stnreport.go

install:
	env GOBIN=$HOME/bin go install cmds/stnparse/stnparse.go
	env GOBIN=$HOME/bin go install cmds/stnfilter/stnfilter.go
	env GOBIN=$HOME/bin go install cmds/stnreport/stnreport.go

uninstall:
	if [ -f $(GOBIN)/stnparse ]; then /bin/rm $(GOBIN)/stnparse; fi
	if [ -f $(GOBIN)/stnfilter ]; then /bin/rm $(GOBIN)/stnfilter; fi
	if [ -f $(GOBIN)/stnreport ]; then /bin/rm $(GOBIN)/stnreport; fi

website:
	./bin/stnparse --version
	shorthand build.shorthand

release:
	./mk-release.sh
