# template

`html/template` rendering.

## What's here

`main.go` — parses an inline template with `{{.Field}}` placeholders; executes with a `map[string]any`.

## Todo

- [ ] Move the template to an embedded file with `//go:embed`; reload it on startup — the standard production pattern.
- [ ] Replace the `map[string]any` with a typed struct; observe how the template engine rejects misspelled field names at execution time.
- [ ] Add a `template.FuncMap` with an `upper` function (`strings.ToUpper`); use `{{upper .Name}}` inside the template.
- [ ] Write a test that executes the template into a `bytes.Buffer` and asserts the exact output string.
