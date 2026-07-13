# iter

`iter.Seq` / range-over-func (Go 1.23+).

## What's here

`iter_test.go` — `Filter` and `Map` iterator adapters with lazy evaluation, early termination, and composition tests.

## Key concepts

- **`iter.Seq[V]`** — `func(yield func(V) bool)` — a push iterator; `yield` returns false on `break`
- **`iter.Seq2[K, V]`** — push iterator over key-value pairs
- **`iter.Pull` / `iter.Pull2`** — convert push → pull; returns `(next func() (V, bool), stop func())`
- **Early termination** — always check `yield`'s return value and return immediately on false
- **Lazy evaluation** — iterator adapters don't process until consumed

## Learnings

**Approximation constraints (`~`):**

The tilde (`~`) in type constraints means "any type whose underlying type is X". For example:

```go
func Values[Slice ~[]E, E any](s Slice) iter.Seq[E]
```

**Without tilde (`Slice []E`)** — only accepts exactly `[]E`:
```go
values := Values([]int{1,2,3})  // ✓ works
```

**With tilde (`Slice ~[]E`)** — accepts `[]E` AND any named type with underlying type `[]E`:
```go
values := Values([]int{1,2,3})  // ✓ works

type MyInts []int
myInts := MyInts{1,2,3}
values := Values(myInts)        // ✓ also works!
```

**Why it matters:**

In Go, you can define named types like `type UserIDs []int` or `type Names []string`. These have underlying types `[]int` and `[]string` but are distinct types. The tilde allows functions to accept both the base type and any named types based on it.

Without the tilde, `slices.Values` would only work with literal slices like `[]int`, not with `UserIDs`. The tilde makes it more flexible for real codebases where named slice types are common.

This is why you'll often see `~[]E`, `~map[K]V`, `~chan E` in generic constraints.

## Todo

- [x] Write `Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V]` and `Map[In, Out any](seq iter.Seq[In], f func(In) Out) iter.Seq[Out]`; chain them and collect with `slices.Collect`.
- [x] Write `Fibonacci() iter.Seq[int]` that yields indefinitely; use `break` after 10 values — practice the early-termination contract.
- [ ] Use `iter.Pull` to take exactly 3 values from a sequence without consuming the rest; call `stop()` and verify no goroutine leaks.
- [ ] Add an `InOrder() iter.Seq2[K, V]` method to a simple BST; verify keys arrive sorted when ranged over.

## Run

```bash
go test ./iter/
```
