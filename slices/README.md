# slices

Generic slice helpers (Go 1.21+).

## Key concepts

- **`slices.Sort` / `slices.SortFunc`** — in-place; `SortFunc` takes a `cmp.Compare`-style comparator
- **`slices.Contains` / `slices.Index`** — linear search
- **`slices.BinarySearch`** — O(log n) on a sorted slice; returns `(index, found)`
- **`slices.Compact`** — removes consecutive duplicates
- **`slices.Collect`** — drains an `iter.Seq[V]` into a `[]V`

## Todo

- [ ] Sort a `[]Person` by two fields using `cmp.Or(cmp.Compare(a.Last, b.Last), cmp.Compare(a.First, b.First))` — this is the idiomatic multi-key sort in Go 1.21+.
- [ ] Implement `Deduplicate[T comparable](s []T) []T` using `slices.Sort` + `slices.Compact` — build the muscle memory for this common pattern.
- [ ] Write `Map[In, Out any](s []In, f func(In) Out) []Out` using a range loop; compare with collecting from an `iter.Seq` via `slices.Collect`.
- [ ] Use `slices.BinarySearch` to implement a fast `Contains` on a sorted slice; write a test asserting correct index and found values at boundaries.

## Run

```bash
go run ./slices/
```
