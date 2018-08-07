//
// stnfilter.go - Simple Timesheet Notation filter.
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
	synopsis = `
%s a standard timesheet notation filter.
`
	description = `
%s will filter the output from stnparse based on date or matching text.
`

	examples = `
Filter TimeSheet.tab from July 4, 2015 through July 14, 2015
and render a stream of JSON blobs.

` + "```" + `
    %s -start 2015-07-04 -end 2015-07-14 -json < TimeSheet.tab
` + "```" + `

To render the same in a tab delimited output

` + "```" + `
    %s -start 2015-07-04 -end 2015-07-14 < TimeSheet.tab
` + "```" + `

Typical usage would be in a pipeline with Unix cat and stnparse

` + "```" + `
   cat Time_Sheet.txt | stnparse | %s -start 2015-07-06 -end 2015-07-010
` + "```" + `

Matching a project name "Fred" for the same week would look like

` + "```" + `
    cat Time_Sheet.txt | stnparse | %s -start 2015-07-06 -end 2015-07-010 -match Fred
` + "```" + `

`

	// Standard Options
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	inputFName           string
	outputFName          string
	quiet                bool
	generateMarkdownDocs bool
	generateManPage      bool

	// App Options
	start  string
	end    string
	match  string
	asJSON bool
)

func main() {
	// Configuration and command line interation
	app := cli.NewCli(stn.Version)
	appName := app.AppName()

	// Add some Help docs
	app.AddHelp("license", []byte(fmt.Sprintf(stn.LicenseText, appName, stn.Version)))
	app.AddHelp("synopsis", []byte(fmt.Sprintf(synopsis, appName)))
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display examples(s)")
	app.StringVar(&inputFName, "i,input", "", "input file name")
	app.StringVar(&outputFName, "o,output", "", "output file name")
	app.BoolVar(&quiet, "quiet", false, "suppress error message")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")

	// App Options
	app.StringVar(&match, "m,match", "", "Match text annotations")
	app.StringVar(&start, "s,start", "", "start of inclusive date range")
	app.StringVar(&end, "e,end", "", "end of inclusive date range")
	app.BoolVar(&asJSON, "j,json", false, "output JSON format")

	// Run the command line interface
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
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
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
			fmt.Fprintf(app.Eout, "Start date error: %s\n", err)
			os.Exit(1)
		}
		if end == "" {
			endTime = activeDate
		} else {
			endTime, err = time.Parse("2006-01-02 15:04:05", end+" 23:59:59")
			if err != nil {
				fmt.Fprintf(app.Eout, "End date error: %s\n", err)
				os.Exit(1)
			}
		}
	}

	reader := bufio.NewReader(app.In)

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
			fmt.Fprintf(app.Eout, "line no. %d: can't filter [%s]\n", lineNo, line)
			os.Exit(1)
		}
		if start != "" {
			showLine = entry.IsInRange(startTime, endTime)
		}
		if showLine == true && match != "" {
			showLine = entry.IsMatch(match)
		}
		if showLine == true {
			fmt.Fprintf(app.Out, "%s", line)
		}
	}
}
