/**
 * stnfilter.go - Simple Timesheet Notation filter.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license.
 * See: http://opensource.org/licenses/BSD-2-Clause
 */
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

EXAMPLE

Filter TimeSheet.tab from July 4, 2015 through July 14, 2015
and render a stream of JSON blobs.

    %s -start 2015-07-04 -end 2015-07-14 -json < timeSheet.tab 


OPTIONS
`, msg, cmdName, cmdName)

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

func main() {
	var (
		showLine   = true
		startTime  time.Time
		endTime    time.Time
		activeDate time.Time
		err        error
	)

	flag.StringVar(&start, "start", "", "start date range.")
	flag.StringVar(&end, "end", "", "end, inclusive, date range.")
	flag.StringVar(&match, "match", "", "column value should match.")
	flag.BoolVar(&asJSON, "json", false, "Output as JSON format.")
	flag.BoolVar(&help, "h", false, "Display this help document.")
	flag.BoolVar(&help, "help", false, "Display this help document.")
	flag.Parse()
	if help == true {
		usage(0, "")
	}

	activeDate = time.Now()
	if start != "" {
		startTime, err = time.Parse("2006-01-02", start)
		if err != nil {
			log.Fatalf("Start date error: %s\n", err)
			os.Exit(1)
		}
		if end == "" {
			endTime = activeDate
		} else {
			endTime, err = time.Parse("2006-01-02", end)
			if err != nil {
				log.Fatalf("End date error: %s\n", err)
				os.Exit(1)
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)

	entry := new(stn.Entry)
	line_no := 0
	for {
		showLine = true
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line_no += 1
		if entry.FromString(line) != true {
			log.Fatalf("line no. %d: can't filter [%s]\n", line_no, line)
			os.Exit(1)
		}
		if start != "" {
			showLine = entry.IsInRange(startTime, endTime)
		}
		if showLine == true && match != "" {
			showLine = entry.IsMatch(match)
		}
		if showLine == true {
			fmt.Printf("%s\n", line)
		}
	}
}
