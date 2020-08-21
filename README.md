# RSA.go

A reimplementation of [RSA.js](https://github.com/stellirin/RSA.js) in Go.

[![codecov](https://codecov.io/gh/stellirin/RSA.go/branch/master/graph/badge.svg)](https://codecov.io/gh/stellirin/RSA.go)

## But why

This is a learning tool, Go has its own `crypto/rsa` package that is probably far superior.

**Don't use this.**

## Code coverage

Go has built-in tools for benchmarking, testing, and visualizing test code coverage.

The GitHub Actions workflow uses this same tooling for CodeCov.

```sh
go test ./... -bench=. -benchmem
go test ./... -covermode=count -coverprofile=coverage.out && go tool cover -html=coverage.out
```

## A note about pointers

Notice in the Git history that I changed `BigInt` operations to use pointers, then reverted it. Using pointers was actually **slower** than copying the values. **[Explanation](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/):**

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

| Function                           | Pointer    | Value      | Difference |
|:---------------------------------- | ----------:| ----------:| ----------:|
| Benchmark_BiCopy-8                 | 59.0 ns/op | 31.3 ns/op | **53%**    |
| Benchmark_BiFromHex-8              |  172 ns/op |  141 ns/op | **82%**    |
| Benchmark_BiAdd-8                  | 82.9 ns/op | 47.3 ns/op | **57%**    |
| Benchmark_BiSubtract-8             | 87.8 ns/op | 50.0 ns/op | **57%**    |
| Benchmark_BiMultiply-8             |  209 ns/op |  162 ns/op | **77%**    |
| Benchmark_BiMultiplyDigit-8        | 78.6 ns/op | 48.1 ns/op | **61%**    |
| Benchmark_BiShiftLeft-8            |  113 ns/op | 85.4 ns/op | **75%**    |
| Benchmark_BiShiftRight-8           |  115 ns/op | 84.3 ns/op | **73%**    |
| Benchmark_BiMultiplyByRadixPower-8 | 67.6 ns/op | 41.4 ns/op | **61%**    |
| Benchmark_BiDivideByRadixPower-8   | 67.6 ns/op | 39.9 ns/op | **59%**    |
| Benchmark_BiModuloByRadixPower-8   | 65.9 ns/op | 38.5 ns/op | **58%**    |
| Benchmark_BiDivideModulo-8         |  543 ns/op |  372 ns/op | **68%**    |
| Benchmark_BiDivide-8               |  585 ns/op |  372 ns/op | **63%**    |
| Benchmark_BiModulo-8               |  578 ns/op |  372 ns/op | **64%**    |

Every operation that involves creating or even only reading a `BigInt` is significantly slower when using a pointer.
