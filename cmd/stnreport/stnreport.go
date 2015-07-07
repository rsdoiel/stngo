/**
 * stnreport.go - Reporting tool for Simple Timesheet Notation. Expects input from either
 * stnfilter or stnparse.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license.
 * See: http://opensource.org/licenses/BSD-2-Clause
 */
package main

import (
	"fmt"
	"os"
	"../../stn"
	"../../stn/report"
	"bufio"
	"log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	entry := new(stn.Entry)
	aggregation := new(report.EntryAggregation)

	lineNo := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "" {
			break
		}
		lineNo++
		//fmt.Printf("DEBUG line %d: [%s]", lineNo, line)
		if entry.FromString(line) != true {
			log.Fatalf("line no. %d: can't filter [%s]\n", lineNo, line)
			os.Exit(1)
		} else {
			aggregation.Aggregate(entry)
		}
	}
	fmt.Println(aggregation.Summarize())
}
