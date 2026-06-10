# atomic

Typed atomics: `atomic.Int64`, `atomic.Bool`, `atomic.Pointer[T]` (Go 1.19+).

## Key concepts

- **`atomic.Int64` / `atomic.Uint64` / `atomic.Bool`** — typed wrappers; no unsafe pointer arithmetic
- **`atomic.Pointer[T]`** — type-safe atomic pointer with `Load`, `Store`, `Swap`, `CompareAndSwap`
- **Compare-and-swap (CAS)** — the foundation of lock-free algorithms; only updates if current value matches expected
- **When to use** — counters, flags, single-pointer swaps; reach for `sync.Mutex` for anything more complex

### Compare-and-Swap (CAS) explained

CAS is a conditional atomic operation: `CompareAndSwap(expected, new)` only updates the value if it currently equals `expected`, returning `true` on success or `false` if another goroutine changed it first.

**Why it matters:**
```go
// Without CAS - race condition, lost updates
head := stack.Load()
newNode := &Node{next: head}
stack.Store(newNode)  // Last writer wins - some pushes lost!

// With CAS - retry loop ensures all updates succeed
for {
    head := stack.Load()
    newNode := &Node{next: head}
    if stack.CompareAndSwap(head, newNode) {
        break  // Success
    }
    // Another goroutine changed head - retry with new value
}
```

**CAS vs Swap:**
- `Swap(new)` — unconditional: always replaces, returns old value
- `CompareAndSwap(expected, new)` — conditional: only replaces if unchanged, enables retry loops

CAS lets you detect concurrent modifications and retry, making it the building block for lock-free data structures.

## Todo

- [x] Increment an `atomic.Int64` from 100 goroutines and verify the final value is exactly 100 — the canonical correctness test; run with `-race`.
- [x] Use `atomic.Bool` as a shutdown flag: set it from a signal handler, poll it in workers, verify clean exit.
- [x] Store a `*Config` in `atomic.Pointer[Config]`; update it from one goroutine while readers `Load` and use it concurrently without a mutex — practice the read-copy-update pattern.
- [x] Implement lock-free stack `Push` with a CAS loop on `atomic.Pointer[node]`; test under concurrent pushes with `-race`.

## Run

```bash
go run ./atomic/
```
