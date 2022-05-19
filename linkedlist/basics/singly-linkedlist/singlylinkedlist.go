package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	first *Node
	last  *Node
	size  int
}

// Add values at the end of linkedList
func (s *SinglyLinkedList) Add(n ...int) {
	for _, v := range n {
		node := &Node{data: v}
		if s.size == 0 {
			s.first = node
			s.last = node
		} else {
			s.last.next = node
			s.last = node
		}
		s.size += 1
	}
}

// Get size of LinkedList
func (s *SinglyLinkedList) Size() int {
	return s.size
}

// Get value at index.
// If index is out of range of linkedlist then function returns -1
func (s *SinglyLinkedList) Get(index int) int {
	if !s.isWithinRange(index) {
		return -1
	}
	node := s.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	return node.data
}

// String implements stringer interface.
func (s *SinglyLinkedList) String() string {
	var strArr []string
	for node := s.first; node != nil; node = node.next {
		strArr = append(strArr, fmt.Sprintf("%d", node.data))
	}
	return strings.Join(strArr, ",")
}

// isWithinRange checks if the index is within range of linkedList
func (s *SinglyLinkedList) isWithinRange(index int) bool {
	return index >= 0 && index <= s.size
}
