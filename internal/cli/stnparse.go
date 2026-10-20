package cli

import (
	"fmt"
	"io"
	"strings"
	"time"

	stngo "github.com/rsdoiel/stngo"
)

/** Parse is stnparse: it reads Simple Timesheet Notation and writes the
 * entries as tab delimited lines or JSON.
 *
 * Parameters:
 *   appName (string) — the program name, for help and usage
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Parse("stnparse", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Parse(appName string, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	var o options
	var asJSON bool
	fs := newFlagSet(appName, stderr, &o)
	fs.BoolVar(&asJSON, "json", false, "output JSON format")
	if code, ok := parseFlags(fs, args); !ok {
		return code
	}
	if standard(&o, appName, ParseHelpText, stdout) {
		return ExitOK
	}
	in, out, closeAll, code := openIO(&o, fs.Args(), stdin, stdout, stderr)
	defer closeAll()
	if code != ExitOK {
		return code
	}

	activeDate := time.Now().Format(stngo.DateFmt)
	entryCnt, badLines := 0, 0
	// The entry before this one in the day, which decides morning or afternoon.
	var previous *stngo.Entry
	if asJSON {
		fmt.Fprint(out, "[")
	}
	code = eachLine(in, stderr, func(n int, line string) bool {
		switch {
		case stngo.IsDateLine(line):
			activeDate = stngo.ParseDateLine(line)
			previous = nil
		case stngo.IsEntry(line):
			entry, err := stngo.ParseEntryAfter(activeDate, line, previous)
			if err != nil {
				fmt.Fprintf(stderr, "line %d: %v\n", n, err)
				badLines++
				return false
			}
			if asJSON {
				if entryCnt > 0 {
					fmt.Fprint(out, ",")
				}
				fmt.Fprint(out, entry.JSON())
			} else {
				fmt.Fprintln(out, entry.String())
			}
			entryCnt++
			previous = entry
		case stngo.LooksLikeEntry(line):
			fmt.Fprintf(stderr, "line %d: looks like an entry but is not one, skipped: %s\n", n, excerpt(line))
		}
		return false
	})
	if asJSON {
		fmt.Fprint(out, "]")
	}
	if code != ExitOK {
		return code
	}
	if badLines > 0 {
		fmt.Fprintf(stderr, "%d line(s) could not be parsed\n", badLines)
		return ExitData
	}
	return ExitOK
}

// excerpt shortens a line for a warning.
func excerpt(line string) string {
	line = strings.TrimSpace(line)
	if r := []rune(line); len(r) > 60 {
		return string(r[:60]) + "..."
	}
	return line
}
