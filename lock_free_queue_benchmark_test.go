package lockfreequeue

import (
	"testing"
)

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
