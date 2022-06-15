package tree

import (
	"github.com/maverick6912/dsa_go/errors"
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
	return &TreeNode[T]{left: nil, right: nil}
}

/*
iArr:= {5,9,3,4,7,1,6}

-

*/

func (b *BinaryTree[T]) AddNodes(vals ...T) error {
	if b == nil {
		return errors.UninitializedError
	}
	for _, v := range vals {
		b.root = b.addNodes(b.root, v)
	}
	return nil
}

func (b *BinaryTree[T]) addNodes(root *TreeNode[T], val T) *TreeNode[T] {
	newNode := &TreeNode[T]{data: val}
	if root == nil {
		return newNode
	}
	if b.cmp(root, newNode) > 0 {
		root.left = b.addNodes(root.left, val)
	} else if b.cmp(root, newNode) < 0 {
		root.right = b.addNodes(root.right, val)
	} else {
		root.left = b.addNodes(root.left, val)
	}

	return root
}

func (b *BinaryTree[T]) Search(elem T) (bool, *TreeNode[T], error) {
	if b == nil {
		return false, nil, errors.UninitializedError
	}
	exist, node := b.searchNode(b.root, &TreeNode[T]{data: elem})
	return exist, node, nil
}

func (b *BinaryTree[T]) searchNode(root, searchNode *TreeNode[T]) (bool, *TreeNode[T]) {
	if root == nil {
		return false, nil
	}
	if b.cmp(root, searchNode) == 0 {
		return true, root
	} else if b.cmp(root, searchNode) > 0 {
		return b.searchNode(root.left, searchNode)
	} else {
		return b.searchNode(root.right, searchNode)
	}
}

func (b *BinaryTree[T]) Remove(elem T) error {
	if b == nil {
		return errors.UninitializedError
	}
	b.removeNode(b.root, &TreeNode[T]{data: elem})
	return nil
}

func (b *BinaryTree[T]) removeNode(root, remNode *TreeNode[T]) *TreeNode[T] {
	if root == nil {
		return root
	}

	if b.cmp(root, remNode) > 0 {
		root.left = b.removeNode(root.left, remNode)
	} else if b.cmp(root, remNode) < 0 {
		root.right = b.removeNode(root.right, remNode)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}
			root.data = temp.data
			root.left = b.removeNode(root.left, temp)
		}
	}
	return root
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
