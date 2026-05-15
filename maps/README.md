# maps

Generic map helpers (Go 1.21+).

## Key concepts

- **`maps.Clone`** — shallow copy; mutations to the clone don't affect the original
- **`maps.DeleteFunc`** — remove entries matching a predicate in one call
- **`maps.Equal`** — deep equality for comparable value types
- **`maps.Keys` / `maps.Values`** — return `iter.Seq`; range over them directly
- **`maps.All`** — `iter.Seq2[K, V]` over all pairs; feed it to `maps.Collect`

## Todo

- [ ] Write `Invert[K, V comparable](m map[K]V) map[V]K` using `maps.All` and a range loop — practice iterating key-value pairs with the new iter API.
- [ ] Use `maps.DeleteFunc` to remove stale cache entries (e.g. entries where `value.ExpiresAt.Before(time.Now())`) — the most common real-world use of DeleteFunc.
- [ ] Use `maps.Clone` + `maps.Copy` to merge two configs where the second overrides the first; verify keys from both maps appear in the result.
- [ ] Collect `maps.Keys` into a slice with `slices.Sorted`; iterate in deterministic order — build the habit of sorting map keys before ranging.

## Run

```bash
go run ./maps/
```
