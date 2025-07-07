package lockfreequeue

import (
	"testing"
)

// BenchmarkLockFreeQueue benchmarks enqueueing and dequeuing values.
func BenchmarkLockFreeQueue(b *testing.B) {
	q := NewLockFreeQ[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		if val := q.Dequeue(); val == nil {
			b.Errorf("Expected a value, got nil at iteration %d", i)
		}
	}
}

// BenchmarkNewLockFreeQ measures the cost of creating new queue instances.
func BenchmarkNewLockFreeQ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewLockFreeQ[int]()
	}
}

// BenchmarkLockFreeQEnqueue benchmarks adding items to the queue.
func BenchmarkLockFreeQEnqueue(b *testing.B) {
	q := NewLockFreeQ[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

// BenchmarkLockFreeQDequeue benchmarks removing items from the queue.
func BenchmarkLockFreeQDequeue(b *testing.B) {
	q := NewLockFreeQ[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if q.Dequeue() == nil {
			b.Errorf("expected a value at iteration %d", i)
		}
	}
}

// BenchmarkLockFreeQIsEmpty benchmarks the IsEmpty check.
func BenchmarkLockFreeQIsEmpty(b *testing.B) {
	q := NewLockFreeQ[int]()
	for i := 0; i < b.N; i++ {
		_ = q.IsEmpty()
	}
}
