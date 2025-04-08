package LockFree

import (
	"testing"
)

func TestLockFreeQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)

	value, ok := queue.Dequeue()
	if !ok || value != 1 {
		t.Errorf("Expected Dequeue() to return (1, true), got: (%v, %v)", value, ok)
	}

	queue.Enqueue(3)

	value, ok = queue.Dequeue()
	if !ok || value != 2 {
		t.Errorf("Expected Dequeue() to return (2, true), got: (%v, %v)", value, ok)
	}

	value, ok = queue.Dequeue()
	if !ok || value != 3 {
		t.Errorf("Expected Dequeue() to return (3, true), got: (%v, %v)", value, ok)
	}

	_, ok = queue.Dequeue()
	if ok {
		t.Error("Expected Dequeue() to return (0, false) for an empty queue")
	}
}