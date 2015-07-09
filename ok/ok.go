/**
 * ok.go - A small library to make assertion like testing statements friendly to the standard testing
 * package available with Go.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the Simplified BSD License.
 * See:
 */

// Package ok provides a functions similar to those in NodeJS's assert module without catastrophic side effects.
package ok

import (
	"testing"
)

func Ok(t *testing.T, expression bool, message string) bool {
	if expression == true {
		return true
	}
	t.Errorf("Failed (expected true): [%s]\n", message)
	return false
}

func NotOk(t *testing.T, expression bool, message string) bool {
	if expression == false {
		return true
	}
	t.Errorf("Failed (expected false): [%s]\n", message)
	return false
}
