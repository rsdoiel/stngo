/**
 * report.go - provides basic reporting features for stnparse output.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license
 * See: http://opensource.org/licenses/BSD-2-Clause
 */

// Package report provides basic report rendering for Simple Timesheet Notation.
package report

import (
    "../../stn"
    "log"
)

// EntryAggregation
type EntryAggregation struct {
    Entries []stn.Entry
}
// Aggregate - add an entry to the EntryAggregate
func (a *EntryAggregation) Aggregate(entry *stn.Entry) bool {
    i := len(a.Entries)
    a.Entries = append(a.Entries, *entry)
    if len(a.Entries) == (i + 1) {
        return true
    }
    return false;
}

// Summarize - give the output of stnparse or stnfilter aggregate the results
// by the first notation and the durration of time.
func (e *EntryAggregation) Summarize() string {
    log.Fatal("EntryAggregation.Summarize() not implemetned.")
    return ""
}

// Detail - given the output of stnparse or stnfilter aggregate the
// results by the first annotation preserving each second to N annotations as a narrative.
func (e *EntryAggregation) Detail() string {
    log.Fatal("EntryAggregation.Detail() not implemetned.")
    return ""
}

// TotalDuration - given the output of stnfilter or stnparse calculate the total durration.
func (e *EntryAggregation) TotalDurration() string {
    log.Fatal("EntryAggregation.TotalDurration() not implemetned.")
    return ""
}
