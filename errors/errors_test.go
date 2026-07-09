package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sentinel temporary error - used when error has no additional data.
// Demonstrates errors.Is pattern: checking for a specific error value.
var Temporary = errors.New("temporary error")

// NetworkError represents a temporary network failure with retry capability.
// Demonstrates errors.As pattern: implementing an interface to indicate behavior.
// Always returns true from Temporary(), making all network errors retryable.
type NetworkError struct {
	Op  string
	Msg string
}

func (n *NetworkError) Error() string {
	return fmt.Sprintf("network error during %s: %s", n.Op, n.Msg)
}

func (n *NetworkError) Temporary() bool {
	return true // Network errors are always temporary
}

// DatabaseError represents a database error that may or may not be temporary.
// Demonstrates dynamic retry logic: Temporary() returns true only when Retryable field is set.
// This shows how errors can carry metadata to determine retry behavior at runtime.
type DatabaseError struct {
	Query     string
	Retryable bool // Dynamic field: true for deadlocks/timeouts, false for constraint violations
}

func (d *DatabaseError) Error() string {
	return fmt.Sprintf("database error executing query: %s", d.Query)
}

func (d *DatabaseError) Temporary() bool {
	return d.Retryable // Dynamic: depends on error type
}

// temporaryError interface defines errors that can be retried.
// Used with errors.As to check if an error implements retry behavior,
// regardless of the concrete error type.
type temporaryError interface {
	Temporary() bool
}

// Retry executes fn up to n times, retrying if the error is temporary.
// Demonstrates two error checking patterns:
// 1. errors.Is - checks if error matches a specific sentinel (e.g., Temporary)
// 2. errors.As - checks if error implements an interface (e.g., Temporary() bool)
// Real code should support both patterns for maximum flexibility.
func Retry(n int, fn func() error) (int, error) {
	var c int
	var err error
RetryLoop:
	for c < n {
		c++
		err = fn()

		// Pattern 1: Check for specific temporary sentinel using errors.Is
		// Use this for simple cases where a single sentinel value indicates retry
		if errors.Is(err, Temporary) {
			continue
		}

		// Pattern 2: Check if error implements Temporary() interface using errors.As
		// Use this for rich error types that carry metadata and dynamic retry logic
		var tempErr temporaryError
		if errors.As(err, &tempErr) && tempErr.Temporary() {
			continue
		}

		// Error is nil, permanent, or doesn't match either pattern - stop retrying
		break RetryLoop
	}
	return c, err
}

// TestRetry_NilError verifies Retry stops immediately when function succeeds (returns nil).
// Expects exactly 1 attempt since success means no retry needed.
func TestRetry_NilError(t *testing.T) {
	c, err := Retry(2, func() error {
		return nil
	})

	assert.Equal(t, 1, c)
	assert.Nil(t, err)
}

// TestRetry_PermError verifies Retry stops immediately on permanent error.
// Errors without Temporary() method and not matching any sentinel are permanent.
// Expects exactly 1 attempt since permanent errors should not be retried.
func TestRetry_PermError(t *testing.T) {
	c, err := Retry(2, func() error {
		return errors.New("permanent error")
	})

	assert.Equal(t, 1, c)
	assert.NotNil(t, err)
}

// TestRetry_TempError verifies Retry exhausts all attempts when error is temporary.
// Uses the Temporary sentinel error which matches via errors.Is check.
// Expects all retry attempts to be used since the error continues being temporary.
func TestRetry_TempError(t *testing.T) {
	c, err := Retry(2, func() error {
		return Temporary
	})

	assert.Equal(t, 2, c) // Retried maximum times
	assert.True(t, errors.Is(err, Temporary))
}

// TestRetry_MultipleTemporaryTypes demonstrates errors.As checking for Temporary() interface
// across different error types (NetworkError, DatabaseError with Retryable=true).
// This shows the advantage of errors.As over errors.Is: multiple error types can indicate
// temporary failure without being the same sentinel value.
func TestRetry_MultipleTemporaryTypes(t *testing.T) {
	t.Run("NetworkError is temporary", func(t *testing.T) {
		c, err := Retry(3, func() error {
			return &NetworkError{Op: "dial", Msg: "connection refused"}
		})

		assert.Equal(t, 3, c) // Retried all attempts
		var netErr *NetworkError
		assert.True(t, errors.As(err, &netErr))
		assert.True(t, netErr.Temporary())
	})

	t.Run("DatabaseError with Retryable=true is temporary", func(t *testing.T) {
		c, err := Retry(3, func() error {
			return &DatabaseError{Query: "SELECT * FROM users", Retryable: true}
		})

		assert.Equal(t, 3, c) // Retried all attempts
		var dbErr *DatabaseError
		assert.True(t, errors.As(err, &dbErr))
		assert.True(t, dbErr.Temporary())
	})

	t.Run("DatabaseError with Retryable=false is permanent", func(t *testing.T) {
		c, err := Retry(3, func() error {
			return &DatabaseError{Query: "SELECT * FROM users", Retryable: false}
		})

		assert.Equal(t, 1, c) // Stopped immediately
		var dbErr *DatabaseError
		assert.True(t, errors.As(err, &dbErr))
		assert.False(t, dbErr.Temporary())
	})
}

// TestErrorIsTraversesWrappedErrors verifies that errors.Is can find a sentinel error
// through multiple levels of wrapping. Creates a chain: err2 -> err1 -> XError, then
// confirms errors.Is(err2, XError) returns true while errors.Is(err2, AnotherError) returns false.
func TestErrorIsTraversesWrappedErrors(t *testing.T) {
	var XError = errors.New("xerror")
	var AnotherError = errors.New("another")

	err1 := fmt.Errorf("error: %w", XError)
	err2 := fmt.Errorf("error: %w", err1)
	err3 := fmt.Errorf("error: %w", err2)

	assert.True(t, errors.Is(err3, XError))
	assert.False(t, errors.Is(err3, AnotherError))
}

// TestErrorsJoin verifies that errors.Join (Go 1.20+) combines multiple errors into one,
// and that errors.Is can find each individual error within the joined result.
// errors.Join internally implements Unwrap() []error (multi-error unwrapping) which allows
// errors.Is to traverse all joined errors. Useful for collecting multiple independent errors
// (e.g., validation failures, cleanup errors) without losing any information.
func TestErrorsJoin(t *testing.T) {
	var XError = errors.New("xerror")
	var YError = errors.New("yerror")

	// Join creates a single error containing both XError and YError
	err := errors.Join(XError, YError)

	// errors.Is can find both sentinels in the joined error
	assert.True(t, errors.Is(err, XError))
	assert.True(t, errors.Is(err, YError))
}

// MultiError demonstrates custom multi-error unwrapping using Unwrap() []error pattern.
// This is the same pattern that errors.Join uses internally. By implementing Unwrap() []error,
// the error type tells errors.Is and errors.As to traverse all wrapped errors in the slice.
// This differs from Unwrap() error (single unwrapping) which only unwraps one level.
//
// How it works: errors.Is checks if an error implements interface{ Unwrap() []error }.
// When found, it iterates through the slice recursively checking each error:
//
//	case interface{ Unwrap() []error }:
//	    for _, err := range x.Unwrap() {
//	        if is(err, target, targetComparable) {
//	            return true
//	        }
//	    }
//
// This is not magic - it's a defined interface pattern that errors.Is explicitly handles.
type MultiError struct {
	Message string
	Errs    []error
}

func (m *MultiError) Error() string {
	return m.Message
}

// Unwrap returns all wrapped errors as a slice, enabling errors.Is to traverse them.
// When errors.Is encounters this method, it iterates through the slice and recursively
// checks each error against the target. This is the mechanism that makes errors.Join work.
func (m *MultiError) Unwrap() []error {
	return m.Errs
}

// TestMultiErrorUnwrap verifies that custom error types implementing Unwrap() []error
// allow errors.Is to traverse all wrapped errors in the slice. This demonstrates the
// internal mechanism that errors.Join uses - it's not magic, just an error type with
// Unwrap() []error. This pattern is useful when you need custom multi-error types with
// additional context or behavior beyond what errors.Join provides.
func TestMultiErrorUnwrap(t *testing.T) {
	var ValidationErr = errors.New("validation failed")
	var AuthErr = errors.New("authentication failed")
	var StorageErr = errors.New("storage failed")

	// Create custom multi-error wrapping three sentinel errors
	multiErr := &MultiError{
		Message: "request processing failed with multiple errors",
		Errs:    []error{ValidationErr, AuthErr, StorageErr},
	}

	// errors.Is traverses the slice returned by Unwrap() []error
	// and can find all three wrapped sentinels
	assert.True(t, errors.Is(multiErr, ValidationErr))
	assert.True(t, errors.Is(multiErr, AuthErr))
	assert.True(t, errors.Is(multiErr, StorageErr))

	// Errors not in the slice cannot be found
	var UnrelatedErr = errors.New("unrelated")
	assert.False(t, errors.Is(multiErr, UnrelatedErr))
}
