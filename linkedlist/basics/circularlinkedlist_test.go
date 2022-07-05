package linkedlist

import (
	"sort"
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestCircularLinkedListAdd(t *testing.T) {
	t.Parallel()
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}

	var cirLinkedList *CircularLinkedList[int]
	cirLinkedList = cirLinkedList.New(CompareCLLInt)
	cirLinkedList.Add(values...)

	actual, err := cirLinkedList.Get(4)
	if err != nil {
		t.Error("error while getting value from queue. error was: ", err.Error())
	}
	assert.Equal(t, 14, actual)
	assert.Equal(t, cirLinkedList.first, cirLinkedList.last.next)
	assert.Equal(t, len(values), cirLinkedList.Size())
}

func TestCircularLinkedListInsert(t *testing.T) {
	t.Parallel()

	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList *CircularLinkedList[int]
	cirLinkedList = cirLinkedList.New(CompareCLLInt)
	cirLinkedList.Add(values...)

	// inserting between linkedList
	err := cirLinkedList.Insert(3, 321)
	if err != nil {
		t.Error("error while inserting to queue")
	}
	actual, err := cirLinkedList.Get(3)
	if err != nil {
		t.Error("error while getting from linkedlist")
	}
	assert.Equal(t, 321, actual)
	assert.Equal(t, 9, cirLinkedList.Size())

	// inserting at start
	err = cirLinkedList.Insert(0, 634)
	if err != nil {
		t.Error("error while inserting to queue")
	}
	actual, _ = cirLinkedList.Get(0)
	assert.Equal(t, 634, actual)
	assert.Equal(t, 10, cirLinkedList.Size())
	assert.Equal(t, cirLinkedList.last.next, cirLinkedList.first)
	assert.Equal(t, 634, cirLinkedList.first.data)

	// inserting at the end
	err = cirLinkedList.Insert(cirLinkedList.Size(), 999)
	if err != nil {
		t.Error("error while inserting to queue")
	}
	actual, err = cirLinkedList.Get(10)
	if err != nil {
		t.Error("error while getting from linkedlist")
	}
	assert.Equal(t, 999, actual)
	assert.Equal(t, 11, cirLinkedList.Size())
}

func TestCircularLinkedListSort(t *testing.T) {
	t.Parallel()

	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList *CircularLinkedList[int]
	cirLinkedList = cirLinkedList.New(CompareCLLInt)
	cirLinkedList.Add(values...)

	sort.Ints(values)
	cirLinkedList.Sort(utils.IntComparator)

	assert.Equal(t, values, cirLinkedList.Values())
}

func TestCircularLinkedListRemove(t *testing.T) {
	t.Parallel()

	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList *CircularLinkedList[int]
	cirLinkedList = cirLinkedList.New(CompareCLLInt)
	cirLinkedList.Add(values...)

	// removing randomly
	err := cirLinkedList.Remove(32)
	if err != nil {
		t.Error("error while removing element from linkedlist. error was: ", err)
	}
	actual, err := cirLinkedList.Get(2)
	if err != nil {
		t.Error("error while getting element from linkedlist. error was: ", err)
	}
	assert.NotEqual(t, values[2], actual)
	actual, err = cirLinkedList.Get(2)
	if err != nil {
		t.Error("error while getting element from linkedlist. error was: ", err)
	}
	assert.Equal(t, values[3], actual)
	assert.Equal(t, len(values)-1, cirLinkedList.Size())

	// removing head node
	err = cirLinkedList.Remove(25)
	if err != nil {
		t.Error("error while removing element from linkedlist. error was: ", err)
	}
	actual, err = cirLinkedList.Get(0)
	if err != nil {
		t.Error("error while getting element from linkedlist. error was: ", err)
	}
	assert.NotEqual(t, values[0], actual)
	assert.Equal(t, values[1], actual)
	assert.Equal(t, len(values)-2, cirLinkedList.Size())

	// removing last node
	err = cirLinkedList.Remove(536)
	if err != nil {
		t.Error("error while removing element from linkedlist. error was: ", err)
	}
	assert.NotEqual(t, values[len(values)-4], cirLinkedList.last.data)
	assert.Equal(t, values[len(values)-2], cirLinkedList.last.data)
	assert.Equal(t, len(values)-3, cirLinkedList.Size())

	// removing a node not in linkedList
	// err = cirLinkedList.Remove(488)
	// if assert.Error(t, err) {
	// 	assert.ErrorIs(t, err, errors.UninitializedError)
	// }
	// assert.Equal(t, len(values)-4, cirLinkedList.Size())
}

func TestCircularLinkedListDelete(t *testing.T) {
	t.Parallel()

	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList *CircularLinkedList[int]
	cirLinkedList = cirLinkedList.New(CompareCLLInt)
	cirLinkedList.Add(values...)

	// deleting randomly
	err := cirLinkedList.Delete(2)
	if err != nil {
		t.Error("error while deleting element from linkedlist. error was: ", err.Error())
	}
	actual, err := cirLinkedList.Get(2)
	if err != nil {
		t.Error("error while getting element from linkedlist. error was: ", err.Error())
	}
	assert.NotEqual(t, values[2], actual)
	assert.Equal(t, values[3], actual)
	assert.Equal(t, len(values)-1, cirLinkedList.Size())

	// deleting head node
	err = cirLinkedList.Delete(0)
	if err != nil {
		t.Error("error while deleting element from linkedlist. error was: ", err.Error())
	}
	actual, err = cirLinkedList.Get(0)
	if err != nil {
		t.Error("error while getting element from linkedlist. error was: ", err.Error())
	}
	assert.NotEqual(t, values[0], actual)
	assert.Equal(t, values[1], actual)
	assert.Equal(t, len(values)-2, cirLinkedList.Size())

	// deleting last node
	err = cirLinkedList.Delete(len(values) - 3)
	if err != nil {
		t.Error("error while deleting element from linkedlist. error was: ", err.Error())
	}
	assert.NotEqual(t, values[len(values)-4], cirLinkedList.last.data)
	assert.Equal(t, values[len(values)-2], cirLinkedList.last.data)
	assert.Equal(t, len(values)-3, cirLinkedList.Size())

	// deleting a node not in linkedList
	err = cirLinkedList.Delete(488)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.IndexOutOfBound)
	}
	assert.Equal(t, len(values)-3, cirLinkedList.Size())
}
