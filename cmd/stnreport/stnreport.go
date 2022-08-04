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
	"fmt"
	"os"
	"strconv"
	"strings"

	// My packages
	"github.com/rsdoiel/stngo"
	"github.com/rsdoiel/stngo/report"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
)

var (
	synopsis = `
%s renders parsed standard timesheet notation reports
`

	description = `
%s takes output from stnparse or stnfilter and renders a
report.
`

	examples = `
This renders columns zero (first column) and one.

` + "```" + `
    stnparse -i TimeSheet.txt | %s -columns 0,1
` + "```" + `

This renders columns zero (first column) and one as a CSV file.

` + "```" + `
    stnparse -i TimeSheet.txt | %s -columns 0,1 -format csv
` + "```" + `


`

	// Standard Options
	showHelp         bool
	showExamples     bool
	showLicense      bool
	showVersion      bool
	inputFName       string
	outputFName      string
	format           string
	quiet            bool
	generateMarkdown bool
	generateManPage  bool

	// App Options
	columns string
	itemize bool
)

func main() {
	// Configure command line interface
	app := cli.NewCli(stn.Version)
	appName := app.AppName()

	// Add Help Docs
	app.AddHelp("license", []byte(fmt.Sprintf(stn.LicenseText, appName, stn.Version)))
	app.AddHelp("synopsis", []byte(fmt.Sprintf(synopsis, appName)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.BoolVar(&itemize, "itemize", false, "report details, for JSON and CSV formats duration is in fractional hours")
	app.StringVar(&inputFName, "i,input", "", "input filename")
	app.StringVar(&outputFName, "o,output", "", "output filename")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")

	// App Options
	app.StringVar(&columns, "c,columns", "0", "a comma delimited List of zero indexed columns to report")
	app.StringVar(&format, "f,format", "", "Set the format of out, text, csv or JSON")

	app.Parse()
	args := app.Args()

	// Setup IO
	var err error
	app.Eout = os.Stderr

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Handle Options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintln(app.Out, app.Help(args...))
		} else if showExamples {
			fmt.Fprintln(app.Out, app.Help("examples"))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}

	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}

	reader := bufio.NewReader(app.In)

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
			fmt.Fprintf(app.Eout, "line no. %d: can't filter [%s]\n", lineNo, line)
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
			fmt.Fprintf(app.Eout, "Column number error: %s, %s", columns, err)
			os.Exit(1)
		}
		cols = append(cols, i)
	}
	if itemize {
		fmt.Fprintln(app.Out, aggregation.Itemize(cols, format))
	} else {
		fmt.Fprintln(app.Out, aggregation.Summarize(cols, format))
	}
}
