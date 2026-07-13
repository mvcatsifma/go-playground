package iter

import (
	"iter"
	"slices"
	"testing"
)

// TestPull demonstrates iter.Pull which converts a push-style iterator (iter.Seq)
// to a pull-style iterator where the consumer controls when to get the next value.
// This test takes exactly 3 values from a 10-element sequence, explicitly calls stop(),
// and leaves the remaining 7 values unconsumed. Calling stop() is critical to prevent
// goroutine leaks - iter.Pull spawns a goroutine that must be cleaned up.
func TestPull(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// iter.Pull converts push iterator to pull iterator
	// Returns: next() to pull values, stop() to clean up the goroutine
	next, stop := iter.Pull(slices.Values(in))

	var result []int
	// Pull exactly 3 values
	for i := 0; i < 3; i++ {
		v, ok := next()
		if !ok {
			t.Fatal("sequence ended prematurely")
		}
		result = append(result, v)
	}

	// Explicitly stop - remaining 7 values are not consumed
	// This cleans up the goroutine spawned by iter.Pull
	stop()

	expected := []int{1, 2, 3}
	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

// TestFibonacci verifies that the Fibonacci iterator generates the correct sequence
// and demonstrates early termination - the iterator is infinite but the consumer
// stops after 10 values using break. This tests the yield contract: the iterator
// must check yield's return value and stop immediately when it returns false.
func TestFibonacci(t *testing.T) {
	var result []int
	count := 0
	for v := range Fibonacci() {
		if count == 10 {
			break
		}
		result = append(result, v)
		count++
	}

	// First 10 Fibonacci numbers: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

// Fibonacci returns an infinite iterator that yields the Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...
// The iterator never terminates on its own - the consumer controls when to stop using break.
// This demonstrates the iterator contract: always check yield's return value and stop when it returns false.
func Fibonacci() iter.Seq[int] {
	return func(yield func(int) bool) {
		// Start with first two Fibonacci numbers
		i := 0
		j := 1

		// Yield first value (0)
		if !yield(i) {
			return
		}

		// Yield second value (1)
		if !yield(j) {
			return
		}

		// Infinite loop: generate subsequent Fibonacci numbers
		// Each number is the sum of the previous two
		for {
			curr := i + j
			if !yield(curr) {
				// Consumer called break, stop generating
				return
			}
			// Shift window: previous becomes i, current becomes j
			i, j = j, curr
		}
	}
}

// TestFilterMapChain verifies that Filter and Map can be chained together and that
// the transformations are applied lazily - nothing happens until slices.Collect consumes
// the iterator. This demonstrates the composability of iterator adapters.
func TestFilterMapChain(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter keeps only even numbers
	filtered := Filter(slices.Values(in), func(v int) bool {
		return v%2 == 0
	})

	// Map multiplies each value by 10
	mapped := Map(filtered, func(n int) int {
		return n * 10
	})

	// Collect materializes the lazy iterator chain
	out := slices.Collect(mapped)

	// Verify result: evens from input, multiplied by 10
	expected := []int{20, 40, 60, 80, 100}
	if !slices.Equal(out, expected) {
		t.Errorf("got %v, want %v", out, expected)
	}
}

// Filter returns a new iterator that yields only values from seq for which keep returns true.
// The filtering happens lazily - values are only checked when the returned iterator is consumed.
// Respects early termination: if the consumer breaks, stops consuming the source sequence.
func Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		// Range over source sequence
		for v := range seq {
			// Apply predicate
			if keep(v) {
				// Yield to consumer; stop if consumer breaks
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Map returns a new iterator that applies f to each value from seq.
// The transformation happens lazily - f is only called when the returned iterator is consumed.
// Respects early termination: if the consumer breaks, stops consuming the source sequence.
func Map[In, Out any](seq iter.Seq[In], f func(In) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		// Range over source sequence
		for in := range seq {
			// Apply transformation
			out := f(in)
			// Yield to consumer; stop if consumer breaks
			if !yield(out) {
				return
			}
		}
	}
}
