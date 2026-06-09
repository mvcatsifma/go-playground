package atomic

import (
	"sync"
	"sync/atomic"
)

// concurrentIncrement spawns nr goroutines that each increment a shared
// atomic counter exactly once. Returns the final count after all goroutines
// complete. With correct atomic operations, the result will always equal nr.
// Run with -race to verify no data races occur.
func concurrentIncrement(nr int) int32 {
	wg := &sync.WaitGroup{}
	var count atomic.Int32

	for i := 0; i < nr; i++ {
		wg.Go(func() { // Go 1.26+: automatically handles Add(1)/Done()
			count.Add(1) // Atomic increment - thread-safe, no mutex needed
		})
	}
	wg.Wait() // Block until all goroutines complete

	return count.Load() // Atomic read of final value
}
