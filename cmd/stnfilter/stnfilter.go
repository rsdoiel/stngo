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
	"../../stn"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	start  string
	end    string
	match  string
	asJSON bool
	help   bool
)

var usage = func(exit_code int, msg string) {
	var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	cmdName := os.Args[0]

	fmt.Fprintf(fh, `%s
USAGE %s [options]

Filter the output from stnparse based on date and/or matching text string.

EXAMPLE

Filter TimeSheet.tab from July 4, 2015 through July 14, 2015
and render a stream of JSON blobs.

    %s -start 2015-07-04 -end 2015-07-14 -json < TimeSheet.tab

To render the same in a tab delimited output

    %s -start 2015-07-04 -end 2015-07-14 < TimeSheet.tab

Typical usage would be in a pipeline with Unix cat and stnparse

   cat Time_Sheet.txt | stnparse | %s -start 2015-07-06 -end 2015-07-010

Matching a project name "Fred" for the same week would look like

    cat Time_Sheet.txt | stnparse | %s -start 2015-07-06 -end 2015-07-010 -match Fred

OPTIONS
`, msg, cmdName, cmdName, cmdName, cmdName, cmdName)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "\t-%s\t\t%s\n", f.Name, f.Usage)
	})

	fmt.Fprintf(fh, `
copyright (c) 2015 all rights reserved.
Released under the BSD 2-Clause license.
See: http://opensource.org/licenses/BSD-2-Clause
`)
	os.Exit(exit_code)
}

func revision() {
	fmt.Printf("%s %s\n", filepath.Base(os.Args[0]), stn.Version)
	os.Exit(0)
}

func main() {
	var (
		version    bool
		showLine   = true
		startTime  time.Time
		endTime    time.Time
		activeDate time.Time
		err        error
	)

	flag.StringVar(&match, "match", "", "Match text in annotations.")
	flag.StringVar(&start, "start", "", "Start of inclusive date range.")
	flag.StringVar(&end, "end", "", "End of inclusive date range.")
	flag.BoolVar(&asJSON, "json", false, "Output in JSON format.")
	flag.BoolVar(&help, "help", false, "Display this help document.")
	flag.BoolVar(&help, "h", false, "Display this help document.")
	flag.BoolVar(&version, "v", false, "Display version information.")
	flag.Parse()
	if help == true {
		usage(0, "")
	}

	if version == true {
		revision()
	}

	activeDate = time.Now()
	if start != "" {
		startTime, err = time.Parse("2006-01-02 15:04:05", start+" 00:00:00")
		if err != nil {
			log.Fatalf("Start date error: %s\n", err)
			os.Exit(1)
		}
		if end == "" {
			endTime = activeDate
		} else {
			endTime, err = time.Parse("2006-01-02 15:04:05", end+" 23:59:59")
			if err != nil {
				log.Fatalf("End date error: %s\n", err)
				os.Exit(1)
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)

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
			log.Fatalf("line no. %d: can't filter [%s]\n", lineNo, line)
			os.Exit(1)
		}
		if start != "" {
			showLine = entry.IsInRange(startTime, endTime)
		}
		if showLine == true && match != "" {
			showLine = entry.IsMatch(match)
		}
		if showLine == true {
			fmt.Printf("%s", line)
		}
	}
}
