// Package report provides basic report rendering for Simple Timesheet Notation.
//
// report.go - provides basic reporting features for stnparse output.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2015 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
package report

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	// stn packages
	"github.com/rsdoiel/stngo"
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

func composeKey(entry *stn.Entry, indexes []int) string {
	var s []string
	for _, col := range indexes {
		if col < len(entry.Annotations) {
			s = append(s, entry.Annotations[col])
		}
	}
	switch len(indexes) {
	case 1:
		return strings.Join(s, " ")
	case 2:
		return strings.Join(s, ": ")
	default:
		return strings.Join(s, "; ")
	}
}

func formatFloat(val float64, format string) string {
	var out string
	switch strings.ToLower(format) {
	case "json":
		out = fmt.Sprintf("%0.2f", val)
	case "csv":
		out = fmt.Sprintf("%0.2f", val)
	default:
		out = fmt.Sprintf("%5.2f", val)
	}
	return out
}

// Summarize - give the output of stnparse or stnfilter aggregate the
// results by the first notation, second notation and durration of time.
func (e *EntryAggregation) Summarize(columns []int, format string) string {
	summary := make(map[string]float64)
	for _, item := range e.Entries {
		// Calc duration
		duration := item.End.Sub(item.Start)
		// Calc key
		key := composeKey(&item, columns)
		// if map entry does not exist create one with key and duration
		// else add the new duration to old and update map
		val, ok := summary[key]
		if ok == true {
			summary[key] = val + duration.Hours()
		} else {
			summary[key] = duration.Hours()
		}
	}
	keys := make([]string, len(summary))
	for k := range summary {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	total := 0.0
	// We're going to write an 2D slice of strings since that
	// is the common tabular data structure.
	records := [][]string{
		{"Hours", "Annotation(s)"},
	}
	var col1 string
	for _, k := range keys {
		v, ok := summary[k]
		if ok == true {
			total += v
			col1 = formatFloat(v, format)
			records = append(records, []string{
				col1,
				fmt.Sprintf("%s", k),
			})
		}
	}
	records = append(records, []string{})

	records = append(records, []string{
		formatFloat(total, format),
		"Total Hours",
	})
	records = append(records, []string{})

	// Now output an appropriate format, default is text,
	// JSON and CSV are options.
	switch strings.ToLower(format) {
	case "csv":
		buf := new(bytes.Buffer)
		w := csv.NewWriter(buf)
		w.WriteAll(records)
		return buf.String()
	case "json":
		src, _ := json.MarshalIndent(records, "", "    ")
		return string(src)
	default:
		outText := []string{}
		for _, record := range records {
			outText = append(outText, strings.Join(record, "\t"))
		}
		return strings.Join(outText, "\n")
	}
}
