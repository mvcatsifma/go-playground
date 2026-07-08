package synctest

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

// TestTime demonstrates synctest's controlled time advancement.
// Time starts at midnight UTC 2000-01-01 and only advances when explicitly sleeping.
// All time operations (Sleep, Now, Since) are deterministic and instant in wall-clock time.
func TestTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now() // Always midnight UTC 2000-01-01 in synctest

		go func() {
			time.Sleep(1 * time.Second) // Simulated time advances by 1s (instant in reality)
			t.Log(time.Since(start))    // Always logs "1s"
		}()

		// Main goroutine sleeps 2s, advancing simulated time.
		// The spawned goroutine runs first (1s < 2s) before this Sleep returns.
		time.Sleep(2 * time.Second)
		t.Log(time.Since(start)) // Always logs "2s"
	})
}

// TestWait demonstrates synctest.Wait() blocking until all goroutines finish.
// Within synctest's deterministic scheduling, the simple bool access is safe.
func TestWait(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		done := false
		go func() {
			done = true
		}()
		// Wait blocks until all spawned goroutines complete.
		// The goroutine runs immediately (no sleep blocking it).
		synctest.Wait()
		t.Log(done) // Always logs "true" - goroutine has finished
	})
}

// TestContextAfterFunc demonstrates testing context.AfterFunc with synctest.
// AfterFunc callbacks execute asynchronously, and synctest.Wait() ensures they complete
// before proceeding, making async behavior testable deterministically.
func TestContextAfterFunc(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a context.Context which can be canceled.
		ctx, cancel := context.WithCancel(t.Context())

		// context.AfterFunc registers a function to be called
		// when a context is canceled (runs in separate goroutine).
		// The callback is registered but dormant - no goroutine is spawned yet.
		afterFuncCalled := false
		context.AfterFunc(ctx, func() {
			afterFuncCalled = true
		})

		// The context has not been canceled, so the AfterFunc is not called.
		// First Wait() tests the negative case: verifies no goroutines are running
		// since the callback hasn't been triggered yet. Returns immediately since
		// nothing is blocked. This proves the callback is dormant and won't run
		// until cancel() is called.
		synctest.Wait()
		if afterFuncCalled {
			t.Fatalf("before context is canceled: AfterFunc called")
		}

		// Cancel the context - this triggers AfterFunc in a new goroutine.
		cancel()

		// Second Wait() blocks until the AfterFunc goroutine completes execution.
		// This demonstrates testing async behavior deterministically.
		synctest.Wait()
		if !afterFuncCalled {
			t.Fatalf("after context is canceled: AfterFunc not called")
		}
	})
}
