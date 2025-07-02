// Package lockfreequeue provides operations for a FIFO structure with
// operations to enqueue and dequeue generic values.
package lockfreequeue

import (
	"sync/atomic"
)

type node[T any] struct {
	value T
	next  atomic.Pointer[node[T]]
}

// LockFreeQ represents a FIFO structure with operations to enqueue
// and dequeue generic values.
// This implementation is concurrent safe for queueing, but not for dequeueing.
// Reference: https://www.cs.rochester.edu/research/synchronization/pseudocode/queues.html
type LockFreeQ[T any] struct {
	head *node[T]
	tail atomic.Pointer[node[T]]
}

// NewLockFreeQ creates and initializes a LockFreeQueue
func NewLockFreeQ[T any]() *LockFreeQ[T] {
	return &LockFreeQ[T]{
		head: &node[T]{},
		tail: atomic.Pointer[node[T]]{},
	}
}

// Enqueue adds a value to the end of the queue in a lock-free, thread-safe manner.
//
// This method performs the following steps:
// - Allocates a new node containing the provided value
// - Atomically swaps the queue's tail pointer to the new node
//   - If the queue was empty, links the head to the new node
//   - Otherwise, links the previous tail's next pointer to the new node
//
// Parameters:
// - v: the value to enqueue; may be any type supported by the generic parameter T
//
// Returns:
// - None
//
// Side Effects:
// - Mutates the internal state of the queue by appending a new node
// - Uses atomic operations to ensure thread safety for concurrent enqueues
func (q *LockFreeQ[T]) Enqueue(v T) {
	newNode := &node[T]{value: v}
	prev := q.tail.Swap(newNode)

	if prev == nil {
		q.head.next.Store(newNode)
		return
	}

	prev.next.Store(newNode)
}

// Dequeue removes and returns the value at the front of the queue in a lock-free queue.
//
// This method performs the following steps:
// - Loads the next node after the current head
//   - If the next node is nil, the queue is empty and nil is returned
//   - Otherwise, advances the head pointer to the next node
//
// - Returns a pointer to the value stored in the dequeued node
//
// Parameters:
// - None
//
// Returns:
// - Pointer to the value of type T at the front of the queue, or nil if the queue is empty
//
// Side Effects:
// - Mutates the internal state of the queue by advancing the head pointer
func (q *LockFreeQ[T]) Dequeue() *T {
	next := q.head.next.Load()

	if next == nil {
		return nil
	}

	q.head = next

	return &next.value
}

// IsEmpty reports whether the queue contains any elements.
//
// This method performs the following steps:
// - Loads the next node after the current head
//   - If the next node is nil, the queue is empty
//   - Otherwise, the queue contains at least one element
//
// Parameters:
// - None
//
// Returns:
// - true if the queue is empty; false otherwise
//
// Side Effects:
// - None
func (q *LockFreeQ[T]) IsEmpty() bool {
	return q.head.next.Load() == nil
}
