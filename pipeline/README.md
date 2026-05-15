# pipeline

Channel pipelines: fan-out, fan-in, cancellation.

## Key concepts

- **Stage** — goroutine reading from `in <-chan T`, writing to returned `out <-chan T`; closes `out` when `in` closes
- **Fan-out** — one channel feeds multiple stages in parallel
- **Fan-in (merge)** — multiple channels merged into one with a `sync.WaitGroup`
- **Cancellation** — `context.Context` passed through every stage so all goroutines stop cleanly

## Todo

- [ ] Build a three-stage pipeline `generate → square → print`; verify values flow through in order.
- [ ] Fan out to two `square` workers and merge with a `merge` function; observe that output order is no longer guaranteed.
- [ ] Add `ctx context.Context` to every stage; cancel after 3 items and verify no goroutines leak (use `-race` and check with `goleak`).
- [ ] Add error propagation: have one worker fail and surface the error to the caller while other stages shut down gracefully.

## Run

```bash
go run ./pipeline/
```
