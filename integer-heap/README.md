# integer-heap

Max-heap over integers using `container/heap`.

## What's here

`main.go` — `IntHeap` implements `heap.Interface` as a max-heap; demonstrates `Init`, `Push`, `Pop`.

## Todo

- [ ] Add `Peek() int` returning the top element without removal — the most common operation after `Pop`.
- [ ] Implement `TopN(nums []int, n int) []int` that returns the N largest values using the heap.
- [ ] Convert to a min-heap by flipping `Less`; verify the pop order reverses — practice the single-line change that swaps heap direction.
- [ ] Call `heap.Remove(h, i)` to delete an element at a known index; verify the heap invariant still holds.
