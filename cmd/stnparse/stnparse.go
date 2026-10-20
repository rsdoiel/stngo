// stnparse is part of stngo, tools for Simple Timesheet Notation. Run stnparse -help.
package main

import (
	"os"
	"path"

	"github.com/rsdoiel/stngo/internal/cli"
)

func main() {
	os.Exit(cli.Parse(path.Base(os.Args[0]), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
