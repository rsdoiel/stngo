//
// libstn.go - is a shared library to make the stn package available
// to other programming languages such as Python3 and Julia.
//
package main

import "C"

import (
	"github.com/rsdoiel/stngo"
)

//export version
func version() *C.char {
	return C.CString(stn.Version)
}

func main() {}
