package parsers

import (
	"strconv"
	"strings"
	"time"
)

// Time gets a str.ParseFunc that converts a string into
// a time.Time with the given layout.
func Time(layout string) func(string) (interface{}, bool) {
	return func(s string) (interface{}, bool) {
		t, err := time.Parse(layout, s)
		if err != nil {
			return nil, false
		}
		return t, true
	}
}

// Quoted is a str.ParseFunc that removes the quotes from the
// specified quoted string.
func Quoted(s string) (interface{}, bool) {
	const quotes = "`'\""
	l := len(s)
	if l < 2 {
		return nil, false
	}
	if s[0] != s[l-1] {
		return nil, false
	}
	if strings.IndexAny(s, quotes) == 0 {
		return s[1 : l-1], true
	}
	return nil, false
}

// Nil is a str.ParseFunc that returns nil.
func Nil(s string) (interface{}, bool) {
	return nil, len(s) == 0
}

// Null is a str.ParseFunc that returns nil if the
// string is "null".
func Null(s string) (interface{}, bool) {
	return nil, strings.ToLower(s) == "null"
}

// Bool is a str.ParseFunc that converts a string
// into a bool.
func Bool(s string) (interface{}, bool) {
	switch strings.ToLower(s) {
	case "true":
		return true, true
	case "false":
		return false, true
	}
	return nil, false
}

// Int is a str.ParseFunc that converts a string into
// an int.
func Int(s string) (interface{}, bool) {
	var err error
	var val int64
	if val, err = strconv.ParseInt(s, 10, 0); err != nil {
		return nil, false
	}
	return int(val), true
}

// Int64 is a str.ParseFunc that converts a string into
// an int64.
func Int64(s string) (interface{}, bool) {
	var err error
	var val int64
	if val, err = strconv.ParseInt(s, 10, 64); err != nil {
		return nil, false
	}
	return int64(val), true
}

// UInt is a str.ParseFunc that converts a string into
// a uint.
func UInt(s string) (interface{}, bool) {
	var err error
	var val uint64
	if val, err = strconv.ParseUint(s, 10, 0); err != nil {
		return nil, false
	}
	return uint(val), true
}

// Uint64 is a str.ParseFunc that converts a string into
// a uint64.
func Uint64(s string) (interface{}, bool) {
	var err error
	var val uint64
	if val, err = strconv.ParseUint(s, 10, 64); err != nil {
		return nil, false
	}
	return uint64(val), true
}

// Float64 is a str.ParseFunc that converts a string into
// a float64.
func Float64(s string) (interface{}, bool) {
	var err error
	var val float64
	if val, err = strconv.ParseFloat(s, 64); err != nil {
		return nil, false
	}
	return float64(val), true
}
