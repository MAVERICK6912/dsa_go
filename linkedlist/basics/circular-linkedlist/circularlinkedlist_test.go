package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

func TestCircularLinkedListAdd(t *testing.T) {
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList CircularLinkedList
	cirLinkedList.Add(values...)

	if actual := cirLinkedList.Get(4); actual != 14 {
		t.Errorf("Got %v expected %v", actual, 14)
	}
	if actual := cirLinkedList.last.next; actual != cirLinkedList.first {
		t.Errorf("Got %v expected %v", actual.data, cirLinkedList.first.data)
	}
	if actual := cirLinkedList.Size(); actual != len(values) {
		t.Errorf("Got %v expected %v", actual, len(values))
	}
}

func TestCircularLinkedListString(t *testing.T) {
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList CircularLinkedList
	cirLinkedList.Add(values...)

	var expectStr []string
	for _, v := range values {
		expectStr = append(expectStr, fmt.Sprintf("%d", v))
	}

	expected := strings.Join(expectStr, ",")

	if actual := fmt.Sprintf("%s", &cirLinkedList); actual != expected {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}
