package str_test

import (
	"fmt"
	"testing"

	"github.com/cheekybits/is"
	"github.com/matryer/str"
)

var tests = []struct {
	s string
	v interface{}
}{

	// strings
	{
		"This is clearly a string",
		string("This is clearly a string"),
	},

	// nils
	{
		"",
		nil,
	},
	{
		"null",
		nil,
	},

	// bools
	{
		"true",
		true,
	},
	{
		"false",
		false,
	},
	{
		"True",
		true,
	},
	{
		"False",
		false,
	},

	// int
	{
		"0",
		int(0),
	},
	{
		fmt.Sprintf("%d", ^int(0)),
		^int(0),
	},

	// uint
	{
		fmt.Sprintf("%d", ^uint(0)),
		^uint(0),
	},

	// float64
	{
		"1.1",
		1.1,
	},

	// forced strings (quoted values will always be string)
	{
		`"true"`,
		"true",
	},
	{
		`'true'`,
		"true",
	},
	{
		"`true`",
		"true",
	},
	{
		`""`,
		"",
	},
	{
		`''`,
		"",
	},
	{
		"``",
		"",
	},
}

func TestParsers(t *testing.T) {
	is := is.New(t)
	for _, test := range tests {
		actual := str.Parse(test.s)
		// assert type and value
		is.Equal(fmt.Sprintf("%T", actual), fmt.Sprintf("%T", test.v))
		is.Equal(actual, test.v)
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
