# iter

`iter.Seq` / range-over-func (Go 1.23+).

## Key concepts

- **`iter.Seq[V]`** — `func(yield func(V) bool)` — a push iterator; `yield` returns false on `break`
- **`iter.Seq2[K, V]`** — push iterator over key-value pairs
- **`iter.Pull` / `iter.Pull2`** — convert push → pull; returns `(next func() (V, bool), stop func())`
- **Early termination** — always check `yield`'s return value and return immediately on false

## Todo

- [ ] Write `Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V]` and `Map[In, Out any](seq iter.Seq[In], f func(In) Out) iter.Seq[Out]`; chain them and collect with `slices.Collect`.
- [ ] Write `Fibonacci() iter.Seq[int]` that yields indefinitely; use `break` after 10 values — practice the early-termination contract.
- [ ] Use `iter.Pull` to take exactly 3 values from a sequence without consuming the rest; call `stop()` and verify no goroutine leaks.
- [ ] Add an `InOrder() iter.Seq2[K, V]` method to a simple BST; verify keys arrive sorted when ranged over.

## Run

```bash
go run ./iter/
```
