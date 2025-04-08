package LockFree

import (
	"testing"
)

func TestLockFreeStack(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)

	value, ok := stack.Pop()
	if !ok || value != 2 {
		t.Errorf("Expected Pop() to return (2, true), got: (%v, %v)", value, ok)
	}

	value, ok = stack.Pop()
	if !ok || value != 1 {
		t.Errorf("Expected Pop() to return (1, true), got: (%v, %v)", value, ok)
	}

	_, ok = stack.Pop()
	if ok {
		t.Error("Expected Pop() to return (0, false) for an empty stack")
	}
}