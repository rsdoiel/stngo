#
# A Simple Makefile for building stn shared library
# for Julia integration.
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
	EXT = .dynlib
endif
ifeq ($(OS), Windows)
	EXT = .dll
endif

build:
	go build -buildmode=c-shared -o $(LIB_NAME)$(EXT) ../$(LIB_NAME)/$(LIB_NAME).go

clean:
	if [ -f "$(LIB_NAME)$(EXT)" ]; then rm "$(LIB_NAME)$(EXT)"; fi
	if [ -f "$(LIB_NAME).h" ]; then rm "$(LIB_NAME).h"; fi


test: clean build
	julia stn_test.jl


