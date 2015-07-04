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
	column string
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
		showLine  = true
		startTime = time.Time
		endTime   = time.Time
	)

	flag.Var(&start, "start", start, "start date range.")
	flag.Var(&end, "end", end, "end, inclusive, date range.")
	flag.VarInt(&column, "column", column, "by annotation column.")
	flag.Var(&match, "match", match, "column value should match.")
	flag.BoolVar(&asJSON, "json", asJSON, "Output as JSON format.")
	flag.BoolVar(&help, "h", help, "Display this help document.")
	flag.Parse()
	if help == true {
		usage(0, "")
	}

	activeDate := time.Now().Format("2006-01-02")
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

	for {
		showLine = true
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		entry, err := stn.ParseTab(line)
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		if showLine == true && match != "" {
			showLine = IsMatch(entry, column, match)
		}
		if showLine == true && start != "" && end != "" {
			showLine = IsInRange(entry, startTime, endTime)
		}
		if sbowLine == true {
			fmt.Printf("%s\n", line)
		}
	}
}
