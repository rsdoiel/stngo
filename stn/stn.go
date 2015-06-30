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
    "fmt"
    "strings"
    "time"
    "errors"
    "regexp"
)

var (
    dateLineRE = regexp.MustCompile("^[0-9][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9]$")
    entryLineRE = regexp.MustCompile("^([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]) - ([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]);")
)

// Entry is the basic data element generated when parsing a file contactining
// Simple Timesheet Notation. It is designed to easily turning to JSON, CSV
// or other useful formats.
type Entry struct {
	Occurrence time.Time
	Start      time.Time
	End        time.Time
	Duration   int      // in minutes
	Notations  []string // cells of contextual data (e.g. project, activity, notes)
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

func splitRange(line string) (time.Time, time.Time, error) {
    parts := strings.SplitN(line, " - ", 2)
    fmt.Println("DEBUG parts -> " + parts[0] + ", " + parts[1])
    start, err1 := time.Parse("15:04", parts[0])
    end, err2 := time.Parse("15:04", parts[1])
    if err1 != nil {
        return start, end, err1
    }
    if err2 != nil {
        return start, end, err2
    }
    return start, end, nil
}

// ParseEntry takes a string and the active date as a string and
// returns a Entry structure and error value.
func ParseEntry(activeDate string, line string) (*Entry, error) {

    if IsDateLine(activeDate) == false {
        return nil, errors.New("active date misformatted.")
    }
    occurrence, err := time.Parse("2006-01-02", activeDate)
    if err != nil {
        return nil, errors.New("Problem parsing active date: " + err.Error())
    }
    if IsEntry(line) == false {
        return nil, errors.New("entry line misformatted.")
    }
    cells := splitCells(line)
    if len(cells) < 2 {
        return nil, errors.New("entry line missing cells")
    }
    start, end, err := splitRange(cells[0])
    if err != nil {
        return nil, err
    }
    var entry *Entry
    entry = &Entry{
        Occurrence: occurrence,
        Start: start,
        End: end,
        Notations: []string{"DEBUG notations"},
    }
    /*
    entry.Occurrence = occurrence
    entry.Start = start
    entry.End = end
    entry.Notations = []string{"DEBUG Notations"}
    */

    fmt.Printf("DEBUG %s, start: %s, end: %s\n", occurrence.Format("2006-01-02"), start.Format("15:04"), end.Format("15:04"))
    //return &Entry{Occurrence: occurrence, Start: start, End: end, Notations: []string{"DEBUG values"}}, nil
    return entry, nil
}
