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

// BenchmarkLockFreeQ_Enqueue benchmarks adding items to the queue.
func BenchmarkLockFreeQ_Enqueue(b *testing.B) {
	q := NewLockFreeQ[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

// BenchmarkLockFreeQ_Dequeue benchmarks removing items from the queue.
func BenchmarkLockFreeQ_Dequeue(b *testing.B) {
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

// BenchmarkLockFreeQ_IsEmpty benchmarks the IsEmpty check.
func BenchmarkLockFreeQ_IsEmpty(b *testing.B) {
	q := NewLockFreeQ[int]()
	for i := 0; i < b.N; i++ {
		_ = q.IsEmpty()
	}
}
