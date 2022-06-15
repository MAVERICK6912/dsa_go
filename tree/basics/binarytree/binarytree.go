package tree

import (
	"github.com/maverick6912/dsa_go/errors"
	queue "github.com/maverick6912/dsa_go/queue/basics/listqueue"
)

type BinaryTree[T any] struct {
	cmp  func(*TreeNode[T], *TreeNode[T]) int
	root *TreeNode[T]
}

type TreeNode[T any] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (b *BinaryTree[T]) New(cmp func(*TreeNode[T], *TreeNode[T]) int) *BinaryTree[T] {
	return &BinaryTree[T]{cmp: cmp}
}

func (tn *TreeNode[T]) New() *TreeNode[T] {
	return &TreeNode[T]{}
}

func (t *BinaryTree[T]) AddLevelOrder(vals ...T) {
	if t == nil {
		return
	}
	t.root = t.addLevelOrder(vals, 0)
}

/*
	arr= [2,5,1,6,9,7]
				2
			   / \
			 5     1
		    / \    /
			6   9  7

	- If we look closely at the above tree constructed using level order:
		- for each root node, the left node is always 2*curr_index+1 of array.
		- for each root node, the right node is always 2*curr_index+2 of array.
*/
func (t *BinaryTree[T]) addLevelOrder(vals []T, index int) *TreeNode[T] {
	node := &TreeNode[T]{left: nil, right: nil}
	if index < len(vals) {
		node.data = vals[index]
		// adding nodes to left of root
		node.left = t.addLevelOrder(vals, 2*index+1)
		// adding nodes to right of root
		node.right = t.addLevelOrder(vals, 2*index+2)
	}
	return node
}

/*
Approach: level order insertion to a given binary tree can be done by traversing the tree and finding any left or right node which is empty.
	- We can achieve the above by using a queue.
	- Push the root if it isn't nil else insert new node as root.
	- Traverse the tree until the queue is empty.
Pseudo code:
	(t *binaryTree)InsertLevelOrder(data):
		if t.root==nil:
			t.root=&TreeNode{data: data}
			return
		q Queue
		for q.Length()!=0:
			tmp:= q.Peek()
			if tmp.left==nil:
				tmp.left=&TreeNode{data: data}
				return
			else:
				q.Enqueue(tmp.left)
			if tmp.right==nil:
				tmp.right=&TreeNode{data: data}
				return
			else:
				q.Enqueue(tmp.right)
*/
func (t *BinaryTree[T]) InsertNodeLevelOrder(val T, cmp func(*TreeNode[T], *TreeNode[T]) int) error {
	if t == nil {
		return errors.UninitializedError
	}
	if t.root == nil {
		t.root = &TreeNode[T]{data: val}
		return nil
	}
	var q *queue.Queue[*TreeNode[T]]
	q = q.New(cmp)

	q.Enqueue(t.root)
	for q.Size() != 0 {
		tmp, err := q.Dequeue()
		if err != nil {
			return errors.PeekError
		}

		if tmp.left == nil {
			tmp.left = &TreeNode[T]{data: val}
			return nil
		} else {
			q.Enqueue(tmp.left)
		}
		if tmp.right == nil {
			tmp.right = &TreeNode[T]{data: val}
			return nil
		} else {
			q.Enqueue(tmp.right)
		}
	}
	return nil
}

func (t *BinaryTree[T]) GetPreOrder() ([]T, error) {
	var ret []T
	if t == nil {
		return ret, errors.UninitializedError
	}
	return *getPreOrder(t.root, &ret), nil
}

func getPreOrder[T any](root *TreeNode[T], ret *[]T) *[]T {
	if root != nil {
		*ret = append(*ret, root.data)
		getPreOrder(root.left, ret)
		getPreOrder(root.right, ret)
	}
	return ret
}

// GetInorder returns inorder traversal(Left, Root, Right) of tree on which it is called.
// returns UninitializedError if tree is nil
func (t *BinaryTree[T]) GetInorder() ([]T, error) {
	var ret []T
	if t == nil {
		return ret, errors.UninitializedError
	}
	return *getInOrder(t.root, &ret), nil
}

func getInOrder[T any](root *TreeNode[T], ret *[]T) *[]T {
	if root != nil {
		getInOrder(root.left, ret)
		*ret = append(*ret, root.data)
		getInOrder(root.right, ret)
	}
	return ret
}

func (t *BinaryTree[T]) GetPostOrder() ([]T, error) {
	var ret []T
	if t == nil {
		return ret, errors.UninitializedError
	}
	return *getPostOrder(t.root, &ret), nil
}

func getPostOrder[T any](root *TreeNode[T], ret *[]T) *[]T {
	if root != nil {
		getPostOrder(root.left, ret)
		getPostOrder(root.right, ret)
		*ret = append(*ret, root.data)
	}
	return ret
}

func (t *BinaryTree[T]) GetLevelorder(cmp func(*TreeNode[T], *TreeNode[T]) int) ([]T, error) {
	var ret []T
	if t == nil {
		return ret, errors.UninitializedError
	}
	var q *queue.Queue[*TreeNode[T]]
	q = q.New(cmp)

	q.Enqueue(t.root)
	for q.Size() != 0 {
		node, err := q.Dequeue()
		ret = append(ret, node.data)
		if err != nil {
			return ret, errors.DequeueError
		}
		if node.left != nil {
			q.Enqueue(node.left)
		}
		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
	return ret, nil
}

func TreeNodeIntComparator(a, b *TreeNode[int]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}

func TreeNodeStringComparator(a, b *TreeNode[string]) int {
	if a.data > b.data {
		return 1
	} else if a.data < b.data {
		return -1
	}
	return 0
}
