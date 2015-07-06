#
# Simple Makefile
#

all: reldate shorthand stnparse stnfilter stnreport

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

test:
	cd stn && go test
	cd stn/shorthand && go test


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
