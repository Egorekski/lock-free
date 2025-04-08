package LockFree

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// Lock-free List
type List struct {
	head unsafe.Pointer
}

type ListNode struct {
	value int
	next unsafe.Pointer
}

func NewList() *List {
	node := &ListNode{}
	return &List{
		head: unsafe.Pointer(node),
	}
}

func (l *List) Add(value int) {
	node := &ListNode{value: value}
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
		node := (*ListNode)(curr)
		fmt.Println(node.value)
		curr = atomic.LoadPointer((*unsafe.Pointer)(node.next))
	}
}

// TODO: Make finding by index (FindByIndex func)

// Lock-free Stack
type StackNode struct {
	value int          
	next  unsafe.Pointer 
}

type Stack struct {
	top unsafe.Pointer
}

func NewStack() *Stack {
	node := &StackNode{}
	return &Stack{
		top: unsafe.Pointer(node),
	}
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
		oldTop := atomic.LoadInt32((*int32)(s.top))
		if oldTop == 0 {
			return 0, false
		}

		oldtop := atomic.LoadPointer(&s.top)

		newTop := (*StackNode)(oldtop).next

		if atomic.CompareAndSwapPointer(&s.top, oldtop, unsafe.Pointer(newTop)) {
			return (*StackNode)(oldtop).value, true
		}
	}
}

// Lock-free Queue
type QueueNode struct {
    value int
    next  *QueueNode
}

type Queue struct {
    head unsafe.Pointer
    tail unsafe.Pointer
}

func NewQueue() *Queue {
	hd := &QueueNode{value: 0}
	return &Queue{
		head: unsafe.Pointer(hd),
	}
}

func (q *Queue) Enqueue(value int) {
	valuePtr := unsafe.Pointer(&value)
	q.tail = atomic.LoadPointer((*unsafe.Pointer)(valuePtr))
}

func (q *Queue) Dequeue() (int, bool) {
	dequeued := atomic.LoadInt32((*int32)(q.head))
	swapped := atomic.CompareAndSwapPointer((*unsafe.Pointer)(q.head), q.head, nil)
	return int(dequeued), swapped
}