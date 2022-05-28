package linkedlist

import (
	"fmt"
	"strings"

	"github.com/maverick6912/dsa_go/utils"
)

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type DoublyLinkedList struct {
	first *Node
	last  *Node
	size  int
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Add nodes to end of linkedList
func (d *DoublyLinkedList) Add(vals ...int) {
	for _, val := range vals {
		node := &Node{data: val}
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
func (d *DoublyLinkedList) Get(index int) int {
	if !d.isWithinRange(index) {
		return -1
	}
	if d.first == nil && d.last == nil {
		return -1
	}

	node := d.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	return node.data
}

// Insert a node with given value at given index of linkedList.
// Does nothing if index is out of bounds of linkedList.
func (d *DoublyLinkedList) Insert(index, val int) {
	if !d.isWithinRange(index) {
		return
	}
	newNode := &Node{data: val}
	node := d.first
	for i := 0; i != index; i, node = i+1, node.next {
	}
	if node == d.first {
		newNode.next = d.first
		d.first.previous = newNode
		d.first = newNode
	} else if node == d.last || index == d.Size() {
		d.Add(val)
		return
	} else {
		newNode.previous = node.previous
		node.previous.next = newNode
		newNode.next = node
		node.previous = newNode
	}
	d.size += 1
}

// Remove a node with given value `val`.
// Does nothing if val doesn't exist in linkedlist or if linkedlist is empty.
func (d *DoublyLinkedList) Remove(val int) {
	if d.Size() == 0 {
		return
	}
	node := d.first
	for ; node.data != val && node != nil; node = node.next {
		if node.data != val && node == d.last {
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
func (d *DoublyLinkedList) Delete(index int) {
	if !d.isWithinRange(index) {
		return
	}
	if d.Size() == 0 {
		return
	}
	node := d.first
	for i := 0; i != index; i, node = i+1, node.next {
		if node == d.last {
			return
		}
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

// Sort a linkedList in ascending order.
func (d *DoublyLinkedList) Sort() {
	var tempArr []int
	node := d.first
	for ; node.next != nil; node = node.next {
	}
	tempArr = append(tempArr, node.data)

	utils.Sort(tempArr, utils.IntComparator)
	d.Clear()
	d.Add(tempArr...)
}

// Clear the linkedlist, remove all links.
func (d *DoublyLinkedList) Clear() {
	d.first = nil
	d.last = nil
	d.size = 0
}

// Values returns all the values in a linkedList as a []int.
func (d *DoublyLinkedList) Values() []int {
	var retArr []int
	node := d.first
	for ; node.next != nil; node = node.next {
		retArr = append(retArr, node.data)
	}
	return retArr
}

// isWithinRange checks if the index is within range of linkedList
func (d *DoublyLinkedList) isWithinRange(index int) bool {
	return index >= 0 && index <= d.size
}

// Get size of LinkedList
func (d *DoublyLinkedList) Size() int {
	return d.size
}

// String implements stringer interface.
func (d *DoublyLinkedList) String() string {
	var strArr []string
	for node := d.first; node != nil; node = node.next {
		strArr = append(strArr, fmt.Sprintf("%d", node.data))
	}
	return strings.Join(strArr, ",")
}
