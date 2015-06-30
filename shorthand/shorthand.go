/**
 * shorthand.go - definition and expantion for stngo project.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved.
 * Released under the BSD 2-Clause license
 */

// Package shorthand provides shorthand storage and expantion for stngo project.
package shorthand

import (
	"strings"
)

// Abbrevations holds the shorthand and translation
var Abbreviations = make(map[string]string)

// IsAssignment checks to see if a string contains an assignment (i.e. has a ' := ' in the string.)
func IsAssignment(text string) bool {
	if strings.Index(text, " := ") == -1 {
		return false
	}
	return true
}

// HasAssignment checks to see if a shortcut has already been assigned.
func HasAssignment(key string) bool {
	_, ok := Abbreviations[key]
	return ok
}

// Assign stores a shorthand and its expantion
func Assign(s string) bool {
	parts := strings.SplitN(strings.TrimSpace(s), " := ", 2)
	key, value := parts[0], parts[1]
	if key == "" || value == "" {
		return false
	}
	Abbreviations[key] = value
	_, ok := Abbreviations[key]
	return ok
}

// Expand takes a text and expands all shorthands
func Expand(text string) string {
	// Iterate through the list of key/values in abbreviations
	for key, value := range Abbreviations {
		text = strings.Replace(text, key, value, -1)
	}
	return text
}

// Clear remove all the elements of a map.
func Clear() {
	for key := range Abbreviations {
		delete(Abbreviations, key)
	}
}
