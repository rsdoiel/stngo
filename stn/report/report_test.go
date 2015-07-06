/**
 * report_test.go test the stn/report package.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license
 * See: http://opensource.org/licenses/BSD-2-Clause
 */
// Package report provides basic report rendering for Simple Timesheet Notation.
package report

import (
	"log"
	"testing"
)

func ok(t *testing.T, expected bool, msg string) {
	if expected != true {
		t.Fatalf("Failed: " + msg)
	}
}

func TestAggregator(t *testing.T) {
	log.Fatal("TestAggregator() not implemented.")
}

func TestHeadings(t *testing.T) {
	log.Fatal("TestHeadings() not implemented.")
}

func TestFooter(t *testing.T) {
	log.Fatal("TestFooter() not implemented.")
}

func TestBody(t *testing.T) {
	log.Fatal("TestBody() not implemented.")
}

func TestReportsDaily(t *testing.T) {
	log.Fatal("TestReportsDaily() not implemented.")
}

func TestReportsWeekly(t *testing.T) {
	log.Fatal("TestReportsWeekly() not implemented.")
}

func TestReportsMonthly(t *testing.T) {
	log.Fatal("TestReportsMonthly() not implemented.")
}

func TestReportsYearly(t *testing.T) {
	log.Fatal("TestReportsMonthly() not implemented.")
}
