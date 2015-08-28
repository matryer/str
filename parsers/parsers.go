package parsers

import (
	"strconv"
	"strings"
)

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

func Nil(s string) (interface{}, bool) {
	return nil, len(s) == 0
}

func Null(s string) (interface{}, bool) {
	return nil, strings.ToLower(s) == "null"
}

func Bool(s string) (interface{}, bool) {
	switch strings.ToLower(s) {
	case "true":
		return true, true
	case "false":
		return false, true
	}
	return nil, false
}

func Int(s string) (interface{}, bool) {
	var err error
	var val int64
	if val, err = strconv.ParseInt(s, 10, 0); err != nil {
		return nil, false
	}
	return int(val), true
}

func Int64(s string) (interface{}, bool) {
	var err error
	var val int64
	if val, err = strconv.ParseInt(s, 10, 64); err != nil {
		return nil, false
	}
	return int64(val), true
}

func UInt(s string) (interface{}, bool) {
	var err error
	var val uint64
	if val, err = strconv.ParseUint(s, 10, 0); err != nil {
		return nil, false
	}
	return uint(val), true
}

func Uint64(s string) (interface{}, bool) {
	var err error
	var val uint64
	if val, err = strconv.ParseUint(s, 10, 64); err != nil {
		return nil, false
	}
	return uint64(val), true
}

func Float64(s string) (interface{}, bool) {
	var err error
	var val float64
	if val, err = strconv.ParseFloat(s, 64); err != nil {
		return nil, false
	}
	return float64(val), true
}
