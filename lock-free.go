package LockFree

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type List struct {
	head unsafe.Pointer
}

type Node struct {
	value int
	next unsafe.Pointer
}

func (l *List) Add(value int) {
	node := &Node{value: value}
	for {
		oldHead := atomic.LoadPointer(&l.head)
		node.next = oldHead

		if atomic.CompareAndSwapPointer(&l.head, oldHead, unsafe.Pointer(node)) {
			break
		}
	}
}

func (l *List) Print() {
	curr := atomic.LoadPointer(&l.head)

	for curr != nil {
		node := (*Node)(curr)
		fmt.Println(node.value)
		curr = atomic.LoadPointer((*unsafe.Pointer)(node.next))
	}
}

// TODO: Make finding by index (FindByIndex func)


type StackNode struct {
	value int          
	next  unsafe.Pointer 
}

type Stack struct {
	top unsafe.Pointer
}

func (s *Stack) Push(value int) {
	node := &StackNode{value: value}

	for {
		oldTop := atomic.LoadPointer(&s.top)
		node.next = oldTop

		if atomic.CompareAndSwapPointer(&s.top, oldTop, unsafe.Pointer(node)) {
			break
		}
	}
}

func (s *Stack) Pop() (int, bool) {
	for {
		oldTop := atomic.LoadPointer(&s.top)
		if oldTop == nil {
			return 0, false
		}

		newTop := (*StackNode)(oldTop).next

		if atomic.CompareAndSwapPointer(&s.top, oldTop, unsafe.Pointer(newTop)) {
			return (*StackNode)(oldTop).value, true
		}
	}
}