// pipeline_test.go - runs the stnparse | stnfilter | stnreport pipeline the way
// the author's day and week report scripts do, on a synthetic timesheet that
// has the same characteristics as a decade of real logs: prose before the
// first date, dates newest first, blank lines between entries, 12 hour
// end times ("8:00 - 4:00"), entries without notes, long annotations with
// commas and colons, and both -columns and --columns spellings.
package stngo

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

const realisticTimesheet = `
Active Project Labels
---------------------

- Acme; Tools; ...

Priorities
----------

1. Thing one
2. Thing two (https://example.org/issues/42)

Timesheet
---------

2026-10-16

8:00 - 9:00; Acme; SysAdmin; Review log emails and slack. Run updates on apps, betawork, and datawork.

9:00 - 10:00; Acme; Widgets; Continue work from Tuesday, hopefully new release today.

10:00 - 11:45; Acme; Meeting; Catch up with Pat

11:45 - 3:00; Acme; Widgets; Continue work from Tuesday: "release" if tests pass, otherwise (sigh) don't.

2026-10-15

8:00 - 4:00; Time Off; Family Sick; Appointment

2026-10-14

8:00 - 4:00; Time Off; Family Sick; Appointments

2026-10-13

8:00 - 9:00; Acme; SysAdmin; Review log emails and slack.

9:00 - 10:00; Acme; Gadgets; Cloud inits

10:30 - 12:30; Acme; Gadgets; NginX install details for the migration

1:00 - 5:00; Acme; Widgets; Restore roles, cli forms, ended with test instance work.

12:30 - 1:00; Acme; Lunch

2026-10-04

8:00 - 9:30; Acme; SysAdmin; Sunday before the week reported on below
`

var (
	buildOnce sync.Once
	binDir    string
	buildErr  error
)

/** buildPrograms builds stnparse, stnfilter and stnreport once per test run.
 *
 * Parameters:
 *   t (*testing.T) — the running test
 *
 * Returns:
 *   string — the directory holding the three binaries
 *
 * Example:
 *   bin := buildPrograms(t)
 */
func buildPrograms(t *testing.T) string {
	t.Helper()
	buildOnce.Do(func() {
		binDir, buildErr = os.MkdirTemp("", "stngo-bin-")
		if buildErr != nil {
			return
		}
		for _, name := range []string{"stnparse", "stnfilter", "stnreport"} {
			cmd := exec.Command("go", "build", "-o", filepath.Join(binDir, name), "./cmd/"+name)
			if out, err := cmd.CombinedOutput(); err != nil {
				buildErr = &buildError{name: name, out: string(out), err: err}
				return
			}
		}
	})
	if buildErr != nil {
		t.Fatal(buildErr)
	}
	return binDir
}

type buildError struct {
	name string
	out  string
	err  error
}

func (e *buildError) Error() string { return "build " + e.name + ": " + e.err.Error() + "\n" + e.out }

/** run runs a built program with standard input, in UTC so the expected
 * times do not depend on the machine's zone.
 *
 * Parameters:
 *   t (*testing.T) — the running test
 *   name (string) — stnparse, stnfilter or stnreport
 *   stdin (string) — text for standard input
 *   args (...string) — command line arguments
 *
 * Returns:
 *   string — standard output
 *   string — standard error
 *   int — the exit code
 *
 * Example:
 *   out, errOut, code := run(t, "stnparse", "2026-10-16\n8:00 - 9:00; A; b\n")
 */
func run(t *testing.T, name string, stdin string, args ...string) (string, string, int) {
	t.Helper()
	return runInZone(t, "UTC", name, stdin, args...)
}

/** runInZone is run with the time zone of the process set to tz.
 *
 * Parameters:
 *   t (*testing.T) — the running test
 *   tz (string) — a zone name such as America/Los_Angeles
 *   name (string) — stnparse, stnfilter or stnreport
 *   stdin (string) — text for standard input
 *   args (...string) — command line arguments
 *
 * Returns:
 *   string — standard output
 *   string — standard error
 *   int — the exit code
 *
 * Example:
 *   out, _, _ := runInZone(t, "America/Los_Angeles", "stnparse", "2026-01-12\n8:00 - 9:00; A; b\n")
 */
func runInZone(t *testing.T, tz string, name string, stdin string, args ...string) (string, string, int) {
	t.Helper()
	if _, err := time.LoadLocation(tz); err != nil {
		t.Skipf("no time zone data for %s: %v", tz, err)
	}
	cmd := exec.Command(filepath.Join(buildPrograms(t), name), args...)
	cmd.Env = append(os.Environ(), "TZ="+tz)
	cmd.Stdin = strings.NewReader(stdin)
	var stdout, stderr strings.Builder
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	code := 0
	if ee, ok := err.(*exec.ExitError); ok {
		code = ee.ExitCode()
	} else if err != nil {
		t.Fatalf("%s: %v", name, err)
	}
	return stdout.String(), stderr.String(), code
}

// pipe runs parse, filter (with filterArgs) and report (with reportArgs) as the author's scripts do.
func pipe(t *testing.T, src string, filterArgs []string, reportArgs ...string) string {
	t.Helper()
	parsed, errOut, code := run(t, "stnparse", src)
	if code != 0 {
		t.Fatalf("stnparse exit %d: %s", code, errOut)
	}
	filtered, errOut, code := run(t, "stnfilter", parsed, filterArgs...)
	if code != 0 {
		t.Fatalf("stnfilter exit %d: %s", code, errOut)
	}
	report, errOut, code := run(t, "stnreport", filtered, reportArgs...)
	if code != 0 {
		t.Fatalf("stnreport exit %d: %s", code, errOut)
	}
	return report
}

func TestParseSkipsProseAndKeepsEveryEntry(t *testing.T) {
	out, errOut, code := run(t, "stnparse", realisticTimesheet)
	if code != 0 || errOut != "" {
		t.Fatalf("exit %d, stderr %q", code, errOut)
	}
	if got := strings.Count(out, "\n"); got != 12 {
		t.Errorf("parsed %d entries, expected 12\n%s", got, out)
	}
}

// The expected text below is worked out by hand from realisticTimesheet. Rows
// are in key order and the report ends with a blank line after the total.

func TestDayReport(t *testing.T) {
	// rpt-time-by-day.bash: stnfilter -start=D -end=D | stnreport --columns=0,1
	want := "Hours\tAnnotation(s)\n" +
		" 3.00\tAcme: Gadgets\n" +
		" 0.50\tAcme: Lunch\n" +
		" 1.00\tAcme: SysAdmin\n" +
		" 4.00\tAcme: Widgets\n" +
		"\n" +
		" 8.50\tTotal Hours\n\n"
	got := pipe(t, realisticTimesheet, []string{"-start=2026-10-13", "-end=2026-10-13"}, "--columns=0,1")
	if got != want {
		t.Errorf("day report\n got: %q\nwant: %q", got, want)
	}
}

func TestWeekReport(t *testing.T) {
	// rpt-time-by-week.bash: Sunday through Saturday, "-start D -end D" with a space.
	// 8:00 - 4:00 counts as eight hours, and 11:45 - 3:00 as 3.25. The entry
	// on 2026-10-04 is outside the week.
	want := "Hours\tAnnotation(s)\n" +
		" 3.00\tAcme: Gadgets\n" +
		" 0.50\tAcme: Lunch\n" +
		" 1.75\tAcme: Meeting\n" +
		" 2.00\tAcme: SysAdmin\n" +
		" 8.25\tAcme: Widgets\n" +
		"16.00\tTime Off: Family Sick\n" +
		"\n" +
		"31.50\tTotal Hours\n\n"
	got := pipe(t, realisticTimesheet, []string{"-start", "2026-10-11", "-end", "2026-10-17"}, "-columns=0,1")
	if got != want {
		t.Errorf("week report\n got: %q\nwant: %q", got, want)
	}
}

func TestFilterAndReportExitCodes(t *testing.T) {
	parsed, _, _ := run(t, "stnparse", realisticTimesheet)
	cases := []struct {
		name  string
		prog  string
		stdin string
		args  []string
		want  int
	}{
		{"filter bad flag", "stnfilter", parsed, []string{"-bogus"}, 2},
		{"filter bad start date", "stnfilter", parsed, []string{"-start=yesterday"}, 2},
		{"filter bad end date", "stnfilter", parsed, []string{"-start=2026-10-01", "-end=soon"}, 2},
		{"filter missing input", "stnfilter", "", []string{filepath.Join(t.TempDir(), "missing")}, 66},
		{"filter text that is not stnparse output", "stnfilter", "8:00 - 9:00; A; b\n", nil, 65},
		{"filter help", "stnfilter", "", []string{"-h"}, 0},
		{"report bad flag", "stnreport", parsed, []string{"-bogus"}, 2},
		{"report bad column list", "stnreport", parsed, []string{"-columns=a,b"}, 2},
		{"report missing input", "stnreport", "", []string{filepath.Join(t.TempDir(), "missing")}, 66},
		{"report text that is not stnparse output", "stnreport", "8:00 - 9:00; A; b\n", nil, 65},
		{"report help", "stnreport", "", []string{"-h"}, 0},
		{"report version", "stnreport", "", []string{"-v"}, 0},
	}
	for _, c := range cases {
		if _, _, code := run(t, c.prog, c.stdin, c.args...); code != c.want {
			t.Errorf("%s: exit %d, want %d", c.name, code, c.want)
		}
	}
}

const laZone = "America/Los_Angeles"

// A timesheet is written in local time: 8:00 is 8:00 whatever the date and
// whatever the date is today. Winter and summer entries differ only in offset.
func TestLocalTimeFollowsTheTimesheetNotToday(t *testing.T) {
	src := "2026-01-12\n8:00 - 9:30; A; b; winter\n2026-07-13\n8:00 - 9:30; A; b; summer\n" +
		"2026-03-08\n8:00 - 9:00; A; b; day clocks went forward\n2026-11-01\n8:00 - 4:00; A; b; day clocks went back\n"
	out, errOut, code := runInZone(t, laZone, "stnparse", src)
	if code != 0 || errOut != "" {
		t.Fatalf("exit %d, stderr %q", code, errOut)
	}
	want := []string{
		"2026-01-12T08:00:00-08:00\t2026-01-12T09:30:00-08:00",
		"2026-07-13T08:00:00-07:00\t2026-07-13T09:30:00-07:00",
		"2026-03-08T08:00:00-07:00\t2026-03-08T09:00:00-07:00",
		"2026-11-01T08:00:00-08:00\t2026-11-01T16:00:00-08:00",
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != len(want) {
		t.Fatalf("got %d entries, want %d\n%s", len(lines), len(want), out)
	}
	for i, w := range want {
		if got := strings.Join(strings.Split(lines[i], "\t")[:2], "\t"); got != w {
			t.Errorf("entry %d\n got: %s\nwant: %s", i, got, w)
		}
	}
}

// A day filter means that day on the wall clock of the timesheet, so an
// entry that starts at 6pm is still in its own day.
func TestDayFilterUsesLocalDays(t *testing.T) {
	src := "2026-10-13\n8:00 - 9:00; A; b; morning\n18:00 - 19:00; A; b; evening\n23:00 - 23:30; A; b; end of day\n"
	parsed, _, _ := runInZone(t, laZone, "stnparse", src)
	for day, want := range map[string]int{"2026-10-12": 0, "2026-10-13": 3, "2026-10-14": 0} {
		out, errOut, code := runInZone(t, laZone, "stnfilter", parsed, "-start="+day, "-end="+day)
		if code != 0 {
			t.Fatalf("stnfilter exit %d: %s", code, errOut)
		}
		if got := strings.Count(out, "\n"); got != want {
			t.Errorf("filter %s kept %d entries, want %d\n%s", day, got, want, out)
		}
	}
	// Winter, where the offset is -08:00 and 4pm is already the next UTC day.
	src = "2026-01-12\n8:00 - 9:00; A; b; morning\n16:30 - 17:30; A; b; late afternoon\n"
	parsed, _, _ = runInZone(t, laZone, "stnparse", src)
	out, _, _ := runInZone(t, laZone, "stnfilter", parsed, "-start=2026-01-12", "-end=2026-01-12")
	if strings.Count(out, "\n") != 2 {
		t.Errorf("winter filter dropped an entry\n%s", out)
	}
}

func TestEntryLikeLinesWarn(t *testing.T) {
	src := `2026-10-13

9:30 - 10:00 DLD; Misc; no semicolon after the end time
8:30 - 9:00; DLD; SysAdmin; fine
12:30 - 5:45: DLD; Feeds; colon where the semicolon belongs
3:45  - 6:00; DLD; Bags; double space before the dash
9:00 - 10:00; DLD; Misc; fine too
Notes without a time at the front are not entries.
`
	out, errOut, code := run(t, "stnparse", src)
	if code != 0 {
		t.Errorf("a warning is not a failure: exit %d, want 0", code)
	}
	if got := strings.Count(out, "\n"); got != 2 {
		t.Errorf("parsed %d entries, want 2\n%s", got, out)
	}
	for _, n := range []string{"line 3:", "line 5:", "line 6:"} {
		if !strings.Contains(errOut, n) {
			t.Errorf("no warning for %s\n%s", n, errOut)
		}
	}
	if strings.Contains(errOut, "line 4:") || strings.Contains(errOut, "line 7:") || strings.Contains(errOut, "line 8:") {
		t.Errorf("warned about a line that is fine\n%s", errOut)
	}
}

func TestLastLineWithoutNewline(t *testing.T) {
	out, _, _ := run(t, "stnparse", "2026-10-13\n8:00 - 9:00; A; b; last line, no newline")
	if strings.Count(out, "\n") != 1 {
		t.Fatalf("stnparse lost an unterminated last line\n%q", out)
	}
	out2, _, _ := run(t, "stnfilter", strings.TrimSuffix(out, "\n"))
	if strings.Count(out2, "\n") != 1 {
		t.Errorf("stnfilter lost an unterminated last line\n%q", out2)
	}
	rep, _, _ := run(t, "stnreport", strings.TrimSuffix(out, "\n"), "-columns=0,1")
	if !strings.Contains(rep, "1.00\tA: b") {
		t.Errorf("stnreport lost an unterminated last line\n%q", rep)
	}
}

// Entries in a day are in order, so an afternoon entry written 1:00 - 4:00
// after a morning one is 13:00 - 16:00, and the first entry of the next date
// starts over.
func TestAfternoonEntriesFollowTheMorning(t *testing.T) {
	src := "2026-10-14\n7:00 - 8:00; A; b; first of the day\n2026-10-13\n8:00 - 8:30; A; b; morning\n" +
		"8:30 - 12:00; A; b; more morning\n12:00 - 1:00; A; Lunch\n1:00 - 4:00; A; b; afternoon\n" +
		"7:00 - 8:00; A; b; evening\n"
	out, errOut, code := runInZone(t, laZone, "stnparse", src)
	if code != 0 || errOut != "" {
		t.Fatalf("exit %d, stderr %q", code, errOut)
	}
	var starts []string
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		f := strings.Split(line, "\t")
		starts = append(starts, f[0][11:16]+"-"+f[1][11:16])
	}
	want := "07:00-08:00 08:00-08:30 08:30-12:00 12:00-13:00 13:00-16:00 19:00-20:00"
	if got := strings.Join(starts, " "); got != want {
		t.Errorf("\n got: %s\nwant: %s", got, want)
	}
}
