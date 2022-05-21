package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	first *Node
	last  *Node
	size  int
}

// Add nodes with given data at the end of linkedList.
func (c *CircularLinkedList) Add(n ...int) {
	for _, val := range n {
		node := &Node{data: val}
		if c.size == 0 {
			c.first = node
			c.last = node
		} else {
			c.last.next = node
			c.last = node
			// making linkedlist circular
			c.last.next = c.first
		}
		c.size += 1
	}
}

// Get element at given index.
// Will return -1 if index is out of range or if linkedList is empty.
func (c *CircularLinkedList) Get(index int) int {
	if !c.isWithinRange(index) {
		return -1
	}
	// check if list is empty
	if c.first == nil && c.last == nil {
		return -1
	}
	node := c.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	return node.data
}

// isWithinRange checks if the index is within range of linkedList.
func (c *CircularLinkedList) isWithinRange(index int) bool {
	return index >= 0 && index <= c.size
}

// Size returns current size of linkedList
func (c *CircularLinkedList) Size() int {
	return c.size
}

// String implements stringer interface.
func (c *CircularLinkedList) String() string {
	var strArr []string
	for node := c.first; node != c.last; node = node.next {
		strArr = append(strArr, fmt.Sprintf("%d", node.data))
	}
	// appending last item
	strArr = append(strArr, fmt.Sprintf("%d", c.last.data))
	return strings.Join(strArr, ",")
}
