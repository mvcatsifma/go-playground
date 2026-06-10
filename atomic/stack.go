package atomic

import (
	"fmt"
	"sync/atomic"
)

// Node represents a single element in the lock-free stack.
type Node struct {
	next  *Node
	value int
}

// Stack is a lock-free stack implementation using atomic.Pointer and compare-and-swap.
// Safe for concurrent access from multiple goroutines without mutexes.
type Stack struct {
	head atomic.Pointer[Node] // Atomic pointer to top of stack
}

// IsEmpty returns true if the stack has no elements.
func (s *Stack) IsEmpty() bool {
	return s.head.Load() == nil
}

// Pop removes and returns the top element from the stack.
// Returns (false, 0) if stack is empty, (true, value) on success.
// Uses CAS retry loop to handle concurrent modifications.
func (s *Stack) Pop() (bool, int) {
	for {
		head := s.head.Load() // Read current head
		if head == nil {
			fmt.Println("Stack: Pop: stack empty")
			return false, 0
		}
		// Try to update head to next node
		if s.head.CompareAndSwap(head, head.next) {
			return true, head.value // Success - return old head's value
		}
		// CAS failed - another goroutine modified stack, retry
		fmt.Println("Stack: Pop: CAS failed, retrying...")
	}
}

// Push adds a new element to the top of the stack.
// Uses CAS retry loop to handle concurrent pushes without losing updates.
func (s *Stack) Push(value int) {
	for {
		head := s.head.Load() // Read current head
		newNode := &Node{     // Create new node pointing to current head
			next:  head,
			value: value,
		}

		// Try to make newNode the new head
		if s.head.CompareAndSwap(head, newNode) {
			fmt.Printf("Stack: Push: value %d\n", value)
			return // Success
		}
		// CAS failed - another goroutine changed head, retry with new head
		fmt.Printf("Stack: Push: CAS failed for value %d, retrying...\n", value)
	}
}
