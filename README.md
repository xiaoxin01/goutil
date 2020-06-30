[![Go Report Card](https://goreportcard.com/badge/github.com/xiaoxin01/goutil)](https://goreportcard.com/report/github.com/xiaoxin01/goutil) [![Build Status](https://travis-ci.org/xiaoxin01/goutil.svg?branch=master)](https://travis-ci.org/xiaoxin01/goutil)

# a util package for go

## strutil

### substring

SubBytesWithRuneCheck/SubStringWithByteLimit get sub string very fast for limit the size of bytes or string when it's very large.

```golang
s := `123456789壹贰叁肆伍陆柒捌玖éèǔ`
SubStringWithByteLimit(s, 9) // 123456789
SubStringWithByteLimit(s, 10) // 123456789
SubStringWithByteLimit(s, 12) // 123456789壹
```

### Benchmark:


|                                   |           |            |         |             |
| --------------------------------- | --------- | ---------- | ------- | ----------- |
| BenchmarkSubBytesWithRuneCheck-8  | 143022060 | 8.38 ns/op | 0 B/op  | 0 allocs/op |
| BenchmarkSubStringWithByteLimit-8 | 24984644  | 44.9 ns/op | 48 B/op | 1 allocs/op |
