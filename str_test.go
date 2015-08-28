package str_test

import (
	"fmt"
	"testing"

	"github.com/matryer/str/parsers"

	"github.com/cheekybits/is"
	"github.com/matryer/str"
)

var tests = []struct {
	s string        // source string
	v interface{}   // expected value
	f str.ParseFunc // expected ParseFunc
}{

	// strings
	{
		"This is clearly a string",
		string("This is clearly a string"),
		nil,
	},

	// nils
	{
		"",
		nil,
		parsers.Nil,
	},
	{
		"null",
		nil,
		parsers.Null,
	},

	// bools
	{
		"true",
		true,
		parsers.Bool,
	},
	{
		"false",
		false,
		parsers.Bool,
	},
	{
		"True",
		true,
		parsers.Bool,
	},
	{
		"False",
		false,
		parsers.Bool,
	},

	// int
	{
		"0",
		int(0),
		parsers.Int,
	},
	{
		fmt.Sprintf("%d", ^int(0)),
		^int(0),
		parsers.Int,
	},

	// uint
	{
		fmt.Sprintf("%d", ^uint(0)),
		^uint(0),
		parsers.UInt,
	},

	// float64
	{
		"1.1",
		1.1,
		parsers.Float64,
	},

	// forced strings (quoted values will always be string)
	{
		`"true"`,
		"true",
		parsers.Quoted,
	},
	{
		`'true'`,
		"true",
		parsers.Quoted,
	},
	{
		"`true`",
		"true",
		parsers.Quoted,
	},
	{
		`""`,
		"",
		parsers.Quoted,
	},
	{
		`''`,
		"",
		parsers.Quoted,
	},
	{
		"``",
		"",
		parsers.Quoted,
	},
}

func TestParseWith(t *testing.T) {
	is := is.New(t)
	for _, test := range tests {
		actual, actualFunc := str.ParseWith(test.s)
		// assert type and value
		is.Equal(fmt.Sprintf("%T", actual), fmt.Sprintf("%T", test.v))
		is.Equal(actual, test.v)
		is.Equal(fmt.Sprint(actualFunc), fmt.Sprint(test.f))
	}
}

func TestNewParser(t *testing.T) {
	is := is.New(t)
	parser := str.New(func(s string) (interface{}, bool) {
		if s == "1" {
			return 1, true
		}
		return nil, false
	})
	is.Equal(1, parser.Parse("1"))
	is.Equal("2", parser.Parse("2"))
}
