# RSA.go

A reimplementation of [RSA.js](https://github.com/stellirin/RSA.js) in Go.

## But why

This is a learning tool, Go has its own `crypto/rsa` package that is probably far superior.

**Don't use this.**

## Code coverage

Go has built-in tools fot benchmarking, testing, and visualizing test code coverage.

```sh
go test -bench=. -benchmem
go test -covermode=count -coverprofile=coverage.out && go tool cover -html=coverage.out
```

## A note about pointers

Notice in the Git history that I changed `bigInt` operations to use pointers, then reverted it. Using pointers was actually **slower** than copying the values. **[Explanation](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/):**

> A common hypothesis derived from intuition goes something like this: _“copying values is expensive, so instead I’ll use a pointer.”_ However, in many cases copying a value is much less expensive than the overhead of using a pointer. _“Why”_ you might ask?
>
> - **The compiler generates checks when dereferencing a pointer.** The purpose is to avoid memory corruption by running panic() if the pointer is nil. This is extra code that must be executed at runtime. When data is passed by value, it cannot be nil.
>
> - **Pointers often have poor locality of reference.** All of the values used within a function are collocated in memory on the stack. Locality of reference is an important aspect of efficient code. It dramatically increases the chance that a value is warm in CPU caches and reduces the risk of a miss penalty during prefetching.
>
> - **Copying objects within a cache line is the roughly equivalent to copying a single pointer.** CPUs move memory between caching layers and main memory on cache lines of constant size. On x86 this is 64 bytes. Further, Go uses a technique called Duff’s device to make common memory operations like copies very efficient.
>
> Pointers should primarily be used to reflect ownership semantics and mutability. In practice, the use of pointers to avoid copies should be infrequent. Don’t fall into the trap of premature optimization. It’s good to develop a habit of passing data by value, only falling back to passing pointers when necessary. An extra bonus is the increased safety of eliminating nil.

## Benchmarks

I confirmed the above with Go benchmarks:

| Function                           | Passing a | Pointer    | Passing a | Value      | Difference |
|:---------------------------------- | ---------:|:---------- | ---------:| ---------- | ---------- |
| Benchmark_biCopy-8                 |  18319030 | 59.0 ns/op |  38715211 | 31.3 ns/op | **53%**    |
| Benchmark_biToHex-8                |    463184 | 2236 ns/op |    462002 | 2220 ns/op | ~~99%~~    |
| Benchmark_biFromHex-8              |   6530492 |  172 ns/op |   8208451 |  141 ns/op | **82%**    |
| Benchmark_biToBytes-8              |   2211871 |  545 ns/op |   2133015 |  529 ns/op | ~~97%~~    |
| Benchmark_biAdd-8                  |  13779902 | 82.9 ns/op |  24048735 | 47.3 ns/op | **57%**    |
| Benchmark_biSubtract-8             |  12815166 | 87.8 ns/op |  22582344 | 50.0 ns/op | **57%**    |
| Benchmark_biHighIndex-8            | 625497319 | 1.87 ns/op | 501739746 | 2.33 ns/op | ~~124%~~   |
| Benchmark_biNumBits-8              | 320694976 | 3.60 ns/op | 279414602 | 4.29 ns/op | ~~119%~~   |
| Benchmark_biMultiply-8             |   5727783 |  209 ns/op |   7060063 |  162 ns/op | **77%**    |
| Benchmark_biMultiplyDigit-8        |  14251012 | 78.6 ns/op |  23229156 | 48.1 ns/op | **61%**    |
| Benchmark_biShiftLeft-8            |  10511432 |  113 ns/op |  13595034 | 85.4 ns/op | **75%**    |
| Benchmark_biShiftRight-8           |   9717110 |  115 ns/op |  13426609 | 84.3 ns/op | **73%**    |
| Benchmark_biMultiplyByRadixPower-8 |  15846832 | 67.6 ns/op |  26826535 | 41.4 ns/op | **61%**    |
| Benchmark_biDivideByRadixPower-8   |  16887123 | 67.6 ns/op |  29196828 | 39.9 ns/op | **59%**    |
| Benchmark_biModuloByRadixPower-8   |  17310244 | 65.9 ns/op |  30306049 | 38.5 ns/op | **58%**    |
| Benchmark_biCompare-8              | 152232939 | 7.75 ns/op | 153687193 | 7.57 ns/op | ~~97%~~    |
| Benchmark_biDivideModulo-8         |   2103934 |  543 ns/op |   3216999 |  372 ns/op | **68%**    |
| Benchmark_biDivide-8               |   2024931 |  585 ns/op |   3150711 |  372 ns/op | **63%**    |
| Benchmark_biModulo-8               |   2050354 |  578 ns/op |   3191080 |  372 ns/op | **64%**    |

Every operation that involves creating or even only reading a `bigInt` is significantly slower when using a pointer.

## Links

- [golang-book.com](https://www.golang-book.com/)
