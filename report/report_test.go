//
// Package report provides basic report rendering for Simple Timesheet Notation.
//
// report_test.go test the stn/report package.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2015 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
//
package report

import (
	"fmt"
	"log"
	"strings"
	"testing"

	// local packages
	"github.com/rsdoiel/ok"
	"github.com/rsdoiel/stngo"
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
	outText := aggregation.Summarize([]int{0})
	outLines := strings.Split(outText, "\n")
	ok.Ok(t, len(outLines) == 8, fmt.Sprintf("lines %d: [%s]\n", linesTotal, outText))
}
