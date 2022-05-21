package linkedlist

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestCircularLinkedListInsert(t *testing.T) {
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList CircularLinkedList
	cirLinkedList.Add(values...)

	// inserting between linkedList
	cirLinkedList.Insert(3, 321)
	if actual := cirLinkedList.Get(3); actual != 321 {
		t.Errorf("Got %v expected %v", actual, 321)
	}
	if actual := cirLinkedList.Size(); actual != 9 {
		t.Errorf("Got %v expected %v", actual, 9)
	}

	// inserting at start
	cirLinkedList.Insert(0, 634)
	if actual := cirLinkedList.Get(0); actual != 634 {
		t.Errorf("Got %v expected %v", actual, 634)
	}
	if actual := cirLinkedList.Size(); actual != 10 {
		t.Errorf("Got %v expected %v", actual, 10)
	}
	if cirLinkedList.last.next != cirLinkedList.first && cirLinkedList.first.data != 634 {
		t.Error("last node is not pointing to correct node(first node)")
	}

	// inserting at the end
	cirLinkedList.Insert(cirLinkedList.Size(), 999)
	if actual := cirLinkedList.Get(10); actual != 999 {
		t.Errorf("Got %v expected %v", actual, 999)
	}
	if actual := cirLinkedList.Size(); actual != 11 {
		t.Errorf("Got %v expected %v", actual, 11)
	}
}

func TestCircularLinkedListSort(t *testing.T) {
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList CircularLinkedList
	cirLinkedList.Add(values...)

	sort.Ints(values)
	cirLinkedList.Sort()

	assert.Equal(t, values, cirLinkedList.Values())
}
