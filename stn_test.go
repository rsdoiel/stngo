// Package stn is a library for processing Simple Timesheet Notation.
//
// stn_test.go - implements tests for stn Go package.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2021 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
package stngo

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestIsDateLine(t *testing.T) {
	// Simple text expected the dates in YYYY-MM-DD
	text := `2015-07-04`
	expected := true
	result := IsDateLine(text)
	if expected != result {
		t.Errorf("%q is valid.", text)
	}

	// Simple text expected the dates in MM/DD/YYYY
	text = `07/04/2015`
	expected = true
	result = IsDateLine(text)
	if expected != result {
		t.Errorf("%q is valid", text)
	}

	// IsDateLine expects dates in YYYYY-MM-DD format.
	text = `January 20th, 2015`
	expected = false
	result = IsDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// Also not valid for IsDateLine...
	text = `07-25-2015`
	expected = false
	result = IsDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// This is an entry not a DateLine
	text = "08:00 - 9:00; misc; email and what not."
	expected = false
	result = IsDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// This is an empty line, not a DateLine
	text = ""
	expected = false
	result = IsDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	expected = false
	result = IsDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}
}

func TestParseDateLine(t *testing.T) {
	// Simple text expected the dates in YYYY-MM-DD
	text := `2015-07-04`
	expected := `2015-07-04`
	result := ParseDateLine(text)
	if expected != result {
		t.Errorf("%q is valid. Get: %q", text, result)
	}

	// IsDateLine expects dates in YYYYY-MM-DD format.
	text = `January 20th, 2015`
	expected = ""
	result = ParseDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// Also not valid for IsDateLine...
	text = `07-25-2015`
	expected = ""
	result = ParseDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// This is an entry not a DateLine
	text = "08:00 - 9:00; misc; email and what not."
	expected = ""
	result = ParseDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// This is an empty line, not a DateLine
	text = ""
	expected = ""
	result = ParseDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	expected = ""
	result = ParseDateLine(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}
}

func TestIsEntry(t *testing.T) {
	// Simple text expected the dates in YYYY-MM-DD
	text := `2015-07-04`
	expected := false
	result := IsEntry(text)
	if expected != result {
		t.Errorf("%q is valid.", text)
	}

	// IsDateLine expects dates in YYYYY-MM-DD format.
	text = `January 20th, 2015`
	expected = false
	result = IsEntry(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// Also not valid for IsDateLine...
	text = `07-25-2015`
	expected = false
	result = IsEntry(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}

	// This is an entry not a DateLine
	text = "08:00 - 9:00; misc; email and what not."
	expected = true
	result = IsEntry(text)
	if expected != result {
		t.Errorf("%q is valid, got: %q", text, strconv.FormatBool(result))
	}

	// This is an empty line, not a DateLine
	text = ""
	expected = false
	result = IsEntry(text)
	if expected != result {
		t.Errorf("%q is NOT valid", text)
	}

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	expected = false
	result = IsEntry(text)
	if expected != result {
		t.Errorf("%q is NOT valid.", text)
	}
}

func TestParseEntry(t *testing.T) {
	cells := splitCells("one; two; three")
	if len(cells) != 3 {
		t.Errorf("Should get three cells: %q", strings.Join(cells, " | "))
	}

	activeDate := "2015-07-04"
	// Simple text expected the dates in YYYY-MM-DD
	// but this is not as an entry
	text := `2015-07-04`
	_, err := ParseEntry(activeDate, text)
	if err == nil {
		t.Errorf("%q produced no error on ParseEntry().", text)
	}

	// niether date nor entry
	text = `January 20th, 2015`
	_, err = ParseEntry(activeDate, text)
	if err == nil {
		t.Errorf("%q produced no error on ParseEntry()", text)
	}

	// Also not valid IsDateLine/entry...
	text = `07-25-2015`
	_, err = ParseEntry(activeDate, text)
	if err == nil {
		t.Errorf("%q produced no error on ParseEntry()", text)
	}

	// This is an entry
	text = "08:00 - 9:30; misc; email and what not."
	entry, err := ParseEntry(activeDate, text)
	if err != nil {
		t.Errorf("%s is Valid, got error: %q", text, err)
	}
	if entry.Start.Hour() != 8 {
		t.Errorf("should start at hour of 8, got %d", entry.Start.Hour())
	}
	if entry.Start.Minute() != 0 {
		t.Errorf("should have start minute 0, %d", entry.Start.Minute())
	}
	if entry.End.Hour() != 9 {
		t.Errorf("should end at hour of 9, got %d", entry.End.Hour())
	}
	if entry.End.Minute() != 30 {
		t.Errorf("should have end minute 30, got %d", entry.End.Minute())
	}
	if len(entry.Annotations) != 2 {
		t.Errorf("Should have two annoations, got %d", len(entry.Annotations))
	}
	if len(entry.Annotations) == 2 {
		if entry.Annotations[0] != "misc" {
			t.Errorf("first cell should be 'misc': [%q]", entry.Annotations[0])
		}
		if entry.Annotations[1] != "email and what not." {
			t.Errorf("first cell should be 'email and what not.': [%q]", entry.Annotations[1])
		}
	}

	jsonString := entry.JSON()
	expectedString := `{"Start":"2015-07-04T08:00:00-07:00","End":"2015-07-04T09:30:00-07:00","Annotations":["misc","email and what not."]}`
	if jsonString != expectedString {
		t.Errorf("entry.toJSON(): %s", jsonString)
	}

	text = entry.String()
	expectedString = "2015-07-04T08:00:00-07:00\t2015-07-04T09:30:00-07:00\tmisc\temail and what not."
	if text != expectedString {
		t.Errorf("entry.String(): %q", text)
	}

	text = "08:22 - 1:34; afternoon; email and what not."
	entry, err = ParseEntry(activeDate, text)
	jsonString = entry.JSON()
	expectedString = `{"Start":"2015-07-04T08:22:00-07:00","End":"2015-07-04T13:34:00-07:00","Annotations":["afternoon","email and what not."]}`
	if jsonString != expectedString {
		t.Errorf("entry.toJSON(): %q", jsonString)
	}
	text = entry.String()
	expectedString = "2015-07-04T08:22:00-07:00\t2015-07-04T13:34:00-07:00\tafternoon\temail and what not."
	if text != expectedString {
		t.Errorf("entry.String(): %q", text)
	}

	// This is an empty line, not a DateLine
	text = ""
	_, err = ParseEntry(activeDate, text)
	if err == nil {
		t.Errorf("%q produced no error on ParseEntry().", text)
	}

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	_, err = ParseEntry(activeDate, text)
	if err == nil {
		t.Errorf("%q produced error on ParseEntry().", text)
	}
}

func TestFilter(t *testing.T) {
	start, _ := time.Parse("2006-01-02", "2015-07-01")
	end, _ := time.Parse("2006-01-02", "2015-07-31")
	s, _ := time.Parse("2006-01-02 15:04", "2015-07-04 08:38")
	e, _ := time.Parse("2006-01-02 15:04", "2015-07-04 13:34")
	t1 := Entry{
		Start:       s,
		End:         e,
		Annotations: []string{"one", "two"},
	}
	expected := true
	result := t1.IsInRange(start, end)
	if expected != result {
		t.Errorf("%q is between %q and %q", t1.String(), start.String(), end.String())
	}

	s, _ = time.Parse("2006-01-02 15:04", "2015-06-04 08:38")
	e, _ = time.Parse("2006-01-02 15:04", "2015-06-04 13:34")
	t1.Start = s
	t1.End = e
	expected = false
	result = t1.IsInRange(start, end)
	if expected != result {
		t.Errorf("%q not is between %q and %q", t1.String(), start.String(), end.String())
	}

	expected = true
	result = t1.IsMatch("one")
	if expected != result {
		t.Errorf("one is an annotation %v, %v", expected, result)
	}

	expected = true
	result = t1.IsMatch("two")
	if expected != result {
		t.Errorf("one is an annotation, %v, %v", expected, result)
	}

	expected = false
	result = t1.IsMatch("three")
	if expected != result {
		t.Errorf("one is an annotation, %v, %v", expected, result)
	}

	t2 := new(Entry)
	if t2.FromString(t1.String()) != true {
		t.Errorf("FromString should work %v", t2.FromString(t1.String()))
	}
	if t1.Start != t2.Start {
		t.Errorf("Start should match, %q == %q", t1.Start, t2.Start)
	}
	if t1.End != t2.End {
		t.Errorf("End should match, %q == %q", t1.End, t2.End)
	}
	for i := 0; i < len(t1.Annotations); i++ {
		if t1.Annotations[i] != t2.Annotations[i] {
			t.Errorf("%s == %s failed\n",
				t1.Annotations[i], t2.Annotations[i])
		}
	}
}

/** TestParseEntryEndOfDay checks that 24:00 is accepted as midnight at the
 * end of the day, e.g. "23:30 - 24:00; DLD; Misc; shutdown systems".
 *
 * Example:
 *   go test -run TestParseEntryEndOfDay
 */
func TestParseEntryEndOfDay(t *testing.T) {
	entry, err := ParseEntry("2015-07-04", "23:30 - 24:00; DLD; Misc; shutdown systems")
	if err != nil {
		t.Fatalf("23:30 - 24:00 should parse, got %v", err)
	}
	if got := entry.End.Sub(entry.Start); got != 30*time.Minute {
		t.Errorf("duration is %s, expected 30m", got)
	}
	if entry.End.Day() != 5 || entry.End.Hour() != 0 || entry.End.Minute() != 0 {
		t.Errorf("end is %s, expected midnight starting 2015-07-05", entry.End)
	}
}

/** TestFromStringDropsTheLineEnding checks that the newline read with a line
 * does not end up inside the last annotation, which made an entry with no
 * notes ("8:00 - 9:00; Acme; Lunch") a different report row from the same
 * activity written with notes.
 *
 * Example:
 *   go test -run TestFromStringDropsTheLineEnding
 */
func TestFromStringDropsTheLineEnding(t *testing.T) {
	for _, ending := range []string{"\n", "\r\n", ""} {
		e := new(Entry)
		line := "2026-10-13T12:30:00Z\t2026-10-13T13:00:00Z\tAcme\tLunch" + ending
		if !e.FromString(line) {
			t.Fatalf("FromString(%q) = false", line)
		}
		if last := e.Annotations[len(e.Annotations)-1]; last != "Lunch" {
			t.Errorf("last annotation is %q for line ending %q, expected \"Lunch\"", last, ending)
		}
	}
}

/** TestParseEntryAfterResolvesMorningAndAfternoon checks that a start time
 * written without am/pm takes the reading, morning or afternoon,
 * that keeps the start times of the day ascending.
 *
 * Example:
 *   go test -run TestParseEntryAfterResolvesMorningAndAfternoon
 */
func TestParseEntryAfterResolvesMorningAndAfternoon(t *testing.T) {
	day := "2026-10-13"
	at := func(clock string) time.Time {
		d, err := time.ParseInLocation("2006-01-02 15:04", day+" "+clock, time.Local)
		if err != nil {
			t.Fatal(err)
		}
		return d
	}
	// prev is an entry of the day that started at start and ended at end.
	prev := func(start, end string) *Entry { return &Entry{Start: at(start), End: at(end)} }
	cases := []struct {
		name       string
		prev       *Entry
		line       string
		start, end string
	}{
		{"first entry of the day is as written", nil, "1:00 - 5:00; A; b", "01:00", "05:00"},
		{"first entry before 7am is as written", nil, "6:30 - 8:00; A; b", "06:30", "08:00"},
		{"crossing noon after a morning entry", prev("08:00", "08:30"), "1:00 - 4:00; A; b", "13:00", "16:00"},
		{"later in the morning stays in the morning", prev("08:00", "08:30"), "9:00 - 10:00; A; b", "09:00", "10:00"},
		{"lunch hour: 12:15 is noon", prev("10:00", "11:45"), "12:15 - 1:15; A; b", "12:15", "13:15"},
		{"an overlap stays where it was written", prev("10:00", "10:30"), "10:00 - 11:00; A; b", "10:00", "11:00"},
		{"the same start as the entry before stays", prev("08:30", "14:45"), "8:30 - 5:30; A; b", "08:30", "17:30"},
		{"an afternoon entry follows an afternoon entry", prev("13:00", "14:00"), "2:00 - 3:00; A; b", "14:00", "15:00"},
		{"an evening session", prev("13:00", "17:00"), "7:00 - 10:00; A; b", "19:00", "22:00"},
		{"8:00 - 4:00 is 8:00 - 16:00 as the first entry", nil, "8:00 - 4:00; A; b", "08:00", "16:00"},
		{"8:00 - 4:00 is 8:00 - 16:00 after an earlier start", prev("07:00", "08:00"), "8:00 - 4:00; A; b", "08:00", "16:00"},
		{"8:00 - 4:00 is 8:00 - 16:00 after the same start", prev("08:00", "08:30"), "8:00 - 4:00; A; b", "08:00", "16:00"},
		{"8:30 - 1:00 ends at 13:00", nil, "8:30 - 1:00; A; b", "08:30", "13:00"},
		{"12:30 - 1:00 ends at 13:00", nil, "12:30 - 1:00; A; b", "12:30", "13:00"},
		{"an afternoon start keeps its end: 1:00 - 5:00", prev("08:00", "12:00"), "1:00 - 5:00; A; b", "13:00", "17:00"},
		{"an entry added out of order is not moved to 9pm", prev("09:15", "09:30"), "9:00 - 3:00; A; b", "09:00", "15:00"},
		{"a start that leaves no room for the end is not moved", prev("13:30", "15:15"), "3:15 - 3:00; A; b", "03:15", "15:00"},
		{"24 hour notation is never moved", prev("08:00", "08:30"), "13:00 - 15:00; A; b", "13:00", "15:00"},
		{"an entry from another day is no guide", &Entry{Start: at("08:00").AddDate(0, 0, 1), End: at("08:30").AddDate(0, 0, 1)}, "1:00 - 4:00; A; b", "01:00", "04:00"},
	}
	for _, c := range cases {
		entry, err := ParseEntryAfter(day, c.line, c.prev)
		if err != nil {
			t.Errorf("%s: %v", c.name, err)
			continue
		}
		if !entry.Start.Equal(at(c.start)) || !entry.End.Equal(at(c.end)) {
			t.Errorf("%s: %q gave %s - %s, want %s - %s", c.name, c.line,
				entry.Start.Format("15:04"), entry.End.Format("15:04"), c.start, c.end)
		}
	}
	// ParseEntry has no previous entry to go by.
	entry, _ := ParseEntry(day, "1:00 - 4:00; A; b")
	if entry.Start.Hour() != 1 {
		t.Errorf("ParseEntry should read as written, got %s", entry.Start.Format("15:04"))
	}
}
