// Package cli holds the command line behavior of stnparse, stnfilter and
// stnreport so that plain go test covers it. The programs under cmd/ only
// call Parse, Filter and Report.
package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	stngo "github.com/rsdoiel/stngo"
)

// The exit codes of every command line tool in the workspace (DR-0003).
const (
	ExitOK         = 0
	ExitUsage      = 2
	ExitData       = 65
	ExitNoInput    = 66
	ExitCantCreate = 73
	ExitIO         = 74
)

// options are the options every program takes.
type options struct {
	help, license, version bool
	input, output          string
}

// newFlagSet makes a FlagSet that reports its errors on stderr and registers
// the standard options, each in a short and a long spelling.
func newFlagSet(appName string, stderr io.Writer, o *options) *flag.FlagSet {
	fs := flag.NewFlagSet(appName, flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		fmt.Fprintf(stderr, "Usage: %s [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]\nTry %s -help for the manual.\n", appName, appName)
	}
	fs.BoolVar(&o.help, "h", false, "display help")
	fs.BoolVar(&o.help, "help", false, "display help")
	fs.BoolVar(&o.license, "l", false, "display license")
	fs.BoolVar(&o.license, "license", false, "display license")
	fs.BoolVar(&o.version, "v", false, "display version")
	fs.BoolVar(&o.version, "version", false, "display version")
	fs.StringVar(&o.input, "i", "", "input filename")
	fs.StringVar(&o.output, "o", "", "output filename")
	return fs
}

// standard answers -help, -license and -version. It reports whether it did.
func standard(o *options, appName string, helpText string, stdout io.Writer) bool {
	switch {
	case o.help:
		fmt.Fprintln(stdout, stngo.FmtHelp(helpText, appName, stngo.Version, stngo.ReleaseDate, stngo.ReleaseHash))
	case o.license:
		fmt.Fprintln(stdout, strings.TrimSpace(stngo.LicenseText))
	case o.version:
		fmt.Fprintf(stdout, "%s %s %s\n", appName, stngo.Version, stngo.ReleaseHash)
	default:
		return false
	}
	return true
}

// openIO picks the input and output: a named file, or standard input and
// output when there is none or the name is "-". The positional arguments
// are the input and output names when the options are not given. The
// returned function closes what was opened.
func openIO(o *options, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (io.Reader, io.Writer, func(), int) {
	closer := func() {}
	if len(args) > 2 {
		fmt.Fprintf(stderr, "surplus argument %q\n", args[2])
		return nil, nil, closer, ExitUsage
	}
	if o.input == "" && len(args) > 0 {
		o.input = args[0]
	}
	if o.output == "" && len(args) > 1 {
		o.output = args[1]
	}
	in, out := stdin, stdout
	var files []*os.File
	closer = func() {
		for _, f := range files {
			f.Close()
		}
	}
	if o.input != "" && o.input != "-" {
		f, err := os.Open(o.input)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return nil, nil, closer, ExitNoInput
		}
		files = append(files, f)
		in = f
	}
	if o.output != "" && o.output != "-" {
		f, err := os.Create(o.output)
		if err != nil {
			fmt.Fprintln(stderr, err)
			closer()
			return nil, nil, func() {}, ExitCantCreate
		}
		files = append(files, f)
		out = f
	}
	return in, out, closer, ExitOK
}

// eachLine calls fn with the number and text of each line, the last one
// included when it has no newline at the end. It returns ExitIO if reading
// fails part way, and otherwise whatever fn last returned when asked to stop.
func eachLine(in io.Reader, stderr io.Writer, fn func(n int, line string) (stop bool)) int {
	reader := bufio.NewReader(in)
	for n := 1; ; n++ {
		line, err := reader.ReadString('\n')
		if line != "" && fn(n, line) {
			return ExitOK
		}
		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Fprintln(stderr, err)
				return ExitIO
			}
			return ExitOK
		}
	}
}

// parseFlags parses args and returns an exit code and whether to carry on.
func parseFlags(fs *flag.FlagSet, args []string) (int, bool) {
	if err := fs.Parse(args); err != nil {
		return ExitUsage, false
	}
	return ExitOK, true
}
