# hash

`crypto/sha256` hashing.

## What's here

`main.go` — hashes two file paths to show how small input differences produce completely different digests.

## Todo

- [ ] Hash a real file by streaming it: `io.Copy(h, f)` — never load large files into memory just to hash them.
- [ ] Compute HMAC-SHA256 with `crypto/hmac`; verify that changing either the key or the message produces a different MAC.
- [ ] Build a file integrity checker: walk a directory, hash every file, write a manifest; re-run and report any changed or missing files.
- [ ] Return the digest as a hex string with `fmt.Sprintf("%x", h.Sum(nil))` — write a `HashFile(path string) (string, error)` helper and call it from a test.
