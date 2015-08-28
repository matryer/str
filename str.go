package str

import "github.com/matryer/str/parsers"

// ParseFunc is a function that attempts to parse the
// string into a real type and returns whether that was
// successful or not.
type ParseFunc func(s string) (interface{}, bool)

// Parser is a slice of ParseFunc functions that will be
// called in order.
type Parser []ParseFunc

// New creates a new Parser with the specified ParseFunc
// functions. Each function will be tried in order until
// one is able to parse the strings.
// If none can, the original string is returned untouched.
func New(funcs ...ParseFunc) Parser {
	return Parser(funcs)
}

// Parse parses s with each ParseFunc in order returning
// the result.
// If no parsers are successful, returns the string untouched.
func (p Parser) Parse(s string) interface{} {
	for _, try := range p {
		if val, ok := try(s); ok {
			return val
		}
	}
	return s
}

// Parse parses s with the DefaultParser.
func Parse(s string) interface{} {
	return DefaultParser.Parse(s)
}

// DefaultParser is the default Parser that includes
// all built-in ParseFunc functions.
var DefaultParser = Parser([]ParseFunc{
	parsers.Quoted,
	parsers.Nil,
	parsers.Null,
	parsers.Bool,
	parsers.Int,
	parsers.Int64,
	parsers.UInt,
	parsers.Uint64,
	parsers.Float64,
})
