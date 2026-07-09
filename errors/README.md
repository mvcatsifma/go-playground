# errors

Sentinel errors, wrapping, `errors.Is` / `errors.As`.

## What's here

`errors_test.go` — sentinel, `%w` wrapping, custom type with `Unwrap`, `errors.Is`, `errors.As`.

## Todo

- [x] Add a multi-level chain: `fmt.Errorf` wrapping `fmt.Errorf` wrapping `XError`; verify `errors.Is` still finds the sentinel at the bottom.
- [x] Use `errors.Join` (Go 1.20) to combine two errors; write a helper `HasError(err, target error) bool` and verify both joined errors are reachable.
- [x] Write a `Retry(n int, f func() error) error` function that retries until the error is not `Temporary()`; use `errors.As` to check.
- [x] Replace `NError.Unwrap() error` with `Unwrap() []error`; verify `errors.Is` traverses the slice — this is the pattern `errors.Join` uses internally.
