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
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	// This package
	"github.com/rsdoiel/stngo/report"
	"github.com/rsdoiel/stngo/stn"
)

func revision() {
	fmt.Printf("%s %s\n", filepath.Base(os.Args[0]), stn.Version)
	os.Exit(0)
}

func main() {
	var (
		version bool
		columns string
	)

	flag.BoolVar(&version, "version", false, "Display version information.")
	flag.BoolVar(&version, "v", false, "Display version information.")
	flag.StringVar(&columns, "columns", "0", "A comma delimited List of zero indexed columns to report")

	flag.Parse()
	if version == true {
		revision()
	}

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
	fmt.Println(aggregation.Summarize(cols))
}
