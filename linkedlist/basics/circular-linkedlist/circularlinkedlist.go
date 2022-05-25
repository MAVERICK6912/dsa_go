package linkedlist

import (
	"fmt"
	"strings"

	"github.com/maverick6912/dsa_go/utils"
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

func (c *CircularLinkedList) Insert(index, val int) {
	if !c.isWithinRange(index) {
		return
	}
	newNode := &Node{data: val}
	// insertion at starting
	if index == 0 {
		newNode.next = c.first
		c.first = newNode
		c.last.next = c.first
		c.size += 1
		return
	}
	// insertion at last
	if index == c.Size() {
		c.last.next = newNode
		c.last = newNode
		c.last.next = c.first
		c.size += 1
		return
	}

	node := c.first
	var previousNode *Node
	for i := 0; i != index; i, node = i+1, node.next {
		previousNode = node
	}
	newNode.next = previousNode.next
	previousNode.next = newNode
	c.size += 1
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

// Remove a node with given value `val`.
// Does nothing if val doesn't exist in linkedlist or if linkedlist is empty.
func (c *CircularLinkedList) Remove(val int) {
	if c.first == nil && c.last == nil {
		return
	}

	node := c.first
	var previousNode *Node

	for ; node.data != val && node != c.last; node = node.next {
		previousNode = node
	}

	if node == c.first {
		c.first = node.next
		c.last.next = c.first
		node = nil
	} else if node == c.last {
		previousNode.next = c.first
		c.last = previousNode
		node = nil
	} else {
		previousNode.next = node.next
		node = nil
	}
	c.size -= 1
}

func (c *CircularLinkedList) Delete(index int) {
	if !c.isWithinRange(index) {
		return
	}
	if c.first == nil && c.last == nil {
		return
	}
	node := c.first
	var previousNode *Node

	for i := 0; i != index && node != c.last; i, node = i+1, node.next {
		previousNode = node
	}
	if node == c.first {
		c.first = node.next
		c.last.next = c.first
		node = nil
	} else if node == c.last {
		c.last = previousNode
		c.last.next = c.first
		node = nil
	} else {
		previousNode.next = node.next
		node = nil
	}
	c.size -= 1
}

// Sort a linkedList in ascending order.
func (c *CircularLinkedList) Sort() {
	if c.Size() < 2 {
		return
	}
	elems := c.Values()
	utils.Sort(elems, utils.IntComparator)
	c.Clear()
	c.Add(elems...)
}

// Values returns all the values in a linkedList as a []int.
func (c *CircularLinkedList) Values() []int {
	var tArr []int
	for node := c.first; node != c.last; node = node.next {
		tArr = append(tArr, node.data)
	}
	tArr = append(tArr, c.last.data)

	return tArr
}

// Clear the linkedlist, remove all links.
func (c *CircularLinkedList) Clear() {
	c.first = nil
	c.last = nil
	c.size = 0
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
