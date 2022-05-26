package stack

import (
	"fmt"
	"strings"
)

type Stack struct {
	list []int
	size int
}

// Add elements to top of stack
func (s *Stack) Add(val ...int) {
	for _, v := range val {
		s.list = append(s.list, v)
		s.size += 1
	}
}

// Push adds given element to top of stack
func (s *Stack) Push(val int) {
	s.list = append(s.list, val)
	s.size += 1
}

// Pop returns and deletes the element on top of stack
func (s *Stack) Pop() int {
	if s.Size() == 0 {
		return -1
	}
	elem := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	s.size -= 1
	return elem
}

// Peek returns the element on top of stack
func (s *Stack) Peek() int {
	return s.list[len(s.list)-1]
}

// Size returns current size of stack
func (s *Stack) Size() int {
	return s.size
}

// String implements stringer interface
func (s *Stack) String() string {
	tArr := make([]string, len(s.list))
	for _, v := range s.list {
		tArr = append(tArr, fmt.Sprintf("%d", v))
	}
	return strings.Join(tArr, ",")
}

// Values returns all elemnts in the stack
func (s *Stack) Values() []int {
	return s.list
}
