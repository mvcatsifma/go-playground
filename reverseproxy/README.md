# reverseproxy

`httputil.ReverseProxy` with the `Rewrite` hook (Go 1.20+).

## Key concepts

- **`Rewrite`** — replaces the deprecated `Director`; receives `*ProxyRequest` with both inbound and outbound requests
- **`ProxyRequest.SetURL`** — rewrite the outbound target host and path
- **`ProxyRequest.SetXForwarded`** — sets `X-Forwarded-For/Host/Proto` correctly in one call
- **`ModifyResponse`** — intercept and modify the upstream response before forwarding
- **`ErrorHandler`** — custom response when the upstream is unreachable

## Todo

- [ ] Start an `httptest.Server` as the upstream; proxy requests to it using `Rewrite` + `SetURL` — the minimum working reverse proxy.
- [ ] Strip a `/api` prefix in `Rewrite` so `/api/users` proxies to `/users` on the upstream — the most common path-rewriting use case.
- [ ] Inject an `X-Internal-Token` header in `Rewrite`; verify the upstream receives it and the original inbound request does not contain it.
- [ ] Add `ModifyResponse` to inject a `Cache-Control` header into every upstream response; shut down the upstream mid-test and verify `ErrorHandler` returns a `502` with a JSON body.

## Run

```bash
go run ./reverseproxy/
```
