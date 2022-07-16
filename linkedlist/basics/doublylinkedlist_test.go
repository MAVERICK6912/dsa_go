package linkedlist

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedListAdd(t *testing.T) {
	t.Parallel()
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll *DoublyLinkedList[int]
	dll = dll.New(CompareDLLInt)
	dll.Add(vals...)

	// Adding random nodes
	actual, err := dll.Get(5)
	if err != nil {
		t.Error("error while getting element from linkedlist")
	}
	assert.Equal(t, 69, actual)
	_, err = dll.Get(9)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.IndexOutOfBound)
	}

	assert.Equal(t, dll.first.next.next.data, dll.first.next.next.next.previous.data)

	assert.Equal(t, len(vals), dll.Size())
}

func TestDoublyLinkedListInsert(t *testing.T) {
	t.Parallel()
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll *DoublyLinkedList[int]
	dll = dll.New(CompareDLLInt)

	dll.Add(vals...)
	// inserting at 2nd index
	err := dll.Insert(2, 9)
	if err != nil {
		t.Error("error while inserting to linked list")
	}
	actual, err := dll.Get(2)
	if err != nil {
		t.Error("error while getting from linked list")
	}
	assert.Equal(t, 9, actual)
	assert.Equal(t, len(vals)+1, dll.Size())
	// inserting a head node
	firstNode := dll.first
	err = dll.Insert(0, 123)
	if err != nil {
		t.Error("error while inserting to linked list")
	}
	actual, err = dll.Get(0)
	if err != nil {
		t.Error("error while getting from linked list")
	}
	assert.Equal(t, 123, actual)
	assert.Equal(t, firstNode.data, dll.first.next.data)
	assert.Equal(t, len(vals)+2, dll.Size())
	// inserting node at last
	lastNode := dll.last
	err = dll.Insert(dll.Size(), 852)
	if err != nil {
		t.Error("error while inserting to linked list")
	}
	actual, err = dll.Get(dll.Size() - 1)
	if err != nil {
		t.Error("error while getting from linked list")
	}
	assert.Equal(t, 852, actual)
	assert.Equal(t, dll.last.previous.data, lastNode.data)
	assert.Equal(t, len(vals)+3, dll.Size())
}

func TestDoublyLinkedListRemove(t *testing.T) {
	t.Parallel()
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll DoublyLinkedList[int]
	dll = *dll.New(CompareDLLInt)
	dll.Add(vals...)
	// removing random node in linkedList
	dll.Remove(45)
	actual, err := dll.Get(3)
	if err != nil {
		t.Error("error while getting from linked list")
	}
	assert.NotEqual(t, vals[3], actual)
	assert.Equal(t, len(vals)-1, dll.Size())
	// removing head node
	dll.Remove(59)
	actual, err = dll.Get(0)
	if err != nil {
		t.Error("error while getting from linked list")
	}
	assert.NotEqual(t, vals[0], actual)
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
	t.Parallel()
	vals := []int{59, 3, 6, 45, 98, 69, 201}
	var dll DoublyLinkedList[int]
	dll = *dll.New(CompareDLLInt)
	dll.Add(vals...)
	// deleting element at index 3
	err := dll.Delete(3)
	if err != nil {
		t.Error("error while deleting in linkedlist")
	}
	actual, err := dll.Get(3)
	if err != nil {
		t.Error("error while getting from linkedlist")
	}
	assert.NotEqual(t, vals[3], actual)
	assert.Equal(t, len(vals)-1, dll.Size())
	// deleting head node
	err = dll.Delete(0)
	if err != nil {
		t.Error("error while deleting in linkedlist")
	}
	actual, err = dll.Get(0)
	if err != nil {
		t.Error("error while getting from linkedlist")
	}
	assert.NotEqual(t, vals[0], actual)
	assert.Equal(t, len(vals)-2, dll.Size())
	// deleting last node
	dll.Delete(len(vals) - 3)
	assert.NotEqual(t, vals[len(vals)-1], dll.last.data)
	assert.Equal(t, len(vals)-3, dll.Size())
	// deleting a node not in linkedList
	err = dll.Delete(57)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.IndexOutOfBound)
	}
	assert.Equal(t, len(vals)-3, dll.Size())
}
