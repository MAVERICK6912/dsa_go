package linkedlist

import (
	"fmt"
	"strings"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
)

type CLLNode[T any] struct {
	data T
	next *CLLNode[T]
}

type CircularLinkedList[T any] struct {
	cmp   func(*CLLNode[T], *CLLNode[T]) int
	first *CLLNode[T]
	last  *CLLNode[T]
	size  int
}

func (c *CircularLinkedList[T]) New(cmp func(*CLLNode[T], *CLLNode[T]) int) *CircularLinkedList[T] {
	return &CircularLinkedList[T]{cmp: cmp}
}

// Add nodes with given data at the end of linkedList.
func (c *CircularLinkedList[T]) Add(n ...T) {
	for _, val := range n {
		node := &CLLNode[T]{data: val}
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

func (c *CircularLinkedList[T]) Insert(index int, val T) error {
	if !c.isWithinRange(index) {
		return errors.IndexOutOfBound
	}
	newNode := &CLLNode[T]{data: val}
	// insertion at starting
	if index == 0 {
		newNode.next = c.first
		c.first = newNode
		c.last.next = c.first
		c.size += 1
		return nil
	}
	// insertion at last
	if index == c.Size() {
		c.last.next = newNode
		c.last = newNode
		c.last.next = c.first
		c.size += 1
		return nil
	}

	node := c.first
	var previousNode *CLLNode[T]
	for i := 0; i != index; i, node = i+1, node.next {
		previousNode = node
	}
	newNode.next = previousNode.next
	previousNode.next = newNode
	c.size += 1
	return nil
}

// Get element at given index.
// Will return -1 if index is out of range or if linkedList is empty.
func (c *CircularLinkedList[T]) Get(index int) (T, error) {
	var ret T
	if !c.isWithinRange(index) {
		return ret, errors.IndexOutOfBound
	}
	// check if list is empty
	if c.first == nil && c.last == nil {
		return ret, errors.UninitializedError
	}
	node := c.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	return node.data, nil
}

// Remove a node with given value `val`.
// Does nothing if val doesn't exist in linkedlist or if linkedlist is empty.
func (c *CircularLinkedList[T]) Remove(val T) error {
	if c.first == nil && c.last == nil {
		return errors.UninitializedError
	}
	remNode := &CLLNode[T]{data: val}
	node := c.first
	var previousNode *CLLNode[T]
	// TODO: add a statement to check if val was not found in linkedlist
	for ; c.cmp(node, remNode) != 0 && node != c.last; node = node.next {
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
	return nil
}

func (c *CircularLinkedList[T]) Delete(index int) error {
	if !c.isWithinRange(index) {
		return errors.IndexOutOfBound
	}
	if c.first == nil && c.last == nil {
		return errors.UninitializedError
	}
	node := c.first
	var previousNode *CLLNode[T]

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
	return nil
}

// Sort a linkedList in ascending order.
func (c *CircularLinkedList[T]) Sort(cmp utils.Comparator[T]) {
	if c.Size() < 2 {
		return
	}
	elems := c.Values()
	utils.Sort(elems, cmp)
	c.Clear()
	c.Add(elems...)
}

// Values returns all the values in a linkedList as a []int.
func (c *CircularLinkedList[T]) Values() []T {
	var tArr []T
	for node := c.first; node != c.last; node = node.next {
		tArr = append(tArr, node.data)
	}
	tArr = append(tArr, c.last.data)

	return tArr
}

// Clear the linkedlist, remove all links.
func (c *CircularLinkedList[T]) Clear() {
	c.first = nil
	c.last = nil
	c.size = 0
}

// isWithinRange checks if the index is within range of linkedList.
func (c *CircularLinkedList[T]) isWithinRange(index int) bool {
	return index >= 0 && index <= c.size
}

// Size returns current size of linkedList
func (c *CircularLinkedList[T]) Size() int {
	return c.size
}

// String implements stringer interface.
func (c *CircularLinkedList[T]) String() string {
	var strArr []string
	for node := c.first; node != c.last; node = node.next {
		strArr = append(strArr, fmt.Sprintf("%v", node.data))
	}
	// appending last item
	strArr = append(strArr, fmt.Sprintf("%v", c.last.data))
	return strings.Join(strArr, ",")
}

func CompareCLLInt(a, b *CLLNode[int]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}

func CompareCLLString(a, b *CLLNode[string]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}
