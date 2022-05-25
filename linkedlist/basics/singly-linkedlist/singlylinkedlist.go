package linkedlist

import (
	"fmt"
	"strings"

	utils "github.com/maverick6912/dsa_go/utils"
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
	// checking if list is empty
	if s.first == nil && s.last == nil {
		return -1
	}
	return node.data
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
func (s *SinglyLinkedList) Remove(val int) {
	node := s.first
	var previousNode *Node
	for i := 0; node != nil; i, node = i+1, node.next {
		if node.data == val {
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
func (s *SinglyLinkedList) Delete(index int) {
	if !s.isWithinRange(index) {
		return
	}
	node := s.first
	var previousNode *Node
	for i := 0; i != index; i, node = i+1, node.next {
		previousNode = node
	}
	if node == s.first {
		if s.size == 1 {
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

// Sort a linkedList in ascending order.
func (s *SinglyLinkedList) Sort() {
	if s.Size() < 2 {
		return
	}
	elems := s.Values()
	utils.Sort(elems, utils.IntComparator)
	s.Clear()
	s.Add(elems...)
}

// Values returns all the values in a linkedList as a []int.
func (s *SinglyLinkedList) Values() []int {
	var tArr []int
	for node := s.first; node != nil; node = node.next {
		tArr = append(tArr, node.data)
	}
	return tArr
}

// Clear the linkedlist, remove all links.
func (s *SinglyLinkedList) Clear() {
	s.first = nil
	s.last = nil
	s.size = 0
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
