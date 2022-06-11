package queue

import "github.com/maverick6912/dsa_go/errors"

/*
Implementation of stack using 2 queues(push heavy):

	Push(val) operation’s step are described below:
        - Enqueue x to pQ
        - One by one dequeue everything from hQ and enqueue to pQ.
        - Swap the names of hQ and pQ
    Pop(s) operation’s function are described below:
        - Dequeue an item from hQ and return it.

---------Dry Run---------------
	pQ=[]
	hQ=[]

	s.Enq(2)

	pQ=[2]
	pQ=pQ+hQ => [2]

	hQ.Clear()

	tQ=[2]
	pQ=[]
	hQ=[2]
	cSize=1
	--------------------------
	pQ=[]
	hQ=[2]

	s.Enq(5)

	pQ=[5]
	pQ=pQ+hQ => [2,5]

	hQ.Clear()
	tQ=[2,5]
	pQ=[]
	hQ=[2,5]

	cSize=2
	----------------------------
	pQ=[]
	hQ=[2]

	s.Enq(9)

	pQ=[9]
	pQ=pQ+hQ => [2,5,9]

	hQ.Clear()
	tQ=[2,5,9]
	pQ=[]
	hQ=[2,5,9]

	cSize=3
	------------------------------

	s.Pop()
	9
	s.Pop()
	5
--------------Glossary-----------------
	hQ-> holdQueue, used to hold elements in stack
	pQ-> pushQueue, used to push new element tp stack
	tQ-> tempQueue, used to swap the queues
	cSize-> currSize of stack
*/

type Stack[T any] struct {
	cmp       func(T, T) int
	pushQueue *Queue[T]
	holdQueue *Queue[T]
	currSize  int
}

// TODO: add Enque and Dequeue ops for queue here.
type Queue[T any] struct {
	elements []T
	// size     int
}

// New initializes a `Stack`
func (s *Stack[T]) New(cmp func(T, T) int) *Stack[T] {
	var p, h *Queue[T]
	p = p.newQueue()
	h = h.newQueue()
	return &Stack[T]{pushQueue: p, holdQueue: h, cmp: cmp}
}

// newQueue initializes a `Queue`.
// Only for internal use.
func (q *Queue[T]) newQueue() *Queue[T] {
	return &Queue[T]{elements: make([]T, 0)}
}

// Push adds given value to the stack.
// Doesn't do anything if stack is un-initialized.
func (s *Stack[T]) Push(val T) error {
	if s == nil {
		return errors.UninitializedError
	}
	if s.holdQueue == nil || s.pushQueue == nil {
		return errors.UninitializedError
	}
	s.pushQueue.elements = append(s.pushQueue.elements, val)
	s.pushQueue.elements = append(s.pushQueue.elements, s.holdQueue.elements...)
	s.holdQueue.Clear()

	tempQ := s.pushQueue
	s.pushQueue = s.holdQueue
	s.holdQueue = tempQ

	s.currSize += 1
	return nil
}

// Pop removes and returns the element on top of stack.
// Returns -1 is stack is un-initialized or empty.
func (s *Stack[T]) Pop() (T, error) {
	var ret T
	if s == nil {
		return ret, errors.UninitializedError
	}
	if s.holdQueue == nil || s.currSize == 0 {
		return ret, errors.UninitializedError
	}
	s.currSize -= 1
	ret = s.holdQueue.elements[0]
	s.holdQueue.elements = s.holdQueue.elements[1:]
	return ret, nil
}

// Pop returns the element on top of stack.
// Returns -1 is stack is un-initialized or empty.
func (s *Stack[T]) Peek() (T, error) {
	var ret T
	if s == nil {
		return ret, errors.UninitializedError
	}
	if s.holdQueue == nil || s.currSize == 0 {
		return ret, errors.UninitializedError
	}
	return s.holdQueue.elements[0], nil
}

// Size returns current size of stack.
func (s *Stack[T]) Size() int {
	return s.currSize
}

// Clear clears the Queue on which it is called.
func (q *Queue[T]) Clear() {
	q.elements = make([]T, 0)
}
