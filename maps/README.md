# maps

Generic map helpers (Go 1.21+).

## Key concepts

- **`maps.Clone`** — shallow copy; mutations to the clone don't affect the original
- **`maps.DeleteFunc`** — remove entries matching a predicate in one call
- **`maps.Equal`** — deep equality for comparable value types
- **`maps.Keys` / `maps.Values`** — return `iter.Seq`; range over them directly
- **`maps.All`** — `iter.Seq2[K, V]` over all pairs; feed it to `maps.Collect`

**When to use `maps.All`:**

Use `maps.All` when you need to compose operations on map entries using iterator adapters (Filter, Map, etc.) before collecting back to a map:

```go
// With maps.All: composable, reusable filters
filtered := maps.Collect(Filter(maps.All(m), predicate))
```

For simple cases, an inline range loop is often clearer:

```go
// Inline: direct, less ceremony for one-off filtering
result := make(map[K]V)
for k, v := range m {
    if predicate(k) {
        result[k] = v
    }
}
```

`maps.All` shines when:
- Chaining multiple transformations (filter → map values → collect)
- Reusing filter/transform functions across many maps
- Building composable map utilities

## Todo

- [x] Write `Invert[K, V comparable](m map[K]V) map[V]K` that swaps keys and values; handle collisions by keeping the last seen value.
- [x] Use `maps.All` with an iterator adapter (e.g., your `Filter` from iter) to filter a map's entries, then collect with `maps.Collect` — demonstrates composing map operations with iterator pipelines.
- [x] Use `maps.DeleteFunc` to remove entries that start with "f".
- [ ] Use `maps.Clone` + `maps.Copy` to merge two configs where the second overrides the first; verify keys from both maps appear in the result.
- [ ] Collect `maps.Keys` into a slice with `slices.Sorted`; iterate in deterministic order — build the habit of sorting map keys before ranging.

## Run

```bash
go run ./maps/
```
