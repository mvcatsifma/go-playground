# synctest

`testing/synctest` — deterministic concurrency testing with a fake clock (Go 1.24+).

## What's here

`synctest_test.go` — demonstrates `synctest.Wait`, fake clock, timeout testing, and goroutine coordination.

## Key concepts

- **`synctest.Test(t, f)`** — runs `f` in an isolated bubble with a fake clock starting at 2000-01-01
- **`synctest.Wait()`** — blocks until every other goroutine in the bubble is durably blocked
- **Fake time** — `time.Sleep` advances the fake clock instantly once all goroutines are idle; tests that would take seconds run in microseconds
- **Durably blocked** — channel ops, `time.Sleep`, `sync.Cond.Wait`, `sync.WaitGroup.Wait`; NOT mutex locks or I/O

## Todo

- [x] Start a goroutine that sets a flag; call `synctest.Wait()` and assert the flag is set — verify you don't need `time.Sleep` to synchronise.
- [x] Test a `context.WithTimeout` without real waiting: sleep to just before the deadline, assert no error; sleep past it, assert `DeadlineExceeded`.
- [x] Test `context.AfterFunc`: verify the callback is not called before cancel and is called after — the example from the official docs, written from scratch.
- [x] Write a ticker-based rate limiter and test N ticks fire in order using fake time; confirm the test completes in microseconds.
- [ ] Test a retry-with-backoff function: advance fake time through multiple retry intervals and assert the correct number of attempts without real waiting.

## Run

```bash
go test ./synctest/
```
