/**
 * stn.go - implements a version of Simple Timesheet Notation as a Go package.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license
 * See: http://opensource.org/licenses/BSD-2-Clause
 */

// Package stn is a library for processing Simple Timesheet Notation.
package stn

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	dateLineRE  = regexp.MustCompile("^[0-9][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9]$")
	entryLineRE = regexp.MustCompile("^([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]) - ([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]);")
)

//FIXME: need to decide how to set timezone offset, assume local unless +/- single time
// referenence found? +8:00, +12:00, -4:00

// Entry is the basic data element generated when parsing a file contactining
// Simple Timesheet Notation. It is designed to easily turning to JSON, CSV
// or other useful formats.
type Entry struct {
	Start      time.Time
	End        time.Time
	Annotations  []string // cells of contextual data (e.g. project, activity, notes)
}

// IsDateLine validates a line as appropriate to pass to ParseDateLine.
func IsDateLine(line string) bool {
	if dateLineRE.MatchString(strings.TrimSpace(line)) {
		return true
	}
	return false
}

// ParseDateLine sets the current date context when parsing Simple Timesheet Notation
// elements. It is what is recorded in Occurrence field of an Entry.
func ParseDateLine(line string) string {
	if IsDateLine(line) {
		return strings.TrimSpace(line)
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
	//FIXME: need to handle the case where someone has entered an end time range
	// smaller than start (e.g. 8:00 - 1:00 meaning 1pm)
	startTime, err1 := time.Parse("2006-01-02 15:04", start)
	endTime, err2 := time.Parse("2006-01-02 15:04", end)
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
		return nil, errors.New("active date misformatted.")
	}
	if IsEntry(line) == false {
		return nil, errors.New("entry line misformatted.")
	}
	cells := splitCells(line)
	if len(cells) < 2 {
		return nil, errors.New("entry line missing cells")
	}

	s, e, err := splitRangeElements(cells[0])
	if err != nil {
		return nil, err
	}

	start, end, err := parseRangeElements(activeDate + " " + s, activeDate + " " + e)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(cells); i += 1 {
		cells[i] = strings.TrimSpace(cells[i])
	}

	var entry *Entry
	entry = &Entry{
		Start:      start,
		End:        end,
		Annotations:  cells[1:],
	}
	return entry, nil
}

func (e *Entry) JSON() string {
	src, _ := json.Marshal(e)
	return string(src)
}

func (e *Entry) String() string {
	return e.Start.Format(time.RFC3339) + "\t" + e.End.Format(time.RFC3339) +
		"\t" + strings.Join(e.Annotations[:], "\t")
}
