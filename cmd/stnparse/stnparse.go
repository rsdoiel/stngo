/**
 * stnparse.go - Simple Timesheet Notation parser.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license.
 * See: http://opensource.org/licenses/BSD-2-Clause
 */
package main

import (
	"../../stn/stn"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	help bool
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

Parse TimeSheet.txt and render a stream of JSON blobs.

    %s -json < timeSheet.txt 


OPTIONS
`, msg, cmdName, cmdName)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "\t-%s\t\t%s\n", f.Name, f.Usage)
	})

	fmt.Fprintf(fh, `
copyright (c) 2015 all rights reserved.
Released under the Simplified BSD License
See: http://opensource.org/licenses/BSD-2-Clause
`)
	os.Exit(exit_code)
}

func main() {
	flag.BoolVar(&help, "h", help, "Display this help document.")
	flag.Parse()
	if help == true {
		usage(0, "")
	}

	activeDate := time.Now().Format("2006-07-15")

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if stn.IsDateLine(line) == true {
			activeDate, _ := stn.ParseDateLine(line)
		} else if stn.IsEntry(line) {
			entry, _ := stn.ParseEntry(line)
		}

		if shorthand.IsAssignment(line) {
			shorthand.Assign(line)
		} else {
			fmt.Print(shorthand.Expand(line))
		}
	}
}
