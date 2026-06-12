# errgroup

Structured concurrency: first error cancels the group (`golang.org/x/sync/errgroup`).

## Key concepts

- **`errgroup.WithContext`** — returns a `*Group` and a derived context; canceled when any task errors
- **`g.Go(f)`** — start a managed goroutine
- **`g.Wait()`** — block until all finish; returns the first non-nil error
- **`g.SetLimit(n)`** — cap concurrently running goroutines
- **`g.TryGo(f)`** — non-blocking submit; returns false if at the limit

## Todo

- [x] Fan out three simulated HTTP fetches with `g.Go` — the most common errgroup pattern.
- [x] Collect all results from the fetches and return them after all finish.
- [ ] Have one task return an error early; verify the context is canceled and other tasks observe `ctx.Err()` and exit — practice the cancellation contract.
- [ ] Use `g.SetLimit(4)` to process 20 items; add a counter to confirm no more than 4 run simultaneously.
- [ ] Use `g.TryGo` to submit work non-blocking; handle `false` returns by queueing rejected tasks for retry.

## Run

```bash
go get golang.org/x/sync
go run ./errgroup/
```
