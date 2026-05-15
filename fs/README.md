# fs

Injectable filesystem via `io/fs`.

## What's here

`main.go` — `StdFS` implements `fs.ReadFileFS` and `fs.StatFS` over the real OS.

## Todo

- [ ] Write `MemFS` implementing `fs.ReadFileFS` with a `map[string][]byte` backing store; swap it into `main` and verify the output is identical.
- [ ] Write `func WalkFiles(fsys fs.FS, root string) ([]string, error)` using `fs.WalkDir`; test it with both `StdFS` and `MemFS`.
- [ ] Accept `fs.FS` instead of `*StdFS` in every function — practice narrowing to the interface at every call boundary.
- [ ] Pass an `embed.FS` to the same functions; confirm any `fs.FS` is interchangeable without code changes.
