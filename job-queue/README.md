# job-queue

Priority queue with index tracking for O(log n) updates.

## What's here

`main.go` — `JobQueue` / `Job` implement `heap.Interface`; jobs store their own index.  
`main_test.go` — unit tests for `Len`, `Push`, `Pop`, `Less`, `Swap`.

## Todo

- [ ] Implement `Update(q *JobQueue, job *Job, newPriority int)` using `heap.Fix`; test that the next `Pop` reflects the updated priority.
- [ ] Implement `Remove(q *JobQueue, job *Job)` using `heap.Remove`; verify the removed job never surfaces during iteration.
- [ ] Add a concurrency-safe wrapper with `sync.Mutex`; write a test with concurrent pushes and pops under `-race`.
- [ ] Add `String() string` to `Job` and log the full queue state after each operation to build intuition for heap shape.
