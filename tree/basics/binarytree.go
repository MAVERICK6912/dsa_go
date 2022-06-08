package tree

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
