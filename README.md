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
| `ctrl-c/` | OS signal handling and graceful shutdown |
| `embed/` | `//go:embed` directives |
| `errors/` | Sentinel errors, `errors.Is` / `errors.As`, custom error types |
| `fs/` | `io/fs` interfaces (`ReadFileFS`, `StatFS`) with injectable filesystem |
| `generics/` | Type parameters and constraints — basic usage |
| `generics3/` | Generics — intermediate patterns |
| `generics5/` | Generics — generic HTTP response handling |
| `godirwalk2/` | Directory walking with error handling |
| `gophercon/` | Channel fan-out and select-based shutdown |
| `hash/` | `crypto/sha256` hashing |
| `httphandler/` | Function-as-`io.Writer` pattern for HTTP handlers |
| `integer-heap/` | Min-heap over integers via `container/heap` |
| `job-queue/` | Heap-based job queue |
| `log/` | `log` package flags, prefix, and UTC formatting |
| `maxgoroutines/` | Goroutine limiting patterns |
| `options/` | Functional options pattern |
| `priority-queue/` | Priority queue via `container/heap` |
| `stringer/` | `go:generate` with the `stringer` tool |
| `template/` | `html/template` rendering |
| `timeout/` | Context cancellation, semaphore limiting, OS signals, directory walking |
| `timer/` | `time.Timer` behaviour (expired timer edge cases) |
| `unitofwork/` | Unit of Work pattern with `go-cmp` |
| `viper/` | Config management with Viper |
| `worker_pool/` | Worker pool / goroutine limiting |
