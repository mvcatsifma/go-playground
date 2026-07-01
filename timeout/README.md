# timeout

Context-aware directory walking with errgroup concurrency limiting, OS signals, and per-task result tracking.

## What's here

`main.go` — walks paths with stdlib `fs.WalkDir`, checks `ctx.Done()` every entry; `errgroup.SetLimit` caps concurrency at 2; SIGINT cancels; distinguishes timeout from cancellation.
`model.go` — `Task`, `TaskResult`, error sentinels (`TaskCanceled`, `TaskTimeout`).
`testdata/` — test directory structure with nested files and restricted permissions.

## Todo

- [ ] Use `testing/synctest` to test the timeout path without actually waiting 60 seconds.
- [ ] Write new tests using stdlib `fs` abstractions - use `testing/fstest.MapFS` for mocking instead of the deleted afero-based test doubles.
- [ ] Add test for restricted file handling - verify permission errors are handled correctly using `testdata/level1/restricted/secret.txt`.
- [x] Distinguish timeout vs cancellation in `TaskResult`: added `canceled` and `timeout` bool fields, set based on `ctx.Err()`.
- [x] Replace godirwalk with stdlib `fs.WalkDir`: simpler error handling, removed dependency on karrick/godirwalk.
- [x] Refactor semaphore to `errgroup.SetLimit`: replaced manual channel-based semaphore with errgroup for cleaner concurrency control.
