package maps

import (
	"iter"
	"maps"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFilterMapByKey demonstrates using maps.All with an iterator adapter (Filter)
// to functionally filter a map's entries, then collect back to a map with maps.Collect.
// This shows the value of maps.All: it converts a map to iter.Seq2[K, V] which can be
// composed with iterator adapters before collecting back to a map.
//
// Note: For simple cases like this, an inline range loop is often clearer. The iterator
// pattern's value emerges when chaining multiple operations or reusing filter functions.
func TestFilterMapByKey(t *testing.T) {
	in := make(map[string]int)
	in["foo"] = 1
	in["bar"] = 2
	in["baz"] = 3

	// Pipeline: map → iter.Seq2 → filter → collect back to map
	result := maps.Collect(Filter(maps.All(in), func(s string) bool {
		return strings.HasPrefix(s, "b")
	}))

	// Should contain exactly "bar" and "baz"
	assert.Len(t, result, 2)
}

// Filter returns a new iter.Seq2[K, V] iterator that yields only entries where
// the predicate returns true for the key. This enables functional-style filtering
// of map entries when used with maps.All and maps.Collect:
//
//	filtered := maps.Collect(Filter(maps.All(m), keepFunc))
//
// The filtering happens lazily - entries are only checked when the returned iterator
// is consumed. Respects early termination: if the consumer breaks, stops consuming
// the source sequence.
func Filter[K comparable, V any](seq iter.Seq2[K, V], keep func(K) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		// Range over source key-value pairs
		for k, v := range seq {
			// Apply predicate to key
			if keep(k) {
				// Yield pair to consumer; stop if consumer breaks
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// TestInvert verifies that Invert swaps map keys and values.
// When multiple keys have the same value (collision), one key arbitrarily wins
// because map iteration order is non-deterministic. This test has a collision
// (both "b" and "c" map to 2) and verifies that either "b" or "c" wins, without
// assuming which one - this makes the test deterministic despite the collision.
func TestInvert(t *testing.T) {
	in := make(map[string]int)
	in["a"] = 1
	in["b"] = 2
	in["c"] = 2 // Collision: same value as "b"

	out := Invert(in)

	// Result should have 2 entries (one collision victim is dropped)
	assert.Len(t, out, 2)

	// Value 1 should always map to "a"
	if out[1] != "a" {
		t.Errorf("expected a, got %v", out[1])
	}

	// Value 2 should map to either "b" or "c" (non-deterministic)
	if out[2] != "b" && out[2] != "c" {
		t.Errorf("expected b or c, got %v", out[2])
	}
}

// Invert swaps map keys and values, returning a new map where the original values
// become keys and original keys become values. Both K and V must be comparable
// to be usable as map keys. When multiple keys map to the same value, the last
// key seen during iteration wins - which key that is depends on Go's randomized
// map iteration order.
func Invert[K, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K)
	for k, v := range m {
		// If v already exists in result, this overwrites it
		// Which key wins is non-deterministic due to map iteration order
		result[v] = k
	}
	return result
}
