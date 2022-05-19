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

/*
	n1->n2->n3->n4->n5
	index= 2, nN
	n1->n2->nN->n3->n4->n5
	previousNode= n2
	n2.Next=nN
	nN.Next=n3
	n3.Next=n4
	n4.Next=n5

	findPreviousNode:
		for i:=0; i<index;i,node=i+1,node.next
	We will get previous node in node
	nN.next=previous.next
	previous.next=nN
*/

// Insert a node with given value at given index of linkedList
func (s *SinglyLinkedList) Insert(index, value int) {
	if !s.isWithinRange(index) {
		return
	}
	if index == s.Size() {
		s.Add(value)
		return
	}
	newNode := Node{
		data: value,
	}
	node := s.first
	var previousNode *Node
	for i := 0; i != index; i, node = i+1, node.next {
		previousNode = node
	}
	if node == s.first {
		newNode.next = node
		s.first = &newNode
	} else {
		newNode.next = previousNode.next
		previousNode.next = &newNode
	}
	s.size += 1
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
