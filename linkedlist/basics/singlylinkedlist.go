package linkedlist

import (
	"fmt"
	"strings"

	"github.com/maverick6912/dsa_go/errors"
	utils "github.com/maverick6912/dsa_go/utils"
)

type SLLNode[T any] struct {
	data T
	next *SLLNode[T]
}

type SinglyLinkedList[T any] struct {
	cmp   func(*SLLNode[T], *SLLNode[T]) int
	first *SLLNode[T]
	last  *SLLNode[T]
	size  int
}

// New initializes the SinglyLinkedist on which it is called upon and returns it.
func (s *SinglyLinkedList[T]) New(cmp func(*SLLNode[T], *SLLNode[T]) int) *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{cmp: cmp}
}

// Add values at the end of linkedList
func (s *SinglyLinkedList[T]) Add(n ...T) {
	for _, v := range n {
		node := &SLLNode[T]{data: v}
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
func (s *SinglyLinkedList[T]) Size() int {
	return s.size
}

// Get value at index.
// If index is out of range of linkedlist then function returns -1
func (s *SinglyLinkedList[T]) Get(index int) (T, error) {
	var ret T
	if !s.isWithinRange(index) {
		return ret, errors.IndexOutOfBound
	}
	node := s.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	// checking if list is empty
	if s.first == nil && s.last == nil {
		return ret, errors.UninitializedError
	}
	ret = node.data
	return ret, nil
}

/*
	Approach for insert:
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
// Insert a node with given value at given index of linkedList.
// Does nothing if index is out of bounds of linkedList.
func (s *SinglyLinkedList[T]) Insert(index int, value T) error {
	if !s.isWithinRange(index) {
		return errors.IndexOutOfBound
	}
	if index == s.Size() {
		s.Add(value)
		return nil
	}
	newNode := SLLNode[T]{
		data: value,
	}
	node := s.first
	var previousNode *SLLNode[T]
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
	return nil
}

/*
	Approach for remove:
	n1->n2->n3->n4->n5
		Remove(n3)
		n1->n2->n4->n5
		currentNode=n3
		previousNode= n2
		n2.Next=n3.next

		findPreviousNode:
			for i:=0; i<index;i,node=i+1,node.next
		We will get previous node in node
		previousNode.next=currentNode.next

		currentNode=nil
*/
// Remove a node with given value `val`
// Does nothing if val doesn't exist in linkedlist or if linkedlist is empty
func (s *SinglyLinkedList[T]) Remove(val T) {
	node := s.first
	var previousNode *SLLNode[T]
	compNode := &SLLNode[T]{data: val}
	for i := 0; node != nil; i, node = i+1, node.next {
		if s.cmp(node, compNode) == 0 {
			break
		}
		previousNode = node
	}
	// this means val was not found in any node
	if node == nil {
		return
	}
	if node == s.first {
		if s.Size() == 1 {
			s.Clear()
			return
		}
		s.first = node.next
	}
	if node == s.last {
		s.last = previousNode
	}
	if previousNode != nil {
		previousNode.next = node.next
	}
	node = nil
	s.size -= 1
}

/*
	Approach for delete:
	n1->n2->n3->n4->n5
		Delete(index: 2)
		n1->n2->n4->n5
		currentNode=n3
		previousNode= n2
		n2.Next=n3.next

		findPreviousNode:
			for i:=0; i<index;i,node=i+1,node.next
		We will get previous node in node
		previousNode.next=currentNode.next

		currentNode=nil
*/
// Delete a node at given index.
// Does nothing if index is not within range of linkedlist or if linkedlist is empty
func (s *SinglyLinkedList[T]) Delete(index int) error {
	if !s.isWithinRange(index) {
		return errors.IndexOutOfBound
	}
	node := s.first
	var previousNode *SLLNode[T]
	for i := 0; i != index; i, node = i+1, node.next {
		previousNode = node
	}
	if node == s.first {
		if s.size == 1 {
			s.Clear()
			return nil
		}
		s.first = node.next
	}
	if node == s.last {
		s.last = previousNode
	}
	if previousNode != nil {
		previousNode.next = node.next
	}
	node = nil
	s.size -= 1
	return nil
}

// Sort a linkedList in ascending order.
func (s *SinglyLinkedList[T]) Sort(cmp utils.Comparator[T]) {
	if s.Size() < 2 {
		return
	}
	elems := s.Values()
	utils.Sort(elems, cmp)
	s.Clear()
	s.Add(elems...)
}

func (s *SinglyLinkedList[T]) Values() []T {
	var ret []T

	for node := s.first; node != nil; node = node.next {
		ret = append(ret, node.data)
	}
	return ret
}

// Clear the linkedlist, remove all links.
func (s *SinglyLinkedList[T]) Clear() {
	s.first = nil
	s.last = nil
	s.size = 0
}

// String implements stringer interface.
func (s *SinglyLinkedList[T]) String() string {
	var strArr []string
	for node := s.first; node != nil; node = node.next {
		strArr = append(strArr, fmt.Sprintf("%v", node.data))
	}
	return strings.Join(strArr, ",")
}

// isWithinRange checks if the index is within range of linkedList
func (s *SinglyLinkedList[T]) isWithinRange(index int) bool {
	return index >= 0 && index <= s.size
}

func CompareSLLInt(a, b *SLLNode[int]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}

func CompareSLLString(a, b *SLLNode[string]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}
