# httpmux

Enhanced `net/http.ServeMux` routing (Go 1.22+).

## Key concepts

- **Method + path patterns** — `"GET /articles/{id}"` matches only GET; no external router needed for basic REST
- **`{id}`** — captures one path segment; `{rest...}` captures the remainder
- **`r.PathValue("id")`** — retrieve a captured wildcard from the request
- **Precedence** — more specific patterns win; `GET /articles/new` beats `GET /articles/{id}`

## Todo

- [ ] Register `GET /articles/{id}` and `POST /articles`; use `httptest` to verify a POST to `/articles/123` returns 405 — method routing with no external dependency.
- [ ] Extract `{id}` with `r.PathValue("id")` and write it back in the response body; test with a real `httptest.Server`.
- [ ] Register both `GET /articles/new` and `GET /articles/{id}`; verify `/articles/new` hits the exact-match handler — understand precedence rules.
- [ ] Wrap the mux with a logging middleware: `func Logging(next http.Handler) http.Handler` — this is the pattern every Go web app uses.

## Run

```bash
go run ./httpmux/
```
