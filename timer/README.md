# timer

`time.Timer` — a negative duration fires immediately.

## What's here

`main.go` — `time.NewTimer(-1 * time.Second)` fires on the first channel read.

## Todo

- [ ] Implement the safe Stop+drain+Reset pattern: `if !t.Stop() { <-t.C }; t.Reset(d)` — this is the pattern you'll need in production and it's easy to get wrong.
- [ ] Write `Debounce(f func(), d time.Duration) func()` using `time.AfterFunc` and `timer.Reset`; test that rapid calls only trigger `f` once after the delay.
- [ ] Add a `time.Ticker` that fires 3 times then stops; compare `Timer` (one-shot) vs `Ticker` (recurring) usage patterns.
- [ ] Write a test using `testing/synctest` so the debounce test completes instantly without real wall time.
