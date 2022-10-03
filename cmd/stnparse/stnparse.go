//
// stnparse.go - Simple Timesheet Notation parser.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2021 all rights reserved.
// Released under the BSD 2-Clause license.
// See: http://opensource.org/licenses/BSD-2-Clause
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	// My packages
	stn "github.com/rsdoiel/stngo"
)

var (

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// App Options
	asJSON bool
)

func main() {
	appName := path.Base(os.Args[0])

	// Configuration and command line interation

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")

	// App Options
	flag.BoolVar(&asJSON, "json", false, "output JSON format")

	flag.Parse()
	args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if showHelp {
		fmt.Fprintln(out, fmtText(helpText, appName, stn.Version))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(out, fmtText(licenseText, appName, stn.Version))
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", appName, stn.Version)
		os.Exit(0)
	}

	if inputFName == "" && len(args) > 0 {
		inputFName = args[0]
	}
	if outputFName == "" && len(args) > 1 {
		outputFName = args[1]
	}
	if inputFName != "" && inputFName != "-" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer in.Close()
	}
	if outputFName != "" && outputFName != "-" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer out.Close()
	}

	activeDate := time.Now().Format(stn.DateFmt)

	reader := bufio.NewReader(in)

	entryCnt := 0
	lineNo := 1
	if asJSON == true {
		fmt.Fprint(out, "[")
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if stn.IsDateLine(line) == true {
			activeDate = stn.ParseDateLine(line)
		} else if stn.IsEntry(line) {
			entry, perr := stn.ParseEntry(activeDate, line)
			if perr != nil {
				fmt.Fprintf(eout, "line %d: %v\n", lineNo, perr)
			}
			if asJSON == true {
				if entryCnt > 0 {
					fmt.Fprint(out, ",")
				}
				fmt.Fprint(out, entry.JSON())
				entryCnt++
			} else {
				fmt.Fprintln(out, entry.String())
			}
		}
		lineNo++
	}
	if asJSON == true {
		fmt.Fprint(out, "]")
	}
}
