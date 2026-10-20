Installation for development of **stngo**
===========================================

**stngo** A set of command line tools and Go package for working with Standard Timesheet Notation.

Quick install with curl or irm
------------------------------

There is an experimental installer.sh script that can be run with the following command to install latest table release. This may work for macOS, Linux and if you’re using Windows with the Unix subsystem. This would be run from your shell (e.g. Terminal on macOS).

~~~shell
curl https://rsdoiel.github.io/stngo/installer.sh | sh
~~~

This will install the programs included in stngo in your `$HOME/bin` directory.

If you are running Windows 10 or 11 use the Powershell command below.

~~~ps1
irm https://rsdoiel.github.io/stngo/installer.ps1 | iex
~~~

### If your are running macOS or Windows

You may get security warnings if you are using macOS or Windows. See the notes for the specific operating system you're using to fix issues.

- [INSTALL_NOTES_macOS.md](INSTALL_NOTES_macOS.md)
- [INSTALL_NOTES_Windows.md](INSTALL_NOTES_Windows.md)

Installing from source
----------------------

### Required software

- Go >= 1.26.4

Optional, for regenerating `version.go`, the man pages and the website:

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3

### Steps

1. git clone https://github.com/rsdoiel/stngo
2. Change directory into the `stngo` directory
3. Make to build, test and install

~~~shell
git clone https://github.com/rsdoiel/stngo
cd stngo
make
make test
make install
~~~

