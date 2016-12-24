#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#
PROJECT=stngo

PROG_LIST="stnparse stnfilter stnreport"

VERSION=$(grep -m1 'Version = ' $PROJECT.go | cut -d\" -f 2)

RELEASE_VERSION=$PROJECT-$VERSION

echo "Preparing $RELEASE_VERSION"
for PROGNAME in $PROG_LIST; do
  echo "Cross compelling $PROGNAME"
  env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  env GOOS=windows GOARCH=amd64 go build -o dist/windows/$PROGNAME.exe cmds/$PROGNAME/$PROGNAME.go
  env GOOS=darwin	GOARCH=amd64 go build -o dist/maxosx/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
done

echo "Assembling dist/"
for FNAME in README.md LICENSE INSTALL.md; do
    cp -v $FNAME dist/
done
echo "Zipping up $RELEASE_NAME"
zip -r $REPONAME-binary-release.zip dist/*
