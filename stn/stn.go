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
    "errors"
	"log"
)

// Entry is the basic data element generated when parsing a file contactining
// Simple Timesheet Notation. It is designed to easily turning to JSON, CSV
// or other useful formats.
type Entry struct {
	Occurrence string
	Start      string
	End        string
	Duration   int      // in minutes
	Notations  []string // cells of contextual data (e.g. project, activity, notes)
}

// IsDateLine validates a line as appropriate to pass to ParseDateLine.
func IsDateLine(line string) bool {
	log.Fatalf("IsActiveDate() not implemented.")
	return false
}

// ParseDateLine sets the current date context when parsing Simple Timesheet Notation
// elements. It is what is recorded in Occurrence field of an Entry.
func ParseDateLine(line string) string {
	log.Fatalf("ParseActiveDate() not implemented.")
	return ""
}

// IsEntry validates a line as an "Entry" to be parsed.
func IsEntry(line string) bool {
	log.Fatalf("IsEntry() not implemented.")
	return false
}

// ParseEntry takes a string and the active date as a string and
// returns a Entry structure and error value.
func ParseEntry(activeDate string, line string) (*Entry, error) {
	return nil, errors.New("ParseEntry() not implemented.")
}
