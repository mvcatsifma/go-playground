# bench

Benchmarks with `testing.B`.

## Key concepts

- **`b.N`** — loop count, auto-tuned by the runner
- **`b.ReportAllocs()`** — show heap allocations per op alongside ns/op
- **`b.ResetTimer`** — exclude setup time from the measurement
- **`b.RunParallel`** — run the body across `GOMAXPROCS` goroutines; essential for measuring shared-state contention
- **`-benchmem`** — print alloc stats globally without touching code

## Todo

- [ ] Benchmark `fmt.Sprintf("k=%d", n)` vs `strconv.AppendInt` into a reused `[]byte`; use `b.ReportAllocs()` and observe the alloc difference — learn to read ns/op and allocs/op together.
- [ ] Benchmark a `sync.Mutex`-protected counter vs an `atomic.Int64` counter using `b.RunParallel` — the result will surprise you at high `GOMAXPROCS`.
- [ ] Add expensive setup before the benchmark loop; use `b.ResetTimer` and verify the setup time is excluded.
- [ ] Write two implementations of the same function and use a `BenchmarkXxx/impl=A` / `BenchmarkXxx/impl=B` sub-benchmark structure — practice the pattern used to compare alternatives.

## Run

```bash
go test -bench=. -benchmem ./bench/
```
