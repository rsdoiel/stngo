//
// stnfilter.go - Simple Timesheet Notation filter.
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
	start  string
	end    string
	match  string
	asJSON bool
)

func main() {
	appName := path.Base(os.Args[0])

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&inputFName, "i", "", "input file name")
	flag.StringVar(&outputFName, "o", "", "output file name")

	// App Options
	flag.StringVar(&match, "match", "", "Match text annotations")
	flag.StringVar(&start, "start", "", "start of inclusive date range")
	flag.StringVar(&end, "end", "", "end of inclusive date range")
	flag.BoolVar(&asJSON, "json", false, "output JSON format")

	// Run the command line interface
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

	// On to running the app
	var (
		showLine   = true
		startTime  time.Time
		endTime    time.Time
		activeDate time.Time
	)
	activeDate = time.Now()
	if start != "" {
		startTime, err = time.Parse("2006-01-02 15:04:05", start+" 00:00:00")
		if err != nil {
			fmt.Fprintf(eout, "Start date error: %s\n", err)
			os.Exit(1)
		}
		if end == "" {
			endTime = activeDate
		} else {
			endTime, err = time.Parse("2006-01-02 15:04:05", end+" 23:59:59")
			if err != nil {
				fmt.Fprintf(eout, "End date error: %s\n", err)
				os.Exit(1)
			}
		}
	}

	reader := bufio.NewReader(in)

	entry := new(stn.Entry)
	lineNo := 0
	for {
		showLine = true
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lineNo++
		if entry.FromString(line) != true {
			fmt.Fprintf(eout, "line no. %d: can't filter [%s]\n", lineNo, line)
			os.Exit(1)
		}
		if start != "" {
			showLine = entry.IsInRange(startTime, endTime)
		}
		if showLine == true && match != "" {
			showLine = entry.IsMatch(match)
		}
		if showLine == true {
			fmt.Fprintf(out, "%s", line)
		}
	}
}
