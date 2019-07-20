//
// stn.go - is a shared library to make the stn package available
// as a Python3 module via ctypes.
package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

import (
	"github.com/rsdoiel/stngo"
)

//export Version
func Version() *C.char {
	cs := C.CString(stn.Version)
	defer C.free(unsafe.Pointer(cs))
	return cs
}

func main() {}
