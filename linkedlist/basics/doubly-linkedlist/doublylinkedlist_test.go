package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedListAdd(t *testing.T) {
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll DoublyLinkedList

	dll.Add(vals...)

	// Adding random nodes
	assert.Equal(t, 69, dll.Get(5))

	assert.Equal(t, -1, dll.Get(9))

	assert.Equal(t, dll.first.next.next.data, dll.first.next.next.next.previous.data)

	assert.Equal(t, len(vals), dll.Size())
}

func TestDoublyLinkedListInsert(t *testing.T) {
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll DoublyLinkedList

	dll.Add(vals...)
	// inserting at 2nd index
	dll.Insert(2, 9)
	assert.Equal(t, 9, dll.Get(2))
	assert.Equal(t, len(vals)+1, dll.Size())
	// inserting a head node
	firstNode := dll.first
	dll.Insert(0, 123)
	assert.Equal(t, 123, dll.Get(0))
	assert.Equal(t, firstNode.data, dll.first.next.data)
	assert.Equal(t, len(vals)+2, dll.Size())
	// inserting node at last
	lastNode := dll.last
	dll.Insert(dll.Size(), 852)
	assert.Equal(t, dll.Get(dll.Size()-1), 852)
	assert.Equal(t, dll.last.previous.data, lastNode.data)
	assert.Equal(t, len(vals)+3, dll.Size())
}

func TestDoublyLinkedListRemove(t *testing.T) {
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll DoublyLinkedList

	dll.Add(vals...)
	// removing random node in linkedList
	dll.Remove(45)
	assert.NotEqual(t, vals[3], dll.Get(3))
	assert.Equal(t, len(vals)-1, dll.Size())
	// removing head node
	dll.Remove(59)
	assert.NotEqual(t, vals[0], dll.Get(0))
	assert.Equal(t, len(vals)-2, dll.Size())
	// removing last node
	dll.Remove(201)
	assert.NotEqual(t, vals[len(vals)-1], dll.last.data)
	assert.Equal(t, len(vals)-3, dll.Size())
	// removing a node not in linkedList
	dll.Remove(2358)
	assert.Equal(t, len(vals)-3, dll.Size())
}

func TestDoublyLinkedListDelete(t *testing.T) {
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll DoublyLinkedList

	dll.Add(vals...)
	// deleting element at index 3
	dll.Delete(3)
	assert.NotEqual(t, vals[3], dll.Get(3))
	assert.Equal(t, len(vals)-1, dll.Size())
	// deleting head node
	dll.Delete(0)
	assert.NotEqual(t, vals[0], dll.Get(0))
	assert.Equal(t, len(vals)-2, dll.Size())
	// deleting last node
	dll.Delete(len(vals) - 3)
	assert.NotEqual(t, vals[len(vals)-1], dll.last.data)
	assert.Equal(t, len(vals)-3, dll.Size())
	// deleting a node not in linkedList
	dll.Delete(57)
	assert.Equal(t, len(vals)-3, dll.Size())
}
