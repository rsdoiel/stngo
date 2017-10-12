//
// stnreport.go - Reporting tool for Simple Timesheet Notation. Expects input from either
// stnfilter or stnparse.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2015 all rights reserved.
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
	"strconv"
	"strings"

	// My packages
	"github.com/rsdoiel/stngo"
	"github.com/rsdoiel/stngo/report"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
)

var (
	usage = `USAGE: %s [OPTIONS]`

	description = `

SYNOPSIS

%s takes output from stnparse or stnfilter and renders a
report.

`

	examples = `
EXAMPLE

    stnparse -i TimeSheet.txt | %s -columns 0,1

This renders columns zero (first column) and one.

`

	// Standard Options
	showHelp     bool
	showExamples bool
	showLicense  bool
	showVersion  bool
	inputFName   string
	outputFName  string

	// App Options
	columns string
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Options
	flag.StringVar(&columns, "c", "0", "a comma delimited List of zero indexed columns to report")
	flag.StringVar(&columns, "columns", "0", "a comma delimited List of zero indexed columns to report")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "STN", stn.Version)
	cfg.LicenseText = fmt.Sprintf(stn.LicenseText, appName, stn.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName)

	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(cfg.Help(args...))
		} else {
			fmt.Println(cfg.Usage())
		}
		os.Exit(0)
	}
	if showExamples == true {
		if len(args) > 0 {
			fmt.Println(cfg.Example(args...))
		} else {
			fmt.Println(cfg.ExampleText)
		}
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	in, err := cli.Open(inputFName, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, out)

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
			fmt.Fprintf(os.Stderr, "line no. %d: can't filter [%s]\n", lineNo, line)
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
			fmt.Fprintf(os.Stderr, "Column number error: %s, %s", columns, err)
			os.Exit(1)
		}
		cols = append(cols, i)
	}
	fmt.Fprintln(out, aggregation.Summarize(cols))
}
