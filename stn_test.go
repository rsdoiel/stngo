//
// Package stn is a library for processing Simple Timesheet Notation.
//
// stn_test.go - implements tests for stn Go package.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2021 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
//
package stn

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
	if err != nil {
		t.Errorf("%q produced error on ParseEntry() - %s.", text, err)
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
	if err != nil {
		t.Errorf("%q produced error on ParseEntry() - %s", text, err)
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
