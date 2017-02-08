
# Installation

*stngo* is a set of command line programs run from a shell like Bash. You can find compiled
version in the [releases](https://github.com/rsdoiel/stngo/releases/latest) 
in the Github repository in a zip file named *stn-VERSION_NO-release.zip*. VERSION_NO reflects
the version number of the release (e.g. v0.0.5). Inside
the zip file look for the directory that matches your computer and copy that someplace
defined in your path (e.g. $HOME/bin). 

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

## Mac OS X

1. Go to [github.com/rsdoiel/stngo/releases/latest](https://github.com/rsdoiel/stngo/releases/latest)
2. Click on the green "stn-VERSION_NO-release.zip" link and download
3. Open a finder window and find the downloaded file and unzip it (e.g. stn-VERSION_NO-release.zip)
4. Look in the unziped folder and find dist/macosx-amd64/
5. Drag (or copy) the *stnparse*, *stnfilter* and *stnreport* to a "bin" directory in your path
6. Open and "Terminal" and run `stnparse -h`

## Windows

1. Go to [github.com/rsdoiel/stngo/releases/latest](https://github.com/rsdoiel/stngo/releases/latest)
2. Click on the green "stn-VERSION_NO-release.zip" link and download
3. Open the file manager find the downloaded file and unzip it (e.g. stn-VERSION_NO-release.zip)
4. Look in the unziped folder and find dist/windows-amd64/
5. Drag (or copy) the *stnparse.exe*, *stnfilter.exe*, and *stnreport.exe* to a "bin" directory in your path
6. Open Bash and and run `stnparse -h`

## Linux

1. Go to [github.com/rsdoiel/stngo/releases/latest](https://github.com/rsdoiel/stngo/releases/latest)
2. Click on the green "stn-VERSION_NO-release.zip" link and download
3. Find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/stn-VERSION_NO-release.zip)
4. In the unziped directory and find for dist/linux-amd64/
5. Copy the *stnparse*, *stnfilter*, and *stnreport* to a "bin" directory (e.g. cp ~/Downloads/stn-VERSION_NO-release/dist/linux-amd64/stngo ~/bin/)
6. From the shell prompt run `stnparse -h`

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/rsdoiel/stngo/releases/latest](https://github.com/rsdoiel/stngo/releases/latest)
2. Click on the green "stn-VERSION_NO-release.zip" link and download
3. Find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/stn-VERSION_NO-release.zip)
4. In the unziped directory and find for dist/raspberrypi-arm7
5. Copy the *stnparse*, *stnfilter*, and *stnreport* to a "bin" directory (e.g. cp ~/Downloads/stn-VERSION_NO-release/dist/raspberrypi-arm7/stngo ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
6. From the shell prompt run `stnparse -h`

