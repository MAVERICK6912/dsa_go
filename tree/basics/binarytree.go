package tree

import queue "github.com/maverick6912/dsa_go/queue/basics/listqueue"

type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func New() *BinaryTree {
	return &BinaryTree{root: newTreeNode()}
}

func newTreeNode() *TreeNode {
	return &TreeNode{}
}

// func (t *BinaryTree) getCurrentLevel() string {
// 	return ""
// }

func (t *BinaryTree) AddLevelOrder(vals ...int) {
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
func (t *BinaryTree) addLevelOrder(vals []int, index int) *TreeNode {
	node := &TreeNode{}
	if index < len(vals) {
		node = &TreeNode{data: vals[index]}
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
func (t *BinaryTree) InsertNodeLevelOrder(val int) {
	if t == nil {
		return
	}
	if t.root == nil {
		t.root = &TreeNode{data: val}
		return
	}
	q := queue.New()
	q.Enqueue(t.root.data)
	for q.Size() != 0 {
		// tmp := t.root
		// TODO: complete this after generic support
	}
}

func (t *BinaryTree) GetPreOrder() string {
	return ""
}

func (t *BinaryTree) GetInorder() string {
	return ""
}

func (t *BinaryTree) GetPostOrder() string {
	return ""
}

func (t *BinaryTree) GetLevelorder() string {
	return ""
}
