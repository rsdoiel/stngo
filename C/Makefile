#
# A Simple Makefile for building libstn shared library
# for C integration.
#

PROJECT = stn

VERSION = $(shell grep -m1 'Version = ' ../$(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = .so
DLL =

ifeq ($(OS), Linux)
	DLL = .so
	EXT =
endif
ifeq ($(OS), Darwin)
	DLL = .dylib
	EXT =
endif
ifeq ($(OS), Windows)
	DLL = .dll
	EXIT = .exe
endif

build: lib/libstn$(DLL) main$(EXT)

lib/libstn$(DLL): lib/libstn.go
	cd lib && go build -o libstn$(DLL) -buildmode=c-shared

main$(EXT):
	gcc -Wall -o main main.c lib/libstn$(DLL)

clean:
	if [ -f "lib/libstn$(DLL)" ]; then rm lib/libstn$(DLL); fi
	if [ -f lib/libstn.h ]; then rm lib/libstn.h; fi
	if [ -f main.o ]; then rm main.o; fi
	if [ -f main$(EXT) ]; then rm main$(EXT); fi
	if [ -f stn.pyc ]; then rm stn.pyc; fi
	if [ -d __pycache__ ]; then rm -fR __pycache__; fi

test: clean build
	./main


build:
	gcc -c main.c 
	gcc -o m main.o
