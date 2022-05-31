package queue

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

type Stack struct {
	pushQueue *Queue
	holdQueue *Queue
	currSize  int
}

// TODO: add Enque and Dequeue ops for queue here.
type Queue struct {
	elements []int
	// size     int
}

// New initializes a `Stack`
func New() *Stack {
	return &Stack{pushQueue: newQueue(), holdQueue: newQueue()}
}

// newQueue initializes a `Queue`.
// Only for internal use.
func newQueue() *Queue {
	return &Queue{elements: make([]int, 0)}
}

// Push adds given value to the stack.
// Doesn't do anything if stack is un-initialized.
func (s *Stack) Push(val int) {
	if s == nil {
		return
	}
	if s.holdQueue == nil || s.pushQueue == nil {
		return
	}
	s.pushQueue.elements = append(s.pushQueue.elements, val)
	s.pushQueue.elements = append(s.pushQueue.elements, s.holdQueue.elements...)
	s.holdQueue.Clear()

	tempQ := s.pushQueue
	s.pushQueue = s.holdQueue
	s.holdQueue = tempQ

	s.currSize += 1
}

// Pop removes and returns the element on top of stack.
// Returns -1 is stack is un-initialized or empty.
func (s *Stack) Pop() int {
	if s == nil {
		return -1
	}
	if s.holdQueue == nil || s.currSize == 0 {
		return -1
	}
	s.currSize -= 1
	val := s.holdQueue.elements[0]
	s.holdQueue.elements = s.holdQueue.elements[1:]
	return val
}

// Pop returns the element on top of stack.
// Returns -1 is stack is un-initialized or empty.
func (s *Stack) Peek() int {
	if s == nil {
		return -1
	}
	if s.holdQueue == nil || s.currSize == 0 {
		return -1
	}
	return s.holdQueue.elements[0]
}

// Size returns current size of stack.
func (s *Stack) Size() int {
	return s.currSize
}

// Clear clears the Queue on which it is called.
func (q *Queue) Clear() {
	q.elements = make([]int, 0)
}
