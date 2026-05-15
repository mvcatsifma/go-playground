# stringer

Auto-generating `String()` methods with `go generate`.

## What's here

`pill.go` — `Pill` enum with `-linecomment` to override display names.  
`pill_string.go` — generated output; do not edit.

## Todo

- [ ] Add a `Caffeine` constant; run `go generate ./stringer/` and verify `pill_string.go` updates automatically — practice the generate loop.
- [ ] Add a second type `Direction int` with `North`, `South`, `East`, `West`; add a second `//go:generate` line and regenerate both at once.
- [ ] Write a test asserting `String()` for every `Pill` value, including the `-linecomment` override (`Paracetamol` → `"Foo"`).
- [ ] Add a `Makefile` target `generate: ; go generate ./...` so generation is one command away in CI.
