# embed

Embedding static files into a Go binary with `//go:embed`.

## What's here

`main.go` — embeds `testdata/` as `embed.FS`; reads and lists files.  
`main_test.go` — embeds a single file as a `string` variable.

## Todo

- [ ] Serve the embedded FS over HTTP with `http.FileServer(http.FS(f))` — the most common production use of `embed.FS`.
- [ ] Embed a JSON config file, unmarshal it at startup, and use the values to configure the program.
- [ ] Write a function `func ReadConfig(fsys fs.FS) (*Config, error)` that accepts `fs.FS` so tests can pass an `os.DirFS` instead of the embedded one.
- [ ] Walk the embedded FS with `fs.WalkDir` and print every file path and size — build the muscle memory for `fs.WalkDir` traversal.
