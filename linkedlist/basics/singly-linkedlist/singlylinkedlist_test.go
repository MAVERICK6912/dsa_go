package linkedlist

import (
	"sort"
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestListAdd(t *testing.T) {
	t.Parallel()
	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist *SinglyLinkedList[int]

	linkedlist = linkedlist.New(CompareSLLInt)
	linkedlist.Add(values...)

	assert.Equal(t, 7, linkedlist.Size())

	_, err := linkedlist.Get(-9)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.IndexOutOfBound)
	}
	actual, err := linkedlist.Get(3)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 2656, actual)
}

func TestLinkedListInsert(t *testing.T) {
	t.Parallel()
	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist *SinglyLinkedList[int]
	linkedlist = linkedlist.New(CompareSLLInt)
	linkedlist.Add(values...)

	err := linkedlist.Insert(2, 43)
	if err != nil {
		t.Error("error while inserting value to linked list")
	}
	actual, err := linkedlist.Get(2)
	if err != nil {
		t.Error("error while getting value from linked list")
	}
	assert.Equal(t, 43, actual)
	assert.Equal(t, 8, linkedlist.Size())

	err = linkedlist.Insert(0, 69)
	if err != nil {
		t.Error("error while inserting value to linked list")
	}
	actual, err = linkedlist.Get(0)
	if err != nil {
		t.Error("error while getting value from linked list")
	}
	assert.Equal(t, 69, actual)
	assert.Equal(t, 9, linkedlist.Size())
}

func TestLinkedListRemove(t *testing.T) {
	t.Parallel()
	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist *SinglyLinkedList[int]
	linkedlist = linkedlist.New(CompareSLLInt)
	linkedlist.Add(values...)

	// remove value 2656 at index 3
	linkedlist.Remove(2656)
	actual, err := linkedlist.Get(3)
	if err != nil {
		t.Error("error while getting value from linked list")
	}
	assert.Equal(t, 48, actual)
	assert.Equal(t, 6, linkedlist.Size())

	// removing first element of linkedList
	linkedlist.Remove(23)
	actual, err = linkedlist.Get(0)
	if err != nil {
		t.Error("error while getting value from linked list")
	}
	assert.Equal(t, values[1], actual)
	assert.Equal(t, 5, linkedlist.Size())

	// removing a node that doesn't exist
	linkedlist.Remove(102)
	assert.Equal(t, 5, linkedlist.Size())

	// testing condition where there's 1 element and we remove it
	linkedlist.Clear()
	linkedlist.Add(69)

	linkedlist.Remove(69)
	_, err = linkedlist.Get(0)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.UninitializedError)
	}
	assert.Equal(t, 0, linkedlist.Size())
}

func TestLinkedListDelete(t *testing.T) {
	t.Parallel()
	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist *SinglyLinkedList[int]
	linkedlist = linkedlist.New(CompareSLLInt)
	linkedlist.Add(values...)

	// deleting node at index 3
	linkedlist.Delete(3)
	actual, err := linkedlist.Get(3)
	if err != nil {
		t.Error("error while getting value from linked list")
	}
	assert.Equal(t, 48, actual)
	assert.Equal(t, 6, linkedlist.Size())

	// deleting first element of linkedList
	linkedlist.Delete(0)
	actual, err = linkedlist.Get(0)
	if err != nil {
		t.Error("error while getting value from linked list")
	}
	assert.Equal(t, 56, actual)
	assert.Equal(t, 5, linkedlist.Size())

	// deleting a node that doesn't exist
	linkedlist.Delete(69)
	assert.Equal(t, 5, linkedlist.Size())
	// testing condition where there's 1 element and we delete it
	linkedlist.Clear()
	linkedlist.Add(69)
	// deleting first(also only node) in linkedlist
	linkedlist.Delete(0)

	_, err = linkedlist.Get(0)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.UninitializedError)
	}
	assert.Equal(t, 0, linkedlist.Size())
}

func TestCircularLinkedListSort(t *testing.T) {
	t.Parallel()

	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var linkedList *SinglyLinkedList[int]
	linkedList = linkedList.New(CompareSLLInt)
	linkedList.Add(values...)

	sort.Ints(values)
	linkedList.Sort(utils.IntComparator)

	assert.Equal(t, values, linkedList.Values())
}
