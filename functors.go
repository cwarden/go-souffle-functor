package main

import "C"

import (
	"strings"
	"time"
)

//export timestampSeconds
func timestampSeconds(t *C.char) C.long {
	s := C.GoString(t)
	p, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return 0
	}
	return C.long(p.Unix())
}

//export toLower
func toLower(s *C.char) *C.char {
	gs := C.GoString(s)
	cs := C.CString(strings.ToLower(gs))
	return cs
}

func main() {}
