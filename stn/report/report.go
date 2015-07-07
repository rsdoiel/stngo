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
	"fmt"
	"strings"
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
	return false
}

// Summarize - give the output of stnparse or stnfilter aggregate the results
// by the first notation and the durration of time.
func (e *EntryAggregation) Summarize() string {
	var outText []string

	summary := make(map[string]float64)
	for _, item := range e.Entries {
		// Calc duration
		duration := item.End.Sub(item.Start)
		// Calc key
		key := item.Annotations[0]
		// if map entry does not exist create one with key and duration
		// else add the new duration to old and update map
		val, ok := summary[key]
		if ok == true {
			summary[key] = val + duration.Hours()
		} else {
			summary[key] = duration.Hours()
		}
	}
	for k, v := range summary {
		s := fmt.Sprintf("%f\t%s", v, k)
		outText = append(outText, s)
	}
	return strings.Join(outText, "\n")
}

// Detail - given the output of stnparse or stnfilter aggregate the
// results by the first annotation preserving each second to N annotations as a narrative.
func (e *EntryAggregation) Detail() string {
	return ""
}

// TotalDuration - given the output of stnfilter or stnparse calculate the total durration.
func (e *EntryAggregation) TotalDuration() string {
	return ""
}
