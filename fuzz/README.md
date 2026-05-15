# fuzz

Property-based testing with `testing.F` (Go 1.18+).

## Key concepts

- **`f.Add(seed)`** — add a seed; the fuzzer generates mutations from these
- **`f.Fuzz(func(t, input))`** — asserts invariants, not exact outputs; called repeatedly
- **Property testing** — roundtrip, no-panic, commutativity, idempotency
- **Seed corpus** — crashing inputs saved to `testdata/fuzz/<Name>/` are always re-run as regression tests

## Todo

- [ ] Write `FuzzReverse` asserting `reverse(reverse(s)) == s` and that valid UTF-8 in produces valid UTF-8 out — the canonical first fuzz target.
- [ ] Write a key=value parser and `FuzzParser` asserting it never panics and `parse(format(kv)) == kv` — practice the encode/decode roundtrip property.
- [ ] Run `go test -fuzz=FuzzReverse -fuzztime=30s ./fuzz/` and inspect any saved corpus inputs; add them as seeds with `f.Add`.
- [ ] Write `FuzzMerge` with two string inputs asserting your merge function is commutative: `merge(a, b) == merge(b, a)`.

## Run

```bash
# Regression mode (seed corpus only)
go test ./fuzz/

# Fuzz mode
go test -fuzz=. -fuzztime=30s ./fuzz/
```
