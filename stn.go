// Package stngo is a library for processing Simple Timesheet Notation.
//
// stn.go - implements a version of Simple Timesheet Notation as a Go package.
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2021 all rights reserved.
// Released under the BSD 2-Clause license
// See: http://opensource.org/licenses/BSD-2-Clause
package stngo

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	/** DateFmt is the layout of a date line, e.g. 2026-10-16.
	 *
	 * Example:
	 *   today := time.Now().Format(stngo.DateFmt)
	 */
	DateFmt = `2006-01-02`
)

var (
	dateLineRE       = regexp.MustCompile("^[0-9][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9]$")
	legacyDateLineRE = regexp.MustCompile("^[0-1][0-9]/[0-3][0-9]/[0-9][0-9][0-9][0-9]$")
	entryLineRE      = regexp.MustCompile("^([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]) - ([0-2][0-9]:[0-6][0-9]|[0-9]:[0-6][0-9]);")
	// A line that starts like a time range, to catch entries with a typo in
	// them (a missing or wrong separator, an extra space).
	startsLikeEntryRE = regexp.MustCompile(`^[0-9]{1,2}:[0-9]{2}\s*-`)
	clockRE           = regexp.MustCompile(`^([0-9]{1,2}):([0-9]{2})$`)
)

/** Entry is the basic data element generated when parsing a file containing
 * Simple Timesheet Notation. It is designed to easily turn into JSON, CSV
 * or other useful formats. Start and End are in the local time zone of the
 * machine that parsed the timesheet, which is the zone the timesheet was
 * written in.
 *
 * Example:
 *   entry, err := stngo.ParseEntry("2026-10-16", "8:00 - 9:00; Acme; Widgets; notes")
 *   fmt.Println(entry.End.Sub(entry.Start)) // 1h0m0s
 */
type Entry struct {
	Start       time.Time
	End         time.Time
	Annotations []string // cells of contextual data (e.g. project, activity, notes)
}

/** IsDateLine validates a line as appropriate to pass to ParseDateLine.
 *
 * Parameters:
 *   line (string) — a line of a timesheet
 *
 * Returns:
 *   bool — true for YYYY-MM-DD and MM/DD/YYYY
 *
 * Example:
 *   stngo.IsDateLine("2026-10-16") // true
 */
func IsDateLine(line string) bool {
	if dateLineRE.MatchString(strings.TrimSpace(line)) {
		return true
	}
	if legacyDateLineRE.MatchString(strings.TrimSpace(line)) {
		return true
	}
	return false
}

/** ParseDateLine sets the current date context when parsing Simple Timesheet
 * Notation elements. It is what is recorded in the Occurrence field of an Entry.
 *
 * Parameters:
 *   line (string) — a date line
 *
 * Returns:
 *   string — the date as YYYY-MM-DD, or "" if the line is not a date line
 *
 * Example:
 *   stngo.ParseDateLine("10/16/2026") // "2026-10-16"
 */
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

/** IsEntry validates a line as an "Entry" to be parsed.
 *
 * Parameters:
 *   line (string) — a line of a timesheet
 *
 * Returns:
 *   bool — true if the line starts with a time range and a semicolon
 *
 * Example:
 *   stngo.IsEntry("8:00 - 9:00; Acme; Widgets") // true
 */
func IsEntry(line string) bool {
	if entryLineRE.MatchString(strings.TrimSpace(line)) {
		return true
	}
	return false
}

/** LooksLikeEntry says a line starts like an entry (a time, a dash) but is
 * not one, which is nearly always a typo: a missing semicolon after the end
 * time, a colon for a semicolon, or an extra space. Such a line is skipped by
 * the parser, so its hours would silently not be counted.
 *
 * Parameters:
 *   line (string) — a line of a timesheet
 *
 * Returns:
 *   bool — true if the line looks like an entry but IsEntry is false
 *
 * Example:
 *   stngo.LooksLikeEntry("8:30 - 9:00 Acme; SysAdmin") // true, no semicolon after 9:00
 */
func LooksLikeEntry(line string) bool {
	line = strings.TrimSpace(line)
	return !IsEntry(line) && startsLikeEntryRE.MatchString(line)
}

func splitCells(line string) []string {
	return strings.Split(line, ";")
}

func splitRangeElements(timeRange string) (string, string, error) {
	if strings.Contains(timeRange, " - ") {
		parts := strings.SplitN(timeRange, " - ", 2)
		return parts[0], parts[1], nil
	}
	return "", "", errors.New("[" + timeRange + "] is not a valid time range string. ")
}

// parseClock reads H:MM or HH:MM into hours and minutes. 24:00 is allowed, and
// means midnight at the end of the day.
func parseClock(value string) (int, int, error) {
	m := clockRE.FindStringSubmatch(strings.TrimSpace(value))
	if m == nil {
		return 0, 0, fmt.Errorf("[%s] is not a time", value)
	}
	hour, _ := strconv.Atoi(m[1])
	minute, _ := strconv.Atoi(m[2])
	if minute > 59 || hour > 24 || (hour == 24 && minute != 0) {
		return 0, 0, fmt.Errorf("parsing time %q: hour or minute out of range", strings.TrimSpace(value))
	}
	return hour, minute, nil
}

// parseRangeElements turns the two clocks of an entry into times on the given
// day in the local time zone. The clocks are wall clock times, so a day on
// which the clocks change still has 8:00 at 8:00 and the offset is the one in
// force at that moment. time.Date normalizes 24:00 to midnight of the next day.
//
// A start time of 1 to 11 o'clock is ambiguous: 1:00 is either 01:00 or 13:00.
// The entries of a day are in order, so their start times ascend. If the time
// as written would start before the entry before it (previous, which may be
// nil) and the afternoon reading would not, the afternoon is meant. With
// nothing before it the time is taken as written. The end time is then the
// first reading that is not before the start, so 8:00 - 4:00 is 8:00 to 16:00.
func parseRangeElements(day time.Time, start string, end string, previous *Entry) (time.Time, time.Time, error) {
	sh, sm, err := parseClock(start)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	eh, em, err := parseClock(end)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	y, mo, d := day.Date()
	startTime := time.Date(y, mo, d, sh, sm, 0, 0, time.Local)
	endTime := rangeEnd(y, mo, d, eh, em, startTime)
	if previous != nil && sh >= 1 && sh < 12 && onDay(previous.Start, day) && startTime.Before(previous.Start) {
		afternoon := time.Date(y, mo, d, sh+12, sm, 0, 0, time.Local)
		// Only if the entry still makes sense: an entry written out of order
		// (9:00 - 3:00 after one that started at 9:15) must not become
		// 21:00 - 3:00.
		if afternoonEnd := rangeEnd(y, mo, d, eh, em, afternoon); !afternoon.Before(previous.Start) && !afternoonEnd.Before(afternoon) {
			startTime, endTime = afternoon, afternoonEnd
		}
	}
	return startTime, endTime, nil
}

// rangeEnd is the end of an entry on a day: the end clock as written, or
// twelve hours later if that is before the start, so 8:00 - 1:00 ends at 13:00.
func rangeEnd(y int, mo time.Month, d int, eh int, em int, start time.Time) time.Time {
	end := time.Date(y, mo, d, eh, em, 0, 0, time.Local)
	//NOTE: need to handle the case where someone has entered an end time
	// smaller than start (e.g. 8:00 - 1:00 meaning 1pm should become 13:00)
	if start.After(end) {
		end = time.Date(y, mo, d, eh+12, em, 0, 0, time.Local)
	}
	return end
}

// onDay says t falls on the same calendar day as day, in the local zone.
func onDay(t time.Time, day time.Time) bool {
	y, m, d := day.Date()
	return !t.Before(time.Date(y, m, d, 0, 0, 0, 0, time.Local)) &&
		t.Before(time.Date(y, m, d+1, 0, 0, 0, 0, time.Local))
}

/** ParseEntry takes the active date and a line and returns an Entry. The
 * times are read as local time on that date, so winter and summer entries
 * each get the offset of their own day. With no entry before it to go by, a
 * start time such as 1:00 is taken as written; see ParseEntryAfter.
 *
 * Parameters:
 *   activeDate (string) — the date of the entry, YYYY-MM-DD or MM/DD/YYYY
 *   line (string) — an entry line, e.g. "8:00 - 9:00; Acme; Widgets; notes"
 *
 * Returns:
 *   *Entry — the parsed entry, nil on error
 *   error — why the line could not be parsed
 *
 * Example:
 *   entry, err := stngo.ParseEntry("2026-10-16", "23:30 - 24:00; Acme; Shutdown")
 */
func ParseEntry(activeDate string, line string) (*Entry, error) {
	return ParseEntryAfter(activeDate, line, nil)
}

/** ParseEntryAfter is ParseEntry for an entry that follows another on the same
 * date. The times of a timesheet are written without am or pm, and the entries
 * of a day are in order, so their start times ascend, and a start time of 1 to
 * 11 o'clock is read as the afternoon when the morning reading would start
 * before the previous entry did. Written after an entry that started at 8:00,
 * "1:00 - 4:00" is 13:00 to 16:00; after one that started at 10:00,
 * "10:00 - 11:00" is an overlap and stays in the morning. A start time of 12:00 or later is never moved, and an entry from
 * another day is no guide.
 *
 * Parameters:
 *   activeDate (string) — the date of the entry, YYYY-MM-DD or MM/DD/YYYY
 *   line (string) — an entry line
 *   previous (*Entry) — the entry before this one in the same day, or nil
 *
 * Returns:
 *   *Entry — the parsed entry, nil on error
 *   error — why the line could not be parsed
 *
 * Example:
 *   first, _ := stngo.ParseEntry("2026-10-16", "8:00 - 8:30; Acme; Widgets")
 *   next, _ := stngo.ParseEntryAfter("2026-10-16", "1:00 - 4:00; Acme; Widgets", first)
 *   fmt.Println(next.Start.Format("15:04")) // 13:00
 */
func ParseEntryAfter(activeDate string, line string, previous *Entry) (*Entry, error) {
	if !IsDateLine(activeDate) {
		return nil, errors.New("invalid format for active date")
	}
	if !IsEntry(line) {
		return nil, errors.New("invalid format for entry")
	}
	day, err := time.ParseInLocation(DateFmt, ParseDateLine(activeDate), time.Local)
	if err != nil {
		return nil, err
	}
	cells := splitCells(line)
	if len(cells) < 2 {
		return nil, errors.New("entry line missing cells")
	}

	s, e, err := splitRangeElements(cells[0])
	if err != nil {
		return nil, err
	}

	start, end, err := parseRangeElements(day, s, e, previous)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(cells); i++ {
		cells[i] = strings.TrimSpace(cells[i])
	}

	return &Entry{
		Start:       start,
		End:         end,
		Annotations: cells[1:],
	}, nil
}

/** JSON converts an Entry struct to JSON notation.
 *
 * Returns:
 *   string — the entry as a JSON object
 *
 * Example:
 *   fmt.Println(entry.JSON())
 */
func (e *Entry) JSON() string {
	src, _ := json.Marshal(e)
	return string(src)
}

/** String converts an Entry struct to a tab delimited string.
 *
 * Returns:
 *   string — start, end and the annotations separated by tabs
 *
 * Example:
 *   fmt.Println(entry.String())
 */
func (e *Entry) String() string {
	return e.Start.Format(time.RFC3339) + "\t" + e.End.Format(time.RFC3339) +
		"\t" + strings.Join(e.Annotations[:], "\t")
}

/** FromString reads a tab delimited string formatted by String into an Entry.
 * The line ending is not part of the last annotation.
 *
 * Parameters:
 *   line (string) — a line written by String
 *
 * Returns:
 *   bool — true if the line was read
 *
 * Example:
 *   entry := new(stngo.Entry)
 *   ok := entry.FromString("2026-10-16T08:00:00-07:00\t2026-10-16T09:00:00-07:00\tAcme\tWidgets")
 */
func (e *Entry) FromString(line string) bool {
	var err error
	// The line comes from a reader and still has its newline; left on, the
	// last cell of an entry with no notes carries it into the report key.
	parts := strings.Split(strings.TrimRight(line, "\r\n"), "\t")
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

/** IsInRange checks the start time of an Entry to see if it is in the range.
 *
 * Parameters:
 *   start (time.Time) — beginning of the range, inclusive
 *   end (time.Time) — end of the range, inclusive
 *
 * Returns:
 *   bool — true if the entry starts inside the range
 *
 * Example:
 *   in := entry.IsInRange(DayStart("2026-10-16"), DayEnd("2026-10-16"))
 */
func (e *Entry) IsInRange(start time.Time, end time.Time) bool {
	t1 := e.Start.Unix()
	if t1 >= start.Unix() && t1 <= end.Unix() {
		return true
	}
	return false
}

/** IsMatch checks the Entry annotations for a matching substring.
 *
 * Parameters:
 *   match (string) — the text to look for
 *
 * Returns:
 *   bool — true if any annotation contains the text
 *
 * Example:
 *   entry.IsMatch("Widgets")
 */
func (e *Entry) IsMatch(match string) bool {
	for i := 0; i < len(e.Annotations); i++ {
		if strings.Contains(e.Annotations[i], match) {
			return true
		}
	}
	return false
}

/** DayStart returns midnight at the start of a date in the local time zone,
 * the zone the timesheet is written in.
 *
 * Parameters:
 *   date (string) — YYYY-MM-DD
 *
 * Returns:
 *   time.Time — 00:00:00 that day
 *   error — if date is not a date
 *
 * Example:
 *   from, err := stngo.DayStart("2026-10-12")
 */
func DayStart(date string) (time.Time, error) {
	return time.ParseInLocation(DateFmt, strings.TrimSpace(date), time.Local)
}

/** DayEnd returns the last moment of a date in the local time zone. It is
 * computed from the start of the next day, so it is right on a day that has
 * 23 or 25 hours.
 *
 * Parameters:
 *   date (string) — YYYY-MM-DD
 *
 * Returns:
 *   time.Time — the last nanosecond of that day
 *   error — if date is not a date
 *
 * Example:
 *   to, err := stngo.DayEnd("2026-10-18")
 */
func DayEnd(date string) (time.Time, error) {
	t, err := DayStart(date)
	if err != nil {
		return t, err
	}
	y, m, d := t.Date()
	return time.Date(y, m, d+1, 0, 0, 0, 0, time.Local).Add(-time.Nanosecond), nil
}
