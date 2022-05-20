package linkedlist

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	first *Node
	last  *Node
	size  int
}
