# timeout

Context-aware directory walking with a semaphore, OS signals, and per-task result tracking.

## What's here

`main.go` — walks a path with godirwalk, checks `ctx.Done()` every entry; semaphore caps concurrency at 2; SIGINT cancels.  
`model.go` — `Task`, `TaskResult`, `Fs` interface.  
`broken_fs.go` / `stub_fs.go` — test doubles.  
`system/fs.go` — production `Fs` via afero.

## Todo

- [ ] Use `testing/synctest` to test the timeout path without actually waiting 60 seconds.
- [ ] Distinguish timeout vs cancellation in `TaskResult` (the existing TODO): add `timedOut bool` and set it when `ctx.Err() == context.DeadlineExceeded`.
- [ ] Make the semaphore limit a parameter to `handleTask` instead of a package-level constant.
- [ ] Replace godirwalk with stdlib `fs.WalkDir` and compare how error handling changes.
