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
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	// My packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/stngo"
)

var (
	usage = `USAGE: %s [OPTIONS] [TIME_DESCRIPTION]`

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
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// App Options
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
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "input", "", "input filename")
	flag.StringVar(&outputFName, "o", "", "output filename")
	flag.StringVar(&outputFName, "output", "", "output filename")

	// App Options
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
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName)

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

	activeDate := time.Now().Format("2006-07-15")

	reader := bufio.NewReader(in)

	entryCnt := 0
	lineNo := 1
	if asJSON == true {
		fmt.Fprint(out, "[")
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
				fmt.Fprintf(os.Stderr, "line %d: %v\n", lineNo, perr)
			}
			if asJSON == true {
				if entryCnt > 0 {
					fmt.Fprint(out, ",")
				}
				fmt.Fprint(out, entry.JSON())
				entryCnt++
			} else {
				fmt.Fprintln(out, entry.String())
			}
		}
		lineNo++
	}
	if asJSON == true {
		fmt.Fprint(out, "]")
	}
}
