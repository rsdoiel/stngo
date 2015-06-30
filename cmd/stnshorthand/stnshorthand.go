/**
 * stnshorthand.go - A command line utility to process an shorthand definitions
 * and render the resulting file with the transformed text and without any shorthand definitions.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license.
 */
package main

import (
	"errors"
	"fmt"
	"os"
	"flag"
    "bufio"
	"../../shorthand"
)

type expressionList []string

var (
    help bool
    expression expressionList
)

var usage = func(exit_code int, msg string) {
    var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
USAGE stnshorthand [options]

stnshorthand reads from standard in and writes to standard out.

Shorthands are defined by a label followed by a space, followed by a colon,
equal sign and another space followed by a value and end of line. The
shorthand is the label, it is replaced by the value (not including the
end of line).

    LABEL := VALUE

To create a shortand for the label "ACME" with the value
"the point at which someone or something is best"
would be done with the following line

    ACME := the point at which someone or something is best

Now each time the shorthand "ACME" is encountered the phrase
"the point at which someone or something is best" will replace it. This the
text

    My, ACME, will come

Would become

    My, the point at which someone or something is best, will come

Normally you would use shorthands for things like long project names,
passing dynamic values (like the current time or date).


EXAMPLE

    stnshorthand -e "@now := $(date +%%H:%%M)" \
	   -e "@today := $(date +%%Y-%%m-%%d)" < myfile.txt > output.txt


OPTIONS
`, msg)

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


func (e *expressionList) String () string {
	return fmt.Sprintf("%s", *e)
}

func (e *expressionList) Set(value string) error {
	if shorthand.IsAssignment(value) == false {
		return errors.New("Shorthand is not valid (LABEL := VALUE)")
	}
	shorthand.Assign(value)
	return nil
}


func init() {
	const (
		expressionUsage = "The shorthand notation(s) you wish at add."
		helpUsage       = "Display this help document."
	)

    flag.Var(&expression, "e", expressionUsage)
	flag.BoolVar(&help, "help", help, helpUsage)
	flag.BoolVar(&help, "h", help, helpUsage)
}

func main() {
    flag.Parse()
    if help == true {
        usage(0, "")
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
