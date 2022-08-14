// stnreport.go - Reporting tool for Simple Timesheet Notation. Expects input from either
// stnfilter or stnparse.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2021 all rights reserved.
// Released under the BSD 2-Clause license.
// See: http://opensource.org/licenses/BSD-2-Clause
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	// My packages

	stn "github.com/rsdoiel/stngo"
	"github.com/rsdoiel/stngo/report"
)

var (

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string
	format      string

	// App Options
	columns string
	itemize bool
)

func main() {
	appName := path.Base(os.Args[0])
	// Configure command line interface

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&itemize, "itemize", false, "report details, for JSON and CSV formats duration is in fractional hours")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")

	// App Options
	flag.StringVar(&columns, "columns", "0", "a comma delimited List of zero indexed columns to report")
	flag.StringVar(&format, "format", "", "Set the format of out, text, csv or JSON")

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

	reader := bufio.NewReader(in)

	entry := new(stn.Entry)
	aggregation := new(report.EntryAggregation)

	lineNo := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "" {
			break
		}
		lineNo++
		if entry.FromString(line) != true {
			fmt.Fprintf(eout, "line no. %d: can't filter [%s]\n", lineNo, line)
			os.Exit(1)
		} else {
			aggregation.Aggregate(entry)
		}
	}
	var cols []int
	s := strings.Split(columns, ",")
	for _, val := range s {
		i, err := strconv.Atoi(val)
		if err != nil {
			fmt.Fprintf(eout, "Column number error: %s, %s", columns, err)
			os.Exit(1)
		}
		cols = append(cols, i)
	}
	if itemize {
		fmt.Fprintln(out, aggregation.Itemize(cols, format))
	} else {
		fmt.Fprintln(out, aggregation.Summarize(cols, format))
	}
}
