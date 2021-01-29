# Intpacker

[![Go Reference](https://pkg.go.dev/badge/github.com/ograycode/intpacker.svg)](https://pkg.go.dev/github.com/ograycode/intpacker)

*Unfortunately I don't believe this is possible at the moment because doing addition/subtraction has the potential to overflow into the next number.*

-------

Just have a few ints to track, but don't want the overhead of channels or mutexex? Neither did I. So I decided to pack many numbers into a single number.

This is experimental.

## Performance

Early indications are that it is good for avoiding mutexes and their overhead by being able to use the `sync/atomic` package on multiple numbers at a time.

Tested with go1.15.2 windows/amd64 on a i7-8750H CPU @ 2.20GHz, 2208 Mhz, 6 Cores, 12 Logical Processors

```
go test -benchmem -bench .                                                                                        
goos: windows
goarch: amd64
pkg: github.com/ograycode/intpacker
BenchmarkUint32AtomicAdd-12             118378352               10.1 ns/op             0 B/op          0 allocs/op
BenchmarkUint32WithMutex-12             63579248                17.1 ns/op             0 B/op          0 allocs/op
BenchmarkUint32AtomicAddParallel-12     45477213                24.6 ns/op             0 B/op          0 allocs/op
BenchmarkUint32WithMutexParallel-12     16742941                72.4 ns/op             0 B/op          0 allocs/op
PASS
```

## TODO

- More tests.
- More benchmarks.
- How to pack and unpack more types, like signed ints.
- How to fully support negative numbers.
