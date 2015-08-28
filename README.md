# str [![GoDoc](https://godoc.org/github.com/matryer/str?status.svg)](https://godoc.org/github.com/matryer/str)
String parsing package for Go. Converts strings to best guess value type.

  * See [parsers.go](https://github.com/matryer/str/blob/master/parsers/parsers.go) for a complete list of parsers.

## Usage

The simplest usage:

```
val := str.Parse("123")
```

To add additional `ParseFunc` functions:

```
parser := append(str.DefaultParser, toMyRangeType)
val := parser.Parse("[1,3]")
```

To use your own parsers entirely, use `New`:

```
parser := str.New(toMyRangeType)
val := parser.Parse("[1,3]")
```
