package linkedlist

import "testing"

func TestListInsert(t *testing.T) {

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
