# generics5

Type-parameterized HTTP response handling.

## What's here

`main.go` — `handleResponse[T httpResp]` dispatches on status code; `unmarshaller[T]` decodes the body.

## Todo

- [ ] Extract the JSON decode logic into `JSONUnmarshaller[T httpResp]() unmarshaller[T]` so callers don't repeat `io.ReadAll` + `json.Unmarshal`.
- [ ] Add `c` for `202 Accepted`; extend the `httpResp` union and add a branch in `handleResponse` — practice the mechanical work of extending a generic union.
- [ ] Add 404 and 500 handling that returns a typed `APIError`; verify the caller can reach it with `errors.As`.
- [ ] Write a table-driven test covering every status code branch; use `httptest` to build the responses.
