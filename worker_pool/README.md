# worker_pool

Semaphore-based goroutine limiting: a buffered channel caps concurrent goroutines at `limit`.

## What's here

`main.go` — acquires a slot before each goroutine; drains the semaphore at the end to wait for all work.

## Todo

- [ ] Add a result channel; collect all results in `main` after the semaphore drains — practice the full produce/consume cycle.
- [ ] Replace the semaphore drain with a `sync.WaitGroup`; compare the two: semaphore drain is clever but WaitGroup communicates intent more clearly.
- [ ] Add context cancellation: stop launching new goroutines when `ctx.Done()` fires; let in-flight goroutines finish.
- [ ] Rewrite using `errgroup.SetLimit(n)` and compare the total line count — understand when to reach for a higher-level abstraction.
