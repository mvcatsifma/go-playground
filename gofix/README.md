# gofix

Automated source rewriting with `go fix`.

## Key concepts

- **`go fix ./...`** — rewrites source files to newer API versions using built-in fixers
- **`go tool fix -list`** — list all fixers and which Go version each targets
- **`go tool fix -r <fixer>`** — apply a single named fixer
- **`gofmt -r 'pattern -> replacement'`** — custom AST-level rewrite for one-off renames

## Todo

- [ ] Run `go tool fix -list` and identify which fixers apply to Go 1.22+; read what each one changes.
- [ ] Write a file using `io/ioutil.ReadAll`; run `go fix ./gofix/` and verify it rewrites to `io.ReadAll` — practice the fix loop.
- [ ] Use `gofmt -r 'a.Foo(b) -> a.Bar(b)'` to rename a method call across a package; verify all call sites update.
- [ ] Add `go generate ./... && git diff --exit-code` to a `Makefile`; understand how `go fix` fits into a CI pipeline.

## Run

```bash
go tool fix -list
go fix ./gofix/
```
