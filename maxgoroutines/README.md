# maxgoroutines

Worker pool: one goroutine per CPU, channel-based work distribution, WaitGroup shutdown.

## What's here

`main.go` — spawns `NumCPU()` workers draining `itemChan`; SIGINT stops the producer and closes the channel.

## Todo

- [ ] Add a result channel so each worker sends output back; collect all results in `main` after `wg.Wait()`.
- [ ] Replace the fixed `cpus` count with a `flag.Int` — make concurrency limits configurable at runtime.
- [ ] Add a `context.Context` so workers check `ctx.Done()` each iteration; compare this to the channel-close shutdown.
- [ ] Refactor into `func RunPool[T, R any](ctx context.Context, workers int, items []T, fn func(T) R) []R` — building a reusable abstraction is the real fluency goal.
