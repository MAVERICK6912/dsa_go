package linkedlist

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAdd(t *testing.T) {

	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist SinglyLinkedList
	linkedlist.Add(values...)

	if actual := linkedlist.Size(); actual != 7 {
		t.Errorf("Got %v expected %v", actual, 7)
	}
	if actual := linkedlist.Get(-9); actual != -1 {
		t.Errorf("Got %v expected %v", actual, -1)
	}
	if actual := linkedlist.Get(3); actual != 2656 {
		t.Errorf("Got %v expected %v", actual, 2656)
	}
}

func TestLinkedListInsert(t *testing.T) {
	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist SinglyLinkedList
	linkedlist.Add(values...)

	linkedlist.Insert(2, 0)
	if actual := linkedlist.Get(2); actual != 0 {
		t.Errorf("Got %v expected %v", actual, 0)
	}
	if actual := linkedlist.Size(); actual != 8 {
		t.Errorf("Got %v expected %v", actual, 8)
	}

	linkedlist.Insert(0, 69)

	if actual := linkedlist.Get(0); actual != 69 {
		t.Errorf("Got %v expected %v", actual, 69)
	}
	if actual := linkedlist.Size(); actual != 9 {
		t.Errorf("Got %v expected %v", actual, 9)
	}
}

func TestLinkedListRemove(t *testing.T) {
	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist SinglyLinkedList
	linkedlist.Add(values...)

	// remove value 2656 at index 3
	linkedlist.Remove(2656)
	if actual := linkedlist.Get(3); actual != 48 {
		t.Errorf("Got %v expected %v", actual, 48)
	}
	if actual := linkedlist.Size(); actual != 6 {
		t.Errorf("Got %v expected %v", actual, 6)
	}

	// removing first element of linkedList
	linkedlist.Remove(23)

	if actual := linkedlist.Get(0); actual != 56 {
		t.Errorf("Got %v expected %v", actual, 56)
	}
	if actual := linkedlist.Size(); actual != 5 {
		t.Errorf("Got %v expected %v", actual, 5)
	}

	// removing a node that doesn't exist
	linkedlist.Remove(102)
	if actual := linkedlist.Size(); actual != 5 {
		t.Errorf("Got %v expected %v", actual, 5)
	}

	// testing condition where there's 1 element and we remove it
	linkedlist.Clear()
	linkedlist.Add(69)

	linkedlist.Remove(69)

	if actual := linkedlist.Get(0); actual != -1 {
		t.Errorf("Got %v expected %v", actual, -1)
	}
	if actual := linkedlist.Size(); actual != 0 {
		t.Errorf("Got %v expected %v", actual, 0)
	}
}

func TestLinkedListDelete(t *testing.T) {
	values := []int{23, 56, 14, 2656, 48, 62, 68}

	var linkedlist SinglyLinkedList
	linkedlist.Add(values...)

	// deleting node and index 3
	linkedlist.Delete(3)
	if actual := linkedlist.Get(3); actual != 48 {
		t.Errorf("Got %v expected %v", actual, 48)
	}
	if actual := linkedlist.Size(); actual != 6 {
		t.Errorf("Got %v expected %v", actual, 6)
	}

	// deleting first element of linkedList
	linkedlist.Delete(0)

	if actual := linkedlist.Get(0); actual != 56 {
		t.Errorf("Got %v expected %v", actual, 56)
	}
	if actual := linkedlist.Size(); actual != 5 {
		t.Errorf("Got %v expected %v", actual, 5)
	}

	// deleting a node that doesn't exist
	linkedlist.Delete(69)
	if actual := linkedlist.Size(); actual != 5 {
		t.Errorf("Got %v expected %v", actual, 5)
	}

	// testing condition where there's 1 element and we delete it
	linkedlist.Clear()
	linkedlist.Add(69)
	// deleting first(also only node) in linkedlist
	linkedlist.Delete(0)

	if actual := linkedlist.Get(0); actual != -1 {
		t.Errorf("Got %v expected %v", actual, -1)
	}
	if actual := linkedlist.Size(); actual != 0 {
		t.Errorf("Got %v expected %v", actual, 0)
	}
}

func TestCircularLinkedListSort(t *testing.T) {
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var linkedList SinglyLinkedList
	linkedList.Add(values...)

	sort.Ints(values)
	linkedList.Sort()

	assert.Equal(t, values, linkedList.Values())
}
