# priority-queue

Priority queue with in-place updates via `heap.Fix`.

## What's here

`priority_queue.go` — `PriorityQueue[*Item]`; `update` changes priority without remove+insert.

## Todo

- [ ] Write tests for `Push`, `Pop`, and `update`; verify pop order matches descending priority.
- [ ] Add `Peek() *Item` returning the top without removal.
- [ ] Make `PriorityQueue` generic: `PriorityQueue[T any]` with a `less func(a, b T) bool` comparator — practice applying generics to an existing data structure.
- [ ] Write a min-heap variant by inverting `less`; run both in `main` to see the difference.
