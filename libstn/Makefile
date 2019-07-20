#
# A Simple Makefile for building stn shared library
# for Python 3 integration.
#

PROJECT = stn

VERSION = $(shell grep -m1 'Version = ' ../$(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = .so

ifeq ($(OS), Linux)
	EXT = .so
endif
ifeq ($(OS), Darwin)
	EXT = .dylib
endif
ifeq ($(OS), Windows)
	EXT = .dll
endif

build:
	go build -buildmode=c-shared -o libstn$(EXT)

clean:
	if [ -f "libstn$(EXT)" ]; then rm libstn$(EXT); fi
	if [ -f "libstn.h" ]; then rm libstn.h; fi


