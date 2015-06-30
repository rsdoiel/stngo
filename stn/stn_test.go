/**
 * stn_test.go - implements tests for stn Go package.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license
 * See: http://opensource.org/licenses/BSD-2-Clause
 */
package stn

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func ok(t *testing.T, expected bool, msg string) {
		if expected != true {
			t.Fatalf("Failed: " + msg)
		}
}

func TestIsDateLine(t *testing.T) {
	// Simple text expected the dates in YYYY-MM-DD
	text := `2015-07-04`
	expected := true
	result := IsDateLine(text)
	ok(t, expected == result, text + " is valid.")

	// IsDateLine expects dates in YYYYY-MM-DD format.
	text = `January 20th, 2015`
	expected = false
	result = IsDateLine(text)
	ok(t, expected == result, text + " is NOT valid.")

	// Also not valid for IsDateLine...
	text = `07-25-2015`
	expected = false
	result = IsDateLine(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is an entry not a DateLine
	text = "08:00 - 9:00; misc; email and what not."
	expected = false
	result = IsDateLine(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is an empty line, not a DateLine
	text = ""
	expected = false
	result = IsDateLine(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	expected = false
	result = IsDateLine(text)
	ok(t, expected == result, text + " is NOT valid")
}

func TestParseDateLine(t *testing.T) {
	// Simple text expected the dates in YYYY-MM-DD
	text := `2015-07-04`
	expected := `2015-07-04`
	result := ParseDateLine(text)
	ok(t, expected == result, text + " is valid. Got: " + result)

	// IsDateLine expects dates in YYYYY-MM-DD format.
	text = `January 20th, 2015`
	expected = ""
	result = ParseDateLine(text)
	ok(t, expected == result, text + " is NOT valid.")

	// Also not valid for IsDateLine...
	text = `07-25-2015`
	expected = ""
	result = ParseDateLine(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is an entry not a DateLine
	text = "08:00 - 9:00; misc; email and what not."
	expected = ""
	result = ParseDateLine(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is an empty line, not a DateLine
	text = ""
	expected = ""
	result = ParseDateLine(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	expected = ""
	result = ParseDateLine(text)
	ok(t, expected == result, text + " is NOT valid")
}

func TestIsEntry(t *testing.T) {
	// Simple text expected the dates in YYYY-MM-DD
	text := `2015-07-04`
	expected := false
	result := IsEntry(text)
	ok(t, expected == result, text + " is valid.")

	// IsDateLine expects dates in YYYYY-MM-DD format.
	text = `January 20th, 2015`
	expected = false
	result = IsEntry(text)
	ok(t, expected == result, text + " is NOT valid.")

	// Also not valid for IsDateLine...
	text = `07-25-2015`
	expected = false
	result = IsEntry(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is an entry not a DateLine
	text = "08:00 - 9:00; misc; email and what not."
	expected = true
	result = IsEntry(text)
	ok(t, expected == result, text + " is valid, got: " + strconv.FormatBool(result))

	// This is an empty line, not a DateLine
	text = ""
	expected = false
	result = IsEntry(text)
	ok(t, expected == result, text + " is NOT valid")

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	expected = false
	result = IsEntry(text)
	ok(t, expected == result, text + " is NOT valid")
}

func TestParseEntry(t *testing.T) {
	cells := splitCells("one; two; three")
	ok(t, len(cells) == 3, "Should get three cells: " + strings.Join(cells, " | "))

	activeDate := "2015-07-04"
	// Simple text expected the dates in YYYY-MM-DD
	text := `2015-07-04`
	_, err := ParseEntry(activeDate, text)
	ok(t, err != nil, text + " produced error on ParseEntry().")

	// IsDateLine expects dates in YYYYY-MM-DD format.
	text = `January 20th, 2015`
	_, err = ParseEntry(activeDate, text)
	ok(t, err != nil, text + " produced error on ParseEntry().")

	// Also not valid for IsDateLine...
	text = `07-25-2015`
	_, err = ParseEntry(activeDate, text)
	ok(t, err != nil, text + " produced error on ParseEntry().")

	// This is an entry not a DateLine
	text = "08:00 - 9:00; misc; email and what not."
	_, err = ParseEntry(activeDate, text)
	ok(t, err == nil, fmt.Sprintf("%s  is Valid, got error: %q", text, err))
	ok(t, false, "FIXME: Need to test Entry() return structure and make sure it converts nicely into String and JSON.")

	// This is an empty line, not a DateLine
	text = ""
	_, err = ParseEntry(activeDate, text)
	ok(t, err != nil, text + " produced error on ParseEntry().")

	// This is just some random text, not a DateLine
	text = "This is just some random text, not a DateLine"
	_, err = ParseEntry(activeDate, text)
	ok(t, err != nil, text + " produced error on ParseEntry().")
}
