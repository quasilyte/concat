## Concat

This package provides simple functions that return concatenation results.
Can work faster than Go "+" operator.

You should not use this package, really. It's just an example.

Benchmark results:

```
BenchmarkConcat2Operator-8   	20000000	        83.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkConcat2Builder-8    	20000000	        70.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkConcat2-8           	20000000	        62.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkConcat3Operator-8   	20000000	       104 ns/op	      32 B/op	       1 allocs/op
BenchmarkConcat3Builder-8    	20000000	        89.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkConcat3-8           	20000000	        82.1 ns/op	      32 B/op	       1 allocs/op
```

Number one is unsafe concatenation, second is `strings.Builder` with preallocated
buffer and "obvious" concatenation is the slowest one... unless [CL123256](https://go-review.googlesource.com/c/go/+/123256) is applied.
