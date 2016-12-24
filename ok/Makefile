#
# Shorthand a text label expander.
#
# @author R. S. Doiel, <rsdoiel@gmail.com>
# copyright (c) 2015 all rights reserved.
# Released under the BSD 2-Clause license
# See: http://opensource.org/licenses/BSD-2-Clause
#
build:
	go build -o bin/shorthand cmds/shorthand/shorthand.go

lint:
	gofmt -w ok.go && golint ok.go
	gofmt -w ok_test.go && golint ok_test.go

test:
	go test

install:
	env GOBIN=$(HOME)/bin go install cmds/shorthand/shorthand.go

release:
	./mk-release.sh

