# str [![GoDoc](https://godoc.org/github.com/matryer/str?status.svg)](https://godoc.org/github.com/matryer/str)
String parsing package for Go. Converts strings to best guess value type.

Supported types (in this order):

  * `nil` or `null`
  * `bool`
  * `int`
  * `int64`
  * `uint`
  * `uint64`
  * `float64`

Quoted values are always strings with the quotes stripped.

For examples see the [test code](https://github.com/matryer/str/blob/master/str_test.go#L11-L115).

## Get started

```
go get gopkg.in/matryer/str.v1
```

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
