package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	first *Node
	last  *Node
	size  int
}

func (c *CircularLinkedList) isWithinRange(index int) bool {
	return index >= 0 && index <= c.size
}

func (c *CircularLinkedList) String() string {
	var strArr []string

	for node := c.first; node != c.last; node = node.next {
		strArr = append(strArr, fmt.Sprintf("%d", node.data))
	}

	return strings.Join(strArr, ",")
}
