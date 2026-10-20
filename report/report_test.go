// Package report provides basic report rendering for Simple Timesheet Notation.
//
// report_test.go test the stn/report package.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2015 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
package report

import (
	"log"
	"strings"
	"testing"

	// local packages

	stn "github.com/rsdoiel/stngo"
)

func TestAggregator(t *testing.T) {
	text := `2015-07-06T08:00:00-07:00	2015-07-06T08:30:00-07:00	misc	email, update basecamp
2015-07-06T08:30:00-07:00	2015-07-06T11:00:00-07:00	ArchivesSpace	running through migration process, updating notes, testing
2015-07-06T11:00:00-07:00	2015-07-06T11:45:00-07:00	misc	update Mac
2015-07-03T08:00:00-07:00	2015-07-03T15:30:00-07:00	Holiday	4th of July observed
2015-07-02T07:45:00-07:00	2015-07-02T09:30:00-07:00	misc	email, review stuff
2015-07-02T09:30:00-07:00	2015-07-02T10:30:00-07:00	DLD meeting
2015-07-02T10:30:00-07:00	2015-07-02T12:00:00-07:00	ArchivesSpace	running through migration process
2015-07-02T03:00:00-07:00	2015-07-02T03:30:00-07:00	ArchivesSpace	Hangouts with Tommy to upgrade cls-arch.library.caltech.edu to v1.3.0, go over migration questions
2015-07-01T07:45:00-07:00	2015-07-01T09:30:00-07:00	ArchivesSpace	continue reading docs, articles about approach and what problems are being solved.
2015-07-01T09:30:00-07:00	2015-07-01T11:00:00-07:00	ArchivesSpace	meeting in SFL's MCR (3rd floot Multi-media Conference Room)`
	aggregation := new(EntryAggregation)
	entry := new(stn.Entry)
	lines := strings.Split(text, "\n")
	linesTotal := 0
	for i, line := range lines {
		if entry.FromString(line) == true {
			if aggregation.Aggregate(entry) != true {
				log.Fatalf("Can't aggregate entry %d: %v", i, entry)
			}
			linesTotal++
		} else {
			log.Fatalf("Can't read line no. %d: [%s]\n", i, line)
		}
	}
	outText := aggregation.Summarize([]int{0}, "")
	outLines := strings.Split(outText, "\n")
	if len(outLines) != 8 {
		t.Errorf("lines %d: [%s]\n", linesTotal, outText)
	}
}

func aggregateOf(t *testing.T, lines ...string) *EntryAggregation {
	t.Helper()
	aggregation := new(EntryAggregation)
	for i, line := range lines {
		entry := new(stn.Entry)
		if !entry.FromString(line) {
			t.Fatalf("can't read line %d: %q", i+1, line)
		}
		aggregation.Aggregate(entry)
	}
	return aggregation
}

const (
	itemizeFirst  = "2026-10-13T08:00:00-07:00\t2026-10-13T09:00:00-07:00\tAcme\tWidgets\tfirst\n"
	itemizeSecond = "2026-10-13T09:00:00-07:00\t2026-10-13T09:30:00-07:00\tAcme\tLunch\n"
)

/** TestItemizeKeepsEveryEntry checks that the header row is added to the
 * entries and does not replace the first one.
 *
 * Example:
 *   go test -run TestItemizeKeepsEveryEntry ./report
 */
func TestItemizeKeepsEveryEntry(t *testing.T) {
	a := aggregateOf(t, itemizeFirst, itemizeSecond)
	want := "Duration\tPrimary\tSecondary\n" +
		" 1.00\tAcme\tWidgets\n" +
		" 0.50\tAcme\tLunch"
	if got := a.Itemize([]int{0, 1}, ""); got != want {
		t.Errorf("text\n got: %q\nwant: %q", got, want)
	}
	wantCSV := "Duration,Primary,Secondary\n1.00,Acme,Widgets\n0.50,Acme,Lunch\n"
	if got := a.Itemize([]int{0, 1}, "csv"); got != wantCSV {
		t.Errorf("csv\n got: %q\nwant: %q", got, wantCSV)
	}
	if got := a.Itemize([]int{0, 1}, "json"); strings.Count(got, "Widgets") != 1 || strings.Count(got, "Lunch") != 1 {
		t.Errorf("json is missing an entry: %s", got)
	}
}

func TestItemizeOneEntry(t *testing.T) {
	a := aggregateOf(t, itemizeFirst)
	want := "Duration\tPrimary\n 1.00\tAcme"
	if got := a.Itemize([]int{0}, ""); got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSummarizeMissingColumnAndFormats(t *testing.T) {
	a := aggregateOf(t, itemizeFirst, itemizeSecond)
	if got := a.Summarize([]int{0, 1}, "csv"); !strings.Contains(got, "1.00,Acme: Widgets") || !strings.Contains(got, "1.50,Total Hours") {
		t.Errorf("csv summary: %q", got)
	}
	if got := a.Itemize([]int{5}, ""); !strings.Contains(got, "N/A") {
		t.Errorf("a column that is not there should be N/A: %q", got)
	}
	if got := a.Summarize([]int{5}, ""); !strings.Contains(got, "1.50\tTotal Hours") {
		t.Errorf("a column that is not there should not lose the hours: %q", got)
	}
	if got := a.Summarize([]int{0}, "json"); !strings.Contains(got, "Total Hours") {
		t.Errorf("json summary: %q", got)
	}
}
