package maps

import (
	"iter"
	"maps"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMergeConfigs demonstrates using maps.Clone + maps.Copy to merge two maps
// where the second map overrides the first on key conflicts. This is a common
// pattern for config merging: base config + override config = final config.
//
// maps.Clone creates a shallow copy, so the original maps remain unchanged.
// maps.Copy(dst, src) copies all entries from src into dst, overwriting any
// existing keys in dst.
//
// Pattern: merged := maps.Clone(base); maps.Copy(merged, override)
func TestMergeConfigs(t *testing.T) {
	// Base config (first)
	cfg1 := make(map[string]int)
	cfg1["foo"] = 1
	cfg1["bar"] = 3

	// Override config (second)
	cfg2 := make(map[string]int)
	cfg2["bar"] = 2 // Will override cfg1's bar=3
	cfg2["baz"] = 4

	// Clone cfg1, then copy cfg2 into it (cfg2 wins on conflicts)
	merged := maps.Clone(cfg1)
	maps.Copy(merged, cfg2)

	// Result should have keys from both maps, with cfg2 overriding cfg1
	expected := make(map[string]int)
	expected["foo"] = 1 // from cfg1
	expected["bar"] = 2 // from cfg2 (overrides cfg1's bar=3)
	expected["baz"] = 4 // from cfg2

	if !maps.Equal(merged, expected) {
		t.Errorf("expected %v, got %v", expected, merged)
	}
}

// TestDeleteFunc demonstrates maps.DeleteFunc which removes entries from a map
// in-place based on a predicate. The predicate returns true to DELETE the entry
// (opposite semantics from Filter's "keep" predicate).
//
// Key differences from Filter:
// - maps.DeleteFunc: modifies original map in-place, no return value
// - Filter: returns new map via maps.Collect, original unchanged
//
// Use DeleteFunc when you want to clean up an existing map without creating a copy.
func TestDeleteFunc(t *testing.T) {
	in := make(map[string]int)
	in["foo"] = 1
	in["bar"] = 2
	in["baz"] = 3

	// Delete entries where key starts with "f"
	// Predicate returns true to DELETE (not keep)
	maps.DeleteFunc(in, func(s string, i int) bool {
		return strings.HasPrefix(s, "f")
	})

	// "foo" should be removed, "bar" and "baz" remain
	assert.Len(t, in, 2)
	assert.Contains(t, in, "bar")
	assert.Contains(t, in, "baz")
	assert.NotContains(t, in, "foo")
}

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
