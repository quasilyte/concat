## Concat

This package provides simple functions that return concatenation results.
Can work faster than Go `+` operator.

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

Using the `benchstat`, here is the difference between `concat` and `+`:

```
name       old time/op  new time/op  delta
Concat2-8  84.2ns ± 1%  62.7ns ± 2%  -25.49%  (p=0.000 n=9+10)
Concat3-8   103ns ± 3%    83ns ± 4%  -19.83%  (p=0.000 n=10+9)
```

If compared with AMD64 asm version for concat2:

```
name       old time/op  new time/op  delta
Concat2-8  84.2ns ± 1%  57.1ns ± 3%  -32.20%  (p=0.000 n=9+9)
```

As a bonus, asm version also makes empty strings concatenation optimization,
just like runtime version of concat would.
