#
# A Simple Makefile for building stn shared library
# for Python 3 integration.
#

PROJECT = stn

LIB_NAME = libstn

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
	go build -buildmode=c-shared -o $(PROJECT)/$(LIB_NAME)$(EXT) ../$(LIB_NAME)/$(LIB_NAME).go

clean:
	if [ -f "$(PROJECT)/$(LIB_NAME)$(EXT)" ]; then rm "$(PROJECT)/$(LIB_NAME)$(EXT)"; fi
	if [ -f "$(PROJECT)/$(LIB_NAME).h" ]; then rm "$(PROJECT)/$(LIB_NAME).h"; fi
	if [ -d "stn/__pycache__" ]; then rm -fR stn/__pycache__; fi


test: clean build
	python3 stn_test.py


