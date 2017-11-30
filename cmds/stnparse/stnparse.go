//
// stnparse.go - Simple Timesheet Notation parser.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2015 all rights reserved.
// Released under the BSD 2-Clause license.
// See: http://opensource.org/licenses/BSD-2-Clause
//
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	// My packages
	"github.com/rsdoiel/stngo"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
)

var (
	description = `

SYNOPSIS

%s parses content in "Standard Timesheet Notation". By default
it parse them into a tabular format but can also optionally
parse them into a stream of JSON blobs.

`

	examples = `

EXAMPLES

	%s < TimeSheet.txt

This will parse the TimeSheet.txt file into a table.

	%s -json < TimeSheet.txt

This will parse TimeSheet.txt file into a stream of JSON blobs.

`

	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	inputFName   string
	outputFName  string

	// App Options
	asJSON bool
)

func main() {
	// Configuration and command line interation
	app := cli.NewCli(stn.Version)
	appName := app.AppName()

	// Document expected parameters (non-option args)
	app.AddParams("[TIME_DESCRIPTION]")

	app.AddHelp("license", []byte(fmt.Sprintf(stn.LicenseText, appName, stn.Version)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")

	// App Options
	app.BoolVar(&asJSON, "j,json", false, "output JSON format")

	app.Parse()
	args := app.Args()

	var err error
	app.In, err = cli.Open(inputFName, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, app.Out)
	app.Err = os.Stderr

	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(app.Help(args...))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}

	if showExamples == true {
		if len(args) > 0 {
			fmt.Println(app.Help(args...))
		} else {
			fmt.Println(app.Help("examples"))
		}
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Println(app.License())
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(app.Version())
		os.Exit(0)
	}

	activeDate := time.Now().Format("2006-07-15")

	reader := bufio.NewReader(app.In)

	entryCnt := 0
	lineNo := 1
	if asJSON == true {
		fmt.Fprint(app.Out, "[")
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
				fmt.Fprintf(app.Err, "line %d: %v\n", lineNo, perr)
			}
			if asJSON == true {
				if entryCnt > 0 {
					fmt.Fprint(app.Out, ",")
				}
				fmt.Fprint(app.Out, entry.JSON())
				entryCnt++
			} else {
				fmt.Fprintln(app.Out, entry.String())
			}
		}
		lineNo++
	}
	if asJSON == true {
		fmt.Fprint(app.Out, "]")
	}
}
