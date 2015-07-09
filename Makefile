#
# Simple Timesheet Notation
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#

all: shorthand reldate stnparse stnfilter stnreport

shorthand: cmd/shorthand/shorthand.go stn/shorthand/shorthand.go
	go build cmd/shorthand/shorthand.go

reldate: cmd/reldate/reldate.go
	go build cmd/reldate/reldate.go

stnparse: cmd/stnparse/stnparse.go stn/stn.go
	go build cmd/stnparse/stnparse.go

stnfilter: cmd/stnfilter/stnfilter.go stn/stn.go
	go build cmd/stnfilter/stnfilter.go

stnreport: cmd/stnreport/stnreport.go stn/report/report.go
	go build cmd/stnreport/stnreport.go

lint:
	gofmt -w stn/stn.go && golint stn/stn.go
	gofmt -w stn/stn_test.go && golint stn/stn_test.go
	gofmt -w stn/report/report.go && golint stn/report/report.go
	gofmt -w stn/report/report_test.go && golint stn/report/report_test.go
	gofmt -w cmd/shorthand/shorthand.go && golint cmd/shorthand/shorthand.go
	gofmt -w stn/shorthand/shorthand_test.go && golint stn/shorthand/shorthand_test.go
	gofmt -w cmd/reldate/reldate.go && golint cmd/reldate/reldate.go
	gofmt -w cmd/stnparse/stnparse.go && golint cmd/stnparse/stnparse.go
	gofmt -w cmd/stnfilter/stnfilter.go && golint cmd/stnfilter/stnfilter.go
	gofmt -w cmd/stnreport/stnreport.go && golint cmd/stnreport/stnreport.go
	gofmt -w ok/ok.go && golint ok/ok.go
	gofmt -w ok/ok_test.go && golint ok/ok_test.go

test:
	cd stn && go test
	cd stn/shorthand && go test
	cd stn/report && go test

# ok test throws false Fail so is skipped
#	cd ok && go test

clean: shorthand reldate stnparse stnfilter stnreport
	if [ -f reldate ]; then rm reldate; fi
	if [ -f shorthand ]; then rm shorthand; fi
	if [ -f stnparse ]; then rm stnparse; fi
	if [ -f stnfilter ]; then rm stnfilter; fi
	if [ -f stnreport ]; then rm stnreport; fi

install: shorthand reldate stnparse stnfilter stnreport
	go install cmd/reldate/reldate.go
	go install cmd/shorthand/shorthand.go
	go install cmd/stnparse/stnparse.go
	go install cmd/stnfilter/stnfilter.go
	go install cmd/stnreport/stnreport.go

uninstall:
	if [ -f $(GOBIN)/reldate ]; then /bin/rm $(GOBIN)/reldate; fi
	if [ -f $(GOBIN)/shorthand ]; then /bin/rm $(GOBIN)/shorthand; fi
	if [ -f $(GOBIN)/stnparse ]; then /bin/rm $(GOBIN)/stnparse; fi
	if [ -f $(GOBIN)/stnfilter ]; then /bin/rm $(GOBIN)/stnfilter; fi
	if [ -f $(GOBIN)/stnreport ]; then /bin/rm $(GOBIN)/stnreport; fi
