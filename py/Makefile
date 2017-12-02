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
	EXT = .dynlib
endif
ifeq ($(OS), Windows)
	EXT = .dll
endif

build:
	cd lib && go build -buildmode=c-shared -o libstn$(EXT)

clean:
	if [ -f "lib/libstn$(EXT)" ]; then rm lib/libstn$(EXT); fi
	if [ -f "lib/libstn.h" ]; then rm lib/libstn.h; fi
	if [ -f "stn.pyc" ]; then rm stn.pyc; fi
	if [ -d "__pycache__" ]; then rm -fR __pycache__; fi

test: clean build
	python3 stn.py

