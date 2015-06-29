/**
 * stnshorthand.go - A command line utility to process an shorthand definitions
 * and render the resulting file with the transformed text and without any shorthand definitions.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license.
 */
package main

import (
	"fmt"
	"os"
	"flag"
    "bufio"
	"../../shorthand"
)

var (
    help bool
    expression string
)

var usage = func(exit_code int, msg string) {
    var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
USAGE %s [options]

%s reads from standard in and writes to standard out.

EXAMPLE

%s -e "@now := $(date +%H:%M)" < myfile.txt > output.txt

 OPTIONS
`, msg, os.Args[0], os.Args[0], os.Args[0], os.Args[0])

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "\t-%s\t\t%s\n", f.Name, f.Usage)
	})

	fmt.Fprintf(fh, `
 copyright (c) 2015 all rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php
`)
	os.Exit(exit_code)
}

func init() {
	const (
		expressionUsage = "The shorthand you wish at add. E.g. -e \"@now = $(date +%H:%M)\""
		helpUsage       = "Display this help document."
	)

    //FIXME: Need to support multiple -e options like sed.
    flag.StringVar(&expression, "e", expression, expressionUsage)
	flag.BoolVar(&help, "help", help, helpUsage)
	flag.BoolVar(&help, "h", help, helpUsage)
}

func main() {
    flag.Parse()
    if help == true {
        usage(0, "")
    }

    //FIXME: for each -e add the assignment.
    if expression != "" {
        if shorthand.IsAssignment(expression) {
            shorthand.Assign(expression)
        } else {
            usage(1, "Error processing expression: " + expression)
        }
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        if shorthand.IsAssignment(line) {
            shorthand.Assign(line)
        } else {
            fmt.Print(shorthand.Expand(line))
        }
    }
}
