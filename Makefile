#
# Simple Timesheet Notation
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#
bin/build: bin/reldate bin/stnparse bin/stnfilter bin/stnreport

bin/reldate: cmd/reldate/reldate.go
	go build -o bin/reldate cmd/reldate/reldate.go

bin/stnparse: cmd/stnparse/stnparse.go stn/stn.go
	go build -o bin/stnparse cmd/stnparse/stnparse.go

bin/stnfilter: cmd/stnfilter/stnfilter.go stn/stn.go
	go build -o bin/stnfilter cmd/stnfilter/stnfilter.go

bin/stnreport: cmd/stnreport/stnreport.go report/report.go
	go build -o bin/stnreport cmd/stnreport/stnreport.go

lint:
	gofmt -w stn/stn.go && golint stn/stn.go
	gofmt -w stn/stn_test.go && golint stn/stn_test.go
	gofmt -w report/report.go && golint report/report.go
	gofmt -w report/report_test.go && golint report/report_test.go
	gofmt -w cmd/reldate/reldate.go && golint cmd/reldate/reldate.go
	gofmt -w cmd/stnparse/stnparse.go && golint cmd/stnparse/stnparse.go
	gofmt -w cmd/stnfilter/stnfilter.go && golint cmd/stnfilter/stnfilter.go
	gofmt -w cmd/stnreport/stnreport.go && golint cmd/stnreport/stnreport.go

test:
	cd stn && go test
	cd report && go test

# ok test throws false Fail so is skipped
#	cd ok && go test

clean: bin/shorthand bin/reldate bin/stnparse bin/stnfilter bin/stnreport
	if [ -f bin/reldate ]; then rm bin/reldate; fi
	if [ -f bin/stnparse ]; then rm bin/stnparse; fi
	if [ -f bin/stnfilter ]; then rm bin/stnfilter; fi
	if [ -f bin/stnreport ]; then rm bin/stnreport; fi

build:
	go build -o bin/reldate cmd/reldate/reldate.go
	go build -o bin/stnparse cmd/stnparse/stnparse.go
	go build -o bin/stnfilter cmd/stnfilter/stnfilter.go
	go build -o bin/stnreport cmd/stnreport/stnreport.go

install:
	go install cmd/reldate/reldate.go
	go install cmd/stnparse/stnparse.go
	go install cmd/stnfilter/stnfilter.go
	go install cmd/stnreport/stnreport.go

uninstall:
	if [ -f $(GOBIN)/reldate ]; then /bin/rm $(GOBIN)/reldate; fi
	if [ -f $(GOBIN)/stnparse ]; then /bin/rm $(GOBIN)/stnparse; fi
	if [ -f $(GOBIN)/stnfilter ]; then /bin/rm $(GOBIN)/stnfilter; fi
	if [ -f $(GOBIN)/stnreport ]; then /bin/rm $(GOBIN)/stnreport; fi

