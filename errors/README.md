# errors

Sentinel errors, wrapping, `errors.Is` / `errors.As`.

## What's here

`errors_test.go` — sentinel, `%w` wrapping, custom type with `Unwrap`, `errors.Is`, `errors.As`.

## Key concepts

**Pointer to interface in errors.As:**

`errors.As` requires a pointer to an interface variable so it can modify what concrete value the interface holds:

```go
var tempErr temporaryError           // interface variable
errors.As(err, &tempErr)              // pass pointer to interface variable
```

The relationship:
- `temporaryError` — the interface (defines the contract)
- `*NetworkError` — a concrete type (implements the contract)
- `tempErr` — an interface variable (can hold any concrete type implementing temporaryError)

errors.As searches the error chain for any concrete error implementing the temporaryError interface. When found (e.g., `*NetworkError`), it assigns that concrete value to your interface variable through the pointer.

## Todo

- [x] Add a multi-level chain: `fmt.Errorf` wrapping `fmt.Errorf` wrapping `XError`; verify `errors.Is` still finds the sentinel at the bottom.
- [x] Use `errors.Join` (Go 1.20) to combine two errors; write a helper `HasError(err, target error) bool` and verify both joined errors are reachable.
- [x] Write a `Retry(n int, f func() error) error` function that retries until the error is not `Temporary()`; use `errors.As` to check.
- [x] Replace `NError.Unwrap() error` with `Unwrap() []error`; verify `errors.Is` traverses the slice — this is the pattern `errors.Join` uses internally.
