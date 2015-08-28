# str
String parsing package for Go. Converts strings to best guess value type.

  * See [parsers.go](https://github.com/matryer/str/blob/master/parsers/parsers.go) for a complete list of parsers.

## Usage

The simplest usage:

```
val := str.Parse("123")
```

To add additional `ParseFunc` functions:

```
parser := append(DefaultParser, toMyRangeType)
val := parser.Parse("[1,3]")
```
