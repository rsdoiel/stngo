//
// Package report provides basic report rendering for Simple Timesheet Notation.
//
// report.go - provides basic reporting features for stnparse output.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2015 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
//
package report

import (
	"../stn"
	"fmt"
	"sort"
	"strings"
)

// EntryAggregation - a structure to hold the aggregated entries.
type EntryAggregation struct {
	Entries []stn.Entry
}

// Aggregate - add an entry to the EntryAggregate
func (e *EntryAggregation) Aggregate(entry *stn.Entry) bool {
	i := len(e.Entries)
	e.Entries = append(e.Entries, *entry)
	if len(e.Entries) == (i + 1) {
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
	outText = append(outText, "Hours\tProject")
	total := 0.0
	keys := make([]string, len(summary))
	for k := range summary {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v, ok := summary[k]
		if ok == true {
			total += v
			outText = append(outText, fmt.Sprintf("%5.2f\t%s", v, k))
		}
	}
	outText = append(outText, "")
	outText = append(outText, fmt.Sprintf("%5.2f\tTotal Hours", total))
	outText = append(outText, "")
	return strings.Join(outText, "\n")
}
