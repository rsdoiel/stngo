package cli

import (
	"fmt"
	"io"
	"strings"
	"time"

	stngo "github.com/rsdoiel/stngo"
)

/** Filter is stnfilter: it passes on the lines of stnparse output that start
 * inside a range of days and match some text.
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
 *   os.Exit(cli.Filter("stnfilter", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Filter(appName string, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	var o options
	var start, end, match string
	var ignored bool
	fs := newFlagSet(appName, stderr, &o)
	fs.StringVar(&match, "match", "", "match text annotations")
	fs.StringVar(&start, "start", "", "start of inclusive date range")
	fs.StringVar(&end, "end", "", "end of inclusive date range")
	fs.BoolVar(&ignored, "json", false, "accepted and ignored")
	if code, ok := parseFlags(fs, args); !ok {
		return code
	}
	if standard(&o, appName, FilterHelpText, stdout) {
		return ExitOK
	}

	// The days are days on the wall clock of the timesheet, in local time.
	var startTime, endTime time.Time
	var err error
	if start != "" {
		if startTime, err = stngo.DayStart(start); err != nil {
			fmt.Fprintf(stderr, "Start date error: %s\n", err)
			return ExitUsage
		}
	}
	if end != "" {
		if endTime, err = stngo.DayEnd(end); err != nil {
			fmt.Fprintf(stderr, "End date error: %s\n", err)
			return ExitUsage
		}
	} else if start != "" {
		endTime = time.Now()
	}
	inRange := start != "" || end != ""

	in, out, closeAll, code := openIO(&o, fs.Args(), stdin, stdout, stderr)
	defer closeAll()
	if code != ExitOK {
		return code
	}

	entry := new(stngo.Entry)
	badLine := 0
	code = eachLine(in, stderr, func(n int, line string) bool {
		if !entry.FromString(line) {
			fmt.Fprintf(stderr, "line no. %d: can't filter [%s]\n", n, excerpt(line))
			badLine = n
			return true
		}
		show := true
		if inRange {
			show = entry.IsInRange(startTime, endTime)
		}
		if show && match != "" {
			show = entry.IsMatch(match)
		}
		if show {
			fmt.Fprintln(out, strings.TrimRight(line, "\r\n"))
		}
		return false
	})
	if code != ExitOK {
		return code
	}
	if badLine > 0 {
		return ExitData
	}
	return ExitOK
}
