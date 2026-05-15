# unique

Value interning with `unique` (Go 1.23+).

## Key concepts

- **`unique.Make[T comparable](v T) Handle[T]`** — interns a value; two calls with equal values return pointer-equal handles
- **`Handle[T].Value() T`** — retrieve the canonical value
- **Use case** — deduplicate high-cardinality repeated values (hostnames, HTTP headers, tags) across many objects without a manual map

## Todo

- [ ] Intern the same string twice; verify the two `Handle[string]` values are equal with `==` — understand what handle equality means.
- [ ] Define `type Header struct{ Name, Value string }` and intern instances; use handles as map keys — practice using handles wherever you'd use the value as a key.
- [ ] Build a `TagSet` that interns each tag string on insertion; measure memory usage against a plain `[]string` with 1 000 000 repeated tags.
- [ ] Verify that `h1 == h2` implies `h1.Value() == h2.Value()` and that interning two equal values always returns the same handle — write this as a test, not just an observation.

## Run

```bash
go run ./unique/
```
