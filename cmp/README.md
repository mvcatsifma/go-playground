# cmp

Generic comparison helpers (Go 1.21+).

## Key concepts

- **`cmp.Compare[T Ordered](x, y T) int`** — returns -1/0/1; the comparator `slices.SortFunc` expects
- **`cmp.Less[T Ordered](x, y T) bool`** — convenience wrapper
- **`cmp.Or[T comparable](vals ...T) T`** — first non-zero value; replaces `if a != "" { return a }` chains
- **`cmp.Ordered`** — satisfied by all integer, float, and string types

## Todo

- [ ] Sort a `[]Person` by last name then first name: `cmp.Or(cmp.Compare(a.Last, b.Last), cmp.Compare(a.First, b.First))` — the canonical multi-key sort idiom.
- [ ] Use `cmp.Or(envValue, flagValue, defaultValue)` to resolve a config string; verify each fallback level works — replace every `if s == ""` chain with `cmp.Or`.
- [ ] Write `Clamp[T cmp.Ordered](v, lo, hi T) T` using `cmp.Compare`; test it at and around the boundaries.
- [ ] Define `type Priority int` and sort tasks by priority using `cmp.Compare` — practice applying `cmp.Ordered` to a named type.

## Run

```bash
go run ./cmp/
```
