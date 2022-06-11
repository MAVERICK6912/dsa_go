package stack

import (
	"fmt"
	"strings"

	"github.com/maverick6912/dsa_go/errors"
)

type Stack[T any] struct {
	cmp  func(T, T) int
	list []T
	size int
}

func (s *Stack[T]) New(cmp func(T, T) int) *Stack[T] {
	return &Stack[T]{list: make([]T, 0), cmp: cmp}
}

// Add elements to top of stack
func (s *Stack[T]) Add(val ...T) {
	for _, v := range val {
		s.list = append(s.list, v)
		s.size += 1
	}
}

// Push adds given element to top of stack.
// Does nothing if stack is un-initialized.
func (s *Stack[T]) Push(val T) {
	if s == nil {
		return
	}
	s.list = append(s.list, val)
	s.size += 1
}

// Pop returns and deletes the element on top of stack.
// If stack is empty or un-initialized returns -1
func (s *Stack[T]) Pop() (T, error) {
	var ret T
	if s == nil {
		return ret, errors.UninitializedError
	}
	if s.Size() == 0 || s.list == nil {
		return ret, errors.UninitializedError
	}
	ret = s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	s.size -= 1
	return ret, nil
}

// Peek returns the element on top of stack
func (s *Stack[T]) Peek() (T, error) {
	if s == nil {
		var ret T
		return ret, errors.UninitializedError
	}
	return s.list[len(s.list)-1], nil
}

// Size returns current size of stack
func (s *Stack[T]) Size() int {
	return s.size
}

// String implements stringer interface
func (s *Stack[T]) String() string {
	tArr := make([]string, len(s.list))
	for _, v := range s.list {
		tArr = append(tArr, fmt.Sprintf("%v", v))
	}
	return strings.Join(tArr, ",")
}

// Values returns all elemnts in the stack
func (s *Stack[T]) Values() []T {
	return s.list
}
