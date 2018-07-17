## Concat

### Overview

This package provides simple functions that return concatenation results.
Can work faster than Go `+` operator.

You should not use this package, really. It's just an example.

### Benchmarks

```
BenchmarkConcat2Operator/short-8         	20000000	        84.4 ns/op
BenchmarkConcat2Operator/longer-8        	10000000	       158 ns/op
BenchmarkConcat2Builder/short-8          	20000000	        70.7 ns/op
BenchmarkConcat2Builder/longer-8         	10000000	       127 ns/op
BenchmarkConcat2/short-8                 	30000000	        57.3 ns/op
BenchmarkConcat2/longer-8                	20000000	       106 ns/op
BenchmarkConcat3Operator/short-8         	20000000	       103 ns/op
BenchmarkConcat3Operator/longer-8        	10000000	       217 ns/op
BenchmarkConcat3Builder/short-8          	20000000	        89.9 ns/op
BenchmarkConcat3Builder/longer-8         	 5000000	       249 ns/op
BenchmarkConcat3/short-8                 	20000000	        85.0 ns/op
BenchmarkConcat3/longer-8                	10000000	       189 ns/op
```

Number one is unsafe concatenation, second is `strings.Builder` with preallocated
buffer and "obvious" concatenation is the slowest one... unless [CL123256](https://go-review.googlesource.com/c/go/+/123256) is applied.

Using the `benchstat`, here is the difference between `concat` and `+`:

```
name              old time/op  new time/op  delta
Concat2/short-8   84.4ns ± 2%  64.3ns ± 4%  -23.85%  (p=0.000 n=14+15)
Concat2/longer-8   138ns ± 1%   118ns ± 1%  -14.83%  (p=0.000 n=13+15)
Concat3/short-8    105ns ± 5%    82ns ± 5%  -22.29%  (p=0.000 n=15+14)
Concat3/longer-8   218ns ± 1%   192ns ± 1%  -11.95%  (p=0.000 n=15+15)
```

If compared with AMD64 asm version for concat2:

```
name              old time/op  new time/op  delta
Concat2/short-8   84.4ns ± 0%  56.9ns ± 5%  -32.54%  (p=0.000 n=15+15)
Concat2/longer-8   138ns ± 1%   107ns ± 0%  -22.51%  (p=0.000 n=13+15)
```

As a bonus, asm version also makes empty strings concatenation optimization,
just like runtime version of concat would.

### Example

```go
package main

import (
	"fmt"

	"github.com/Quasilyte/concat"
)

func main() {
	v := "world!"
	fmt.Println(concat.Strings("hello, ", v)) // => "hello, world!"
}
```
