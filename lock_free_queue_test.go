package lockfreequeue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEnqueueDequeue tests the basic functionality of the lock-free queue
func TestEnqueueDequeue(t *testing.T) {
	q := NewLockFreeQ[int]() // Assuming your LockFreeQ works with int for this example

	// Enqueue elements
	q.Enqueue(1)
	q.Enqueue(2)

	// Dequeue and check elements
	val := q.Dequeue()
	assert.NotNil(t, val, "Expected 1, got nil")

	if val != nil {
		assert.Equal(t, 1, *val, "Expected 1, got %v", *val)
	}

	val = q.Dequeue()
	assert.NotNil(t, val, "Expected 2, got nil")

	if val != nil {
		assert.Equal(t, 2, *val, "Expected 2, got %v", *val)
	}

	// Check if the queue is empty
	assert.True(t, q.IsEmpty(), "Expected queue to be empty after dequeuing all elements")
	assert.Nil(t, q.Dequeue())
}

// TestConcurrentEnqueue tests concurrent enqueuing into the lock-free queue
func TestConcurrentEnqueue(t *testing.T) {
	q := NewLockFreeQ[int]()

	var wg sync.WaitGroup

	numWorkers := 100 // Number of concurrent goroutines
	numEnqueues := 10 // Number of enqueues per goroutine

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numEnqueues; j++ {
				q.Enqueue(workerID*numEnqueues + j)
			}
		}(i)
	}

	wg.Wait()

	// Assuming your dequeue is not concurrently safe, this part is tricky.
	// We can't guarantee the order of elements, but we can check if all elements are present.
	// This part of the test will need adjustment based on your dequeue method's thread safety.
	seen := make(map[int]bool)

	for i := 0; i < numWorkers*numEnqueues; i++ {
		val := q.Dequeue()
		assert.NotNil(t, val, "Expected a value, got nil at iteration %d", i)
		assert.False(t, seen[*val], "Duplicate value detected: %v", *val)

		seen[*val] = true
	}

	if !q.IsEmpty() {
		assert.Fail(t, "Expected queue to be empty after all dequeues")
	}
}
