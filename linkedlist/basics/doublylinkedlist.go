package linkedlist

import (
	"fmt"
	"strings"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
)

type DLLNode[T any] struct {
	data     T
	next     *DLLNode[T]
	previous *DLLNode[T]
}

type DoublyLinkedList[T any] struct {
	cmp   func(*DLLNode[T], *DLLNode[T]) int
	first *DLLNode[T]
	last  *DLLNode[T]
	size  int
}

func (d *DoublyLinkedList[T]) New(cmp func(*DLLNode[T], *DLLNode[T]) int) *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{cmp: cmp}
}

// Add nodes to end of linkedList
func (d *DoublyLinkedList[T]) Add(vals ...T) {
	for _, val := range vals {
		node := &DLLNode[T]{data: val}
		if d.Size() == 0 && d.first == nil {
			d.first = node
			d.last = node
		} else {
			d.last.next = node
			node.previous = d.last
			d.last = node
		}
		d.size += 1
	}
}

// Get value at index.
// If index is out of range of linkedlist then function returns -1
func (d *DoublyLinkedList[T]) Get(index int) (T, error) {
	var ret T
	if !d.isWithinRange(index) {
		return ret, errors.IndexOutOfBound
	}
	if d.first == nil && d.last == nil {
		return ret, errors.UninitializedError
	}

	node := d.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	ret = node.data
	return ret, nil
}

// Insert a node with given value at given index of linkedList.
// Does nothing if index is out of bounds of linkedList.
func (d *DoublyLinkedList[T]) Insert(index int, val T) error {
	if !d.isWithinRange(index) {
		return errors.IndexOutOfBound
	}
	newNode := &DLLNode[T]{data: val}
	node := d.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	if node == d.first {
		newNode.next = d.first
		d.first.previous = newNode
		d.first = newNode
	} else if node == d.last || index == d.Size() {
		d.Add(val)
		return nil
	} else {
		newNode.previous = node.previous
		node.previous.next = newNode
		newNode.next = node
		node.previous = newNode
	}
	d.size += 1
	return nil
}

// Remove a node with given value `val`.
// Does nothing if val doesn't exist in linkedlist or if linkedlist is empty.
func (d *DoublyLinkedList[T]) Remove(val T) {
	if d.Size() == 0 {
		return
	}
	cmpNode := &DLLNode[T]{data: val}
	node := d.first
	for ; d.cmp(node, cmpNode) != 0 && node != nil; node = node.next {
		if d.cmp(node, cmpNode) != 0 && node == d.last {
			return
		}
	}
	if node == nil {
		return
	}
	if node == d.first {
		d.first = node.next
		node.next = nil
		d.size -= 1
		return
	} else if node == d.last {
		d.last = node.previous
		node.previous = nil
		d.size -= 1
		return
	}
	node.previous.next = node.next
	node.next.previous = node.previous
	d.size -= 1
}

// Delete a node at given index.
// Does nothing if index is not within range of linkedlist or if linkedlist is empty.
func (d *DoublyLinkedList[T]) Delete(index int) error {
	if !d.isWithinRange(index) {
		return errors.IndexOutOfBound
	}
	if d.Size() == 0 {
		return errors.UninitializedError
	}
	node := d.first
	for i := 0; i != index; i, node = i+1, node.next {
		if node == d.last {
			return errors.NotFound
		}
	}
	if node == d.first {
		d.first = node.next
		node.next = nil
		d.size -= 1
		return nil
	} else if node == d.last {
		d.last = node.previous
		node.previous = nil
		d.size -= 1
		return nil
	}
	node.previous.next = node.next
	node.next.previous = node.previous
	d.size -= 1
	return nil
}

// Sort a linkedList in ascending order.
func (d *DoublyLinkedList[T]) Sort(cmp utils.Comparator[T]) {
	tempArr := d.Values()

	utils.Sort(tempArr, cmp)
	d.Clear()
	d.Add(tempArr...)
}

// Clear the linkedlist, remove all links.
func (d *DoublyLinkedList[T]) Clear() {
	d.first = nil
	d.last = nil
	d.size = 0
}

// Values returns all the values in a linkedList as a []int.
func (d *DoublyLinkedList[T]) Values() []T {
	var retArr []T
	node := d.first
	for ; node.next != nil; node = node.next {
		retArr = append(retArr, node.data)
	}
	return retArr
}

// isWithinRange checks if the index is within range of linkedList
func (d *DoublyLinkedList[T]) isWithinRange(index int) bool {
	return index >= 0 && index <= d.size
}

// Get size of LinkedList
func (d *DoublyLinkedList[T]) Size() int {
	return d.size
}

// String implements stringer interface.
func (d *DoublyLinkedList[T]) String() string {
	var strArr []string
	for node := d.first; node != nil; node = node.next {
		strArr = append(strArr, fmt.Sprintf("%v", node.data))
	}
	return strings.Join(strArr, ",")
}

func CompareDLLInt(a, b *DLLNode[int]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}
func CompareDLLString(a, b *DLLNode[string]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}
