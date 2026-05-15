package main

import (
	"crypto/sha256"
	"log"
)

// Hashes two file paths with SHA-256 to show how similar inputs produce distinct digests.
func main() {
	const (
		input1 = "/foo/bar/baz/bax.txt"
		input2 = "/foo/bar/baz/bad.txt"
	)

	h := sha256.New()
	h.Write([]byte(input1))

	log.Printf("%x", h.Sum(nil))

	h2 := sha256.New()
	h2.Write([]byte(input2))

	log.Printf("%x", h2.Sum(nil))
}
