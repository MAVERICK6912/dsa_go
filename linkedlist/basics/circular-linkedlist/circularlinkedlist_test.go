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

func TestCircularLinkedListRemove(t *testing.T) {
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList CircularLinkedList
	cirLinkedList.Add(values...)

	// removing randomly
	cirLinkedList.Remove(32)
	assert.NotEqual(t, values[2], cirLinkedList.Get(2))
	assert.Equal(t, values[3], cirLinkedList.Get(2))
	assert.Equal(t, len(values)-1, cirLinkedList.Size())

	// removing head node
	cirLinkedList.Remove(25)
	assert.NotEqual(t, values[0], cirLinkedList.Get(0))
	assert.Equal(t, values[1], cirLinkedList.Get(0))
	assert.Equal(t, len(values)-2, cirLinkedList.Size())

	// removing last node
	cirLinkedList.Remove(536)
	assert.NotEqual(t, values[len(values)-4], cirLinkedList.last.data)
	assert.Equal(t, values[len(values)-2], cirLinkedList.last.data)
	assert.Equal(t, len(values)-3, cirLinkedList.Size())

	// removing a node not in linkedList
	cirLinkedList.Remove(488)
	assert.Equal(t, len(values)-4, cirLinkedList.Size())
}

func TestCircularLinkedListDelete(t *testing.T) {
	values := []int{25, 46, 32, 15, 14, 69, 47, 536}
	var cirLinkedList CircularLinkedList
	cirLinkedList.Add(values...)

	// deleting randomly
	cirLinkedList.Delete(2)
	assert.NotEqual(t, values[2], cirLinkedList.Get(2))
	assert.Equal(t, values[3], cirLinkedList.Get(2))
	assert.Equal(t, len(values)-1, cirLinkedList.Size())

	// deleting head node
	cirLinkedList.Delete(0)
	assert.NotEqual(t, values[0], cirLinkedList.Get(0))
	assert.Equal(t, values[1], cirLinkedList.Get(0))
	assert.Equal(t, len(values)-2, cirLinkedList.Size())

	// deleting last node
	cirLinkedList.Delete(len(values) - 3)
	assert.NotEqual(t, values[len(values)-4], cirLinkedList.last.data)
	assert.Equal(t, values[len(values)-2], cirLinkedList.last.data)
	assert.Equal(t, len(values)-3, cirLinkedList.Size())

	// deleting a node not in linkedList
	cirLinkedList.Delete(488)
	assert.Equal(t, len(values)-3, cirLinkedList.Size())
}
