#!/bin/bash
#
GO=$(which go)
if [ "$GO" = "" ]; then
    echo "Must install Golang first"
    echo "See http://golang.org for instructions"
    exit 1
fi

# Install dependent libraries
# Add ok test package
go get github.com/rsdoiel/ok
# Add shorthand package
go get github.com/rsdoiel/shorthand

make
make test
