package linkedlist

import (
	"testing"
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
