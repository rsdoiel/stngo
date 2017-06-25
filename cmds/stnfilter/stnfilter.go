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
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	// My packages
	"github.com/rsdoiel/stngo"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
)

var (
	usage = `USAGE: %s [OPTIONS]`

	description = `
SYNOPSIS

%s will filter the output from stnparse based on date or matching text.
`

	examples = `
EXAMPLES

Filter TimeSheet.tab from July 4, 2015 through July 14, 2015
and render a stream of JSON blobs.

    %s -start 2015-07-04 -end 2015-07-14 -json < TimeSheet.tab

To render the same in a tab delimited output

    %s -start 2015-07-04 -end 2015-07-14 < TimeSheet.tab

Typical usage would be in a pipeline with Unix cat and stnparse

   cat Time_Sheet.txt | stnparse | %s -start 2015-07-06 -end 2015-07-010

Matching a project name "Fred" for the same week would look like

    cat Time_Sheet.txt | stnparse | %s -start 2015-07-06 -end 2015-07-010 -match Fred
`

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

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")

	// App Options
	flag.StringVar(&match, "m", "", "match text annotations")
	flag.StringVar(&match, "match", "", "Match text annotations")
	flag.StringVar(&start, "s", "", "start of inclusive date range")
	flag.StringVar(&start, "start", "", "start of inclusive date range")
	flag.StringVar(&end, "e", "", "end of inclusive date range")
	flag.StringVar(&end, "end", "", "end of inclusive date range")
	flag.BoolVar(&asJSON, "j", false, "output JSON format")
	flag.BoolVar(&asJSON, "json", false, "output JSON format")
}
func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	// Configuration and command line interation
	cfg := cli.New(appName, "STN", fmt.Sprintf(stn.LicenseText, appName, stn.Version), stn.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName)

	if showHelp == true {
		fmt.Println(cfg.Usage())
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
			fmt.Fprintf(os.Stderr, "Start date error: %s\n", err)
			os.Exit(1)
		}
		if end == "" {
			endTime = activeDate
		} else {
			endTime, err = time.Parse("2006-01-02 15:04:05", end+" 23:59:59")
			if err != nil {
				fmt.Fprintf(os.Stderr, "End date error: %s\n", err)
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
			fmt.Fprintf(os.Stderr, "line no. %d: can't filter [%s]\n", lineNo, line)
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
