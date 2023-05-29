// Package stn is a library for processing Simple Timesheet Notation.
//
// stn.go - implements a version of Simple Timesheet Notation as a Go package.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2021 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
package stn

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Version of stn.go package.
const (

	DateFmt = `2006-01-02`
)

var (
	dateLineRE       = regexp.MustCompile("^[0-9][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9]$")
	legacyDateLineRE = regexp.MustCompile("^[0-1][0-9]/[0-3][0-9]/[0-9][0-9][0-9][0-9]$")
	entryLineRE      = regexp.MustCompile("^([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]) - ([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]);")
)

// Entry is the basic data element generated when parsing a file contactining
// Simple Timesheet Notation. It is designed to easily turning to JSON, CSV
// or other useful formats.
type Entry struct {
	Start       time.Time
	End         time.Time
	Annotations []string // cells of contextual data (e.g. project, activity, notes)
}

// IsDateLine validates a line as appropriate to pass to ParseDateLine.
func IsDateLine(line string) bool {
	if dateLineRE.MatchString(strings.TrimSpace(line)) {
		return true
	}
	if legacyDateLineRE.MatchString(strings.TrimSpace(line)) {
		return true
	}
	return false
}

// ParseDateLine sets the current date context when parsing Simple Timesheet Notation
// elements. It is what is recorded in Occurrence field of an Entry.
func ParseDateLine(line string) string {
	if dateLineRE.MatchString(strings.TrimSpace(line)) {
		return strings.TrimSpace(line)
	}
	if legacyDateLineRE.MatchString(strings.TrimSpace(line)) {
		parts := strings.SplitN(strings.TrimSpace(line), "/", 3)
		return fmt.Sprintf("%s-%s-%s", parts[2], parts[0], parts[1])
	}
	return ""
}

// IsEntry validates a line as an "Entry" to be parsed.
func IsEntry(line string) bool {
	if entryLineRE.MatchString(strings.TrimSpace(line)) {
		return true
	}
	return false
}

func splitCells(line string) []string {
	return strings.Split(line, ";")
}

func splitRangeElements(timeRange string) (string, string, error) {
	if strings.Index(timeRange, " - ") != -1 {
		parts := strings.SplitN(timeRange, " - ", 2)
		return parts[0], parts[1], nil
	}
	return "", "", errors.New("[" + timeRange + "] is not a valid time range string. ")
}

func parseRangeElements(start string, end string) (time.Time, time.Time, error) {
	startTime, err1 := time.Parse("2006-01-02 15:04 MST", start)
	endTime, err2 := time.Parse("2006-01-02 15:04 MST", end)
	//NOTE: need to handle the case where someone has entered an end time ran
	// smaller than start (e.g. 8:00 - 1:00 meaning 1pm should become 13:00)
	if startTime.Unix() > endTime.Unix() {
		plus12hr, _ := time.ParseDuration("+12h")
		endTime = endTime.Add(plus12hr)
	}
	if err1 != nil {
		return startTime, endTime, err1
	}
	if err2 != nil {
		return startTime, endTime, err2
	}
	return startTime, endTime, nil
}

// ParseEntry takes a string and the active date as a string and
// returns a Entry structure and error value.
func ParseEntry(activeDate string, line string) (*Entry, error) {
	if IsDateLine(activeDate) == false {
		return nil, errors.New("invalid format for active date")
	}
	if IsEntry(line) == false {
		return nil, errors.New("invalid format for entry")
	}
	cells := splitCells(line)
	if len(cells) < 2 {
		return nil, errors.New("entry line missing cells")
	}

	s, e, err := splitRangeElements(cells[0])
	if err != nil {
		return nil, err
	}

	// NOTE: for now I am assume timesheets are in local time.
	// Need to think about supporting other timezone for things like
	// timesheets during event travel.
	zone, _ := time.Now().Zone()
	start, end, err := parseRangeElements(activeDate+" "+s+" "+zone,
		activeDate+" "+e+" "+zone)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(cells); i++ {
		cells[i] = strings.TrimSpace(cells[i])
	}

	var entry *Entry
	entry = &Entry{
		Start:       start,
		End:         end,
		Annotations: cells[1:],
	}
	return entry, nil
}

// JSON converts an Entry struct to JSON notation.
func (e *Entry) JSON() string {
	src, _ := json.Marshal(e)
	return string(src)
}

// String converts an Entry struct to a tab delimited string.
func (e *Entry) String() string {
	return e.Start.Format(time.RFC3339) + "\t" + e.End.Format(time.RFC3339) +
		"\t" + strings.Join(e.Annotations[:], "\t")
}

// FromString reads a tab delimited string formatted with Stringback into a Entry struct
func (e *Entry) FromString(line string) bool {
	var err error
	parts := strings.Split(line, "\t")
	if len(parts) < 3 {
		return false
	}
	e.Start, err = time.Parse(time.RFC3339, parts[0])
	if err != nil {
		return false
	}
	e.End, err = time.Parse(time.RFC3339, parts[1])
	if err != nil {
		return false
	}
	e.Annotations = parts[2:]
	return true
}

// IsInRange checks the start and end times of an Entry structure to see if it is in the time range
func (e *Entry) IsInRange(start time.Time, end time.Time) bool {
	t1 := e.Start.Unix()
	if t1 >= start.Unix() && t1 <= end.Unix() {
		return true
	}
	return false
}

// IsMatch checks the Entry struct Annotations for matching substring
func (e *Entry) IsMatch(match string) bool {
	matched := false
	//NOTE: search all columns
	for i := 0; i < len(e.Annotations); i++ {
		if strings.Contains(e.Annotations[i], match) == true {
			matched = true
			break
		}
	}
	return matched
}
