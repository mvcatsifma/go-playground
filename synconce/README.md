# synconce

`sync.Once`, `sync.OnceValue`, `sync.Map`.

## Key concepts

- **`sync.Once`** — `Do(f)` runs `f` exactly once; safe for concurrent callers; ideal for lazy init
- **`sync.OnceValue[T]`** (Go 1.21) — memoises the return value; cleaner than Once + package var
- **`sync.Map`** — concurrent map; `Load`, `Store`, `LoadOrStore`, `Delete`, `Range`
- **`LoadOrStore`** — atomic get-or-create; prevents the load+store race in caches

## Todo

- [ ] Use `sync.Once` to initialise an expensive resource exactly once across 10 concurrent goroutines; assert the init function runs only once with a counter.
- [ ] Rewrite the singleton using `sync.OnceValue[*Resource](func() *Resource { ... })`; notice how it eliminates the package-level variable.
- [ ] Implement a simple cache with `LoadOrStore`: concurrent callers for the same key should compute the value only once — practice the canonical cache pattern.
- [ ] Use `sm.Range` to collect all entries into a `map[string]any` at the end of a concurrent test; verify every stored key is present.

## Run

```bash
go run ./synconce/
```
