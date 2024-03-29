
Installation
============

*stngo* is a set of command line programs run from a shell like Bash. You can find compiled
version in the [releases](https://github.com/rsdoiel/stngo/releases/latest) 

## Quick install with curl

The following curl command can be used to run the installer on most
POSIX systems. Programs are installed into `$HOME/bin`. `$HOME/bin` will
need to be in your path. From a shell (or terminal session) run the
following.

~~~
curl https://rsdoiel.github.io/stngo/installer.sh | sh
~~~

## Compiled version

This is generalized instructions for a release. 

Compiled versions are available for Mac OS X (Intel and M1 processors, macOS-x86_64, macOS-arm64), 
Linux (Intel process, Linux-x86_64), Windows (Intel processor, Windows-x86_64), 
Rapsberry Pi (arm7 processor, RaspberryPiOS-arm7) and Pine64 (arm64 processor, Linux-aarch64)


VERSION_NUMBER is a [symantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on Github and click latest release

>    https://github.com/rsdoiel/stngo/releases/latest


| Platform    | Zip Filename                           |
|-------------|----------------------------------------|
| Windows     | stn-VERSION_NUMBER-Windows-x86_64.zip |
| Mac OS X    | stn-VERSION_NUMBER-macOS-x86_64.zip  |
| Mac OS X    | stn-VERSION_NUMBER-macos-arm64.zip  |
| Linux/Intel | stn-VERSION_NUMBER-Linux-x86_64.zip   |
| Raspbery Pi | stn-VERSION_NUMBER-RaspberryPiOS-arm7.zip |
| Pine64      | stn-VERSION_NUMBER-Linux-aarch64.zip   |


## The basic recipe

+ Find the Zip file listed matching the architecture you're running and download it
    + (e.g. if you're on a Windows 10 laptop/Surface with a amd64 style CPU you'd choose the Zip file with "windows-amd64" in the name).
+ Download the zip file and unzip the file.  
+ Copy the contents of the folder named "bin" to a folder that is in your path 
    + (e.g. "$HOME/bin" is common).
+ Adjust your PATH if needed
    + (e.g. export PATH="$HOME/bin:$PATH")
+ Test


### macOS

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

```shell
    cd Downloads/
    unzip stngo-*-macOS-x86_64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    stnparse -version
```

### Windows

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

```shell
    cd Downloads/
    unzip stngo-*-Windows-x86_64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    stnparse -version
```


### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip stngo-*-Linux-x86_64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    stnparse -version
```


### Raspberry Pi

Released version is for a Raspberry Pi 2 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip stngo-*-RaspberryPiOS-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    stnparse -version
```


Compiling from source
---------------------

_stngo_ is "go gettable".  Use the "go get" command to download the dependant packages
as well as _stngo_'s source code.

```shell
    go get -u github.com/rsdoiel/stngo/...
```

Or clone the repstory and then compile

```shell
    cd
    git clone https://github.com/rsdoiel/stngo src/github.com/rsdoiel/stngo
    cd src/github.com/rsdoiel/stngo
    make
    make test
    make install
```

