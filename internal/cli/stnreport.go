package cli

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	stngo "github.com/rsdoiel/stngo"
	"github.com/rsdoiel/stngo/report"
)

/** Report is stnreport: it totals the hours in stnparse or stnfilter output.
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
 *   os.Exit(cli.Report("stnreport", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Report(appName string, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	var o options
	var columns, format string
	var itemize bool
	fs := newFlagSet(appName, stderr, &o)
	fs.BoolVar(&itemize, "itemize", false, "report each entry, for JSON and CSV formats duration is in fractional hours")
	fs.StringVar(&columns, "columns", "0", "a comma delimited list of zero indexed columns to report")
	fs.StringVar(&format, "format", "", "set the format of the output, text, csv or json")
	if code, ok := parseFlags(fs, args); !ok {
		return code
	}
	if standard(&o, appName, ReportHelpText, stdout) {
		return ExitOK
	}

	var cols []int
	for val := range strings.SplitSeq(columns, ",") {
		i, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil || i < 0 {
			fmt.Fprintf(stderr, "Column number error: %s\n", columns)
			return ExitUsage
		}
		cols = append(cols, i)
	}

	in, out, closeAll, code := openIO(&o, fs.Args(), stdin, stdout, stderr)
	defer closeAll()
	if code != ExitOK {
		return code
	}

	aggregation := new(report.EntryAggregation)
	badLine := 0
	code = eachLine(in, stderr, func(n int, line string) bool {
		entry := new(stngo.Entry)
		if !entry.FromString(line) {
			fmt.Fprintf(stderr, "line no. %d: can't report [%s]\n", n, excerpt(line))
			badLine = n
			return true
		}
		aggregation.Aggregate(entry)
		return false
	})
	if code != ExitOK {
		return code
	}
	if badLine > 0 {
		return ExitData
	}
	if itemize {
		fmt.Fprintln(out, aggregation.Itemize(cols, format))
	} else {
		fmt.Fprintln(out, aggregation.Summarize(cols, format))
	}
	return ExitOK
}
