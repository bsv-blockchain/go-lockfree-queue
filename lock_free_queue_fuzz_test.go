package lockfreequeue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// FuzzNewLockFreeQ verifies that a newly created queue is empty and can accept a value.
func FuzzNewLockFreeQ(f *testing.F) {
	f.Add(1)

	f.Fuzz(func(t *testing.T, seed int) {
		q := NewLockFreeQ[int]()
		require.NotNil(t, q)
		assert.True(t, q.IsEmpty())

		q.Enqueue(seed)
		assert.False(t, q.IsEmpty())
		val := q.Dequeue()
		require.NotNil(t, val)
		assert.Equal(t, seed, *val)
	})
}

// FuzzLockFreeQEnqueue ensures that enqueued values are returned in order.
func FuzzLockFreeQEnqueue(f *testing.F) {
	f.Add(0)
	f.Add(123)

	f.Fuzz(func(t *testing.T, v int) {
		q := NewLockFreeQ[int]()
		q.Enqueue(v)
		val := q.Dequeue()
		require.NotNil(t, val)
		assert.Equal(t, v, *val)
	})
}

// FuzzLockFreeQDequeue confirms dequeue returns values in FIFO order and nil when empty.
func FuzzLockFreeQDequeue(f *testing.F) {
	f.Add(1, 2)

	f.Fuzz(func(t *testing.T, first, second int) {
		q := NewLockFreeQ[int]()
		q.Enqueue(first)
		q.Enqueue(second)

		val1 := q.Dequeue()
		require.NotNil(t, val1)
		assert.Equal(t, first, *val1)

		val2 := q.Dequeue()
		require.NotNil(t, val2)
		assert.Equal(t, second, *val2)

		assert.Nil(t, q.Dequeue())
	})
}

// FuzzLockFreeQIsEmpty checks the IsEmpty logic across enqueue and dequeue cycles.
func FuzzLockFreeQIsEmpty(f *testing.F) {
	f.Add(10)
	f.Fuzz(func(t *testing.T, v int) {
		q := NewLockFreeQ[int]()
		assert.True(t, q.IsEmpty())

		q.Enqueue(v)
		assert.False(t, q.IsEmpty())

		require.NotNil(t, q.Dequeue())
		assert.True(t, q.IsEmpty())
	})
}
