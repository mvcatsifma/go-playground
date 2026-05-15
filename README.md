# go-playground

A flat collection of self-contained Go experiments and learning exercises. Each top-level directory is an independent program (`package main`) that can be run directly.

**Module:** `dsen.nl/go-playground`

## Running experiments

```bash
# Run a single experiment
go run ./generics/

# Run all tests
go test ./...

# Run tests for a single package
go test ./errors/
```

## Experiments

| Directory | Topic |
|-----------|-------|
| `atomic/` | Typed atomics — `atomic.Int64`, `atomic.Pointer[T]`, CAS |
| `bench/` | `testing.B` benchmarks, `b.ReportAllocs`, `b.RunParallel` |
| `cmp/` | Generic comparison — `cmp.Compare`, `cmp.Or`, `cmp.Ordered` |
| `ctrl-c/` | OS signal handling and graceful shutdown |
| `embed/` | `//go:embed` directives |
| `errgroup/` | `golang.org/x/sync/errgroup` — structured concurrency with error propagation |
| `errors/` | Sentinel errors, `errors.Is` / `errors.As`, custom error types |
| `fs/` | `io/fs` interfaces (`ReadFileFS`, `StatFS`) with injectable filesystem |
| `fuzz/` | `testing.F` fuzz targets, property-based testing |
| `generics/` | Type parameters and constraints — basic usage |
| `generics3/` | Generics — intermediate patterns |
| `generics5/` | Generics — generic HTTP response handling |
| `godirwalk2/` | Directory walking with error handling |
| `gofix/` | `go fix` tool — automated source rewriting for API migrations |
| `gophercon/` | Channel fan-out and select-based shutdown |
| `hash/` | `crypto/sha256` hashing |
| `httphandler/` | Function-as-`io.Writer` pattern for HTTP handlers |
| `httpmux/` | `net/http.ServeMux` enhanced routing — method patterns, wildcards |
| `iface/` | Interface composition, type assertions, type switches |
| `integer-heap/` | Max-heap over integers via `container/heap` |
| `iopipe/` | `io.Pipe` for buffer-free streaming between writer and reader |
| `iter/` | `iter.Seq` / range-over-func, custom iterators, `iter.Pull` |
| `job-queue/` | Heap-based priority queue with O(log n) updates |
| `log/` | `log` package flags, prefix, and UTC formatting |
| `maps/` | Generic map helpers — `maps.Clone`, `maps.DeleteFunc`, `maps.All` |
| `maxgoroutines/` | Goroutine limiting patterns |
| `options/` | Functional options pattern |
| `pipeline/` | Channel pipelines: fan-out, fan-in, cancellation |
| `priority-queue/` | Priority queue via `container/heap` with in-place updates |
| `reverseproxy/` | `httputil.ReverseProxy` with `Rewrite`, `ModifyResponse`, `ErrorHandler` |
| `slices/` | Generic slice helpers — `slices.Sort`, `slices.Compact`, `slices.Collect` |
| `slog/` | Structured logging with `log/slog`, handlers, attrs, groups |
| `stringer/` | `go:generate` with the `stringer` tool |
| `synconce/` | `sync.Once`, `sync.OnceValue`, `sync.Map` |
| `synctest/` | `testing/synctest` — fake clock and goroutine-aware test bubbles |
| `template/` | `html/template` rendering |
| `timeout/` | Context cancellation, semaphore limiting, OS signals, directory walking |
| `timer/` | `time.Timer` behaviour (expired timer edge cases) |
| `unique/` | Value interning with `unique.Make` (Go 1.23+) |
| `unitofwork/` | Unit of Work pattern with `go-cmp` |
| `viper/` | Config management with Viper |
| `worker_pool/` | Semaphore-based goroutine limiting |