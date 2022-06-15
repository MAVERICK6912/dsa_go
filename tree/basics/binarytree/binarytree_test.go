package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLevelOrder(t *testing.T) {
	t.Parallel()
	var bt *BinaryTree[int]
	bt = bt.New(TreeNodeIntComparator)

	vals := []int{21, 3, 54, 69, 2, 1, 7}

	bt.AddLevelOrder(vals...)
	/*
		Level order should look like this:
						21
						/ \
					   3    54
					  / \  / \
					 69  2 1  7
	*/

	assert.Equal(t, bt.root.data, 21)
	assert.Equal(t, bt.root.left.data, 3)
	assert.Equal(t, bt.root.right.data, 54)
	assert.Equal(t, bt.root.left.left.data, 69)
	assert.Equal(t, bt.root.left.right.data, 2)
	assert.Equal(t, bt.root.right.left.data, 1)
	assert.Equal(t, bt.root.right.right.data, 7)
}

func TestInsertNodeLevelOrder(t *testing.T) {
	t.Parallel()
	var bt *BinaryTree[int]
	bt = bt.New(TreeNodeIntComparator)

	err := bt.InsertNodeLevelOrder(87, TreeNodeIntComparator)
	if err != nil {
		t.Error("error while inserting to tree")
	}
	assert.Equal(t, 87, bt.root.data)

	err = bt.InsertNodeLevelOrder(69, TreeNodeIntComparator)
	if err != nil {
		t.Error("error while inserting to tree")
	}
	assert.Equal(t, 69, bt.root.left.data)

	err = bt.InsertNodeLevelOrder(1, TreeNodeIntComparator)
	if err != nil {
		t.Error("error while inserting to tree")
	}
	assert.Equal(t, 1, bt.root.right.data)

	err = bt.InsertNodeLevelOrder(25, TreeNodeIntComparator)
	if err != nil {
		t.Error("error while inserting to tree")
	}
	assert.Equal(t, 25, bt.root.left.left.data)

	err = bt.InsertNodeLevelOrder(34, TreeNodeIntComparator)
	if err != nil {
		t.Error("error while inserting to tree")
	}
	assert.Equal(t, 34, bt.root.left.right.data)
}

func TestGetLevelorder(t *testing.T) {
	t.Parallel()
	var bt *BinaryTree[int]
	bt = bt.New(TreeNodeIntComparator)

	vals := []int{21, 3, 54, 69, 2, 1, 7}
	for _, v := range vals {
		bt.InsertNodeLevelOrder(v, TreeNodeIntComparator)
	}

	/*
			21
			/ \
		   3    54
		  / \  / \
		69   2 1  7
	*/
	actual, err := bt.GetLevelorder(TreeNodeIntComparator)
	if err != nil {
		t.Error("error while getting levelOrder of tree")
	}

	assert.Equal(t, vals, actual)
}

func TestGetInorder(t *testing.T) {
	t.Parallel()
	var bt *BinaryTree[int]
	bt = bt.New(TreeNodeIntComparator)

	vals := []int{1, 2, 3, 4, 5}
	/*
		Binary tree for above will look like this:
				1
			   / \
			  2   3
			 / \
			4   5
		Inorder traversal is left,root,right. Hence in-order traversal will look like this:
			4: is left of left
			2: is curr_root of left of left
			5: is right of curr_root
			1: is the root of curr_root
			3: is the right of root
		So, final result of traversal will be [4,2,5,1,3]
	*/

	exp := []int{4, 2, 5, 1, 3}
	for _, v := range vals {
		bt.InsertNodeLevelOrder(v, TreeNodeIntComparator)
	}

	actual, err := bt.GetInorder()
	if err != nil {
		t.Error("error while inorder traversal")
	}
	assert.Equal(t, exp, actual)
}

func TestGetPreOrder(t *testing.T) {
	t.Parallel()
	var bt *BinaryTree[int]
	bt = bt.New(TreeNodeIntComparator)

	vals := []int{1, 2, 3, 4, 5}
	/*
		Binary tree for above will look like this:
				1
			   / \
			  2   3
			 / \
			4   5
		Inorder traversal is root,left,right. Hence in-order traversal will look like this:
			1: is root of tree
			2: is left of root
			4: is left of left root
			5: is right of left of root
			3: is the right of root
		So, final result of traversal will be [4,2,5,1,3]
	*/

	exp := []int{1, 2, 4, 5, 3}
	for _, v := range vals {
		bt.InsertNodeLevelOrder(v, TreeNodeIntComparator)
	}

	actual, err := bt.GetPreOrder()
	if err != nil {
		t.Error("error while inorder traversal")
	}
	assert.Equal(t, exp, actual)
}

func TestGetPostOrder(t *testing.T) {
	t.Parallel()
	var bt *BinaryTree[int]
	bt = bt.New(TreeNodeIntComparator)

	vals := []int{1, 2, 3, 4, 5}
	/*
		Binary tree for above will look like this:
				1
			   / \
			  2   3
			 / \
			4   5
		Inorder traversal is left,right,root. Hence post-order traversal will look like this:
			4: is left of left of root
			5: is right of left of left
			2: is curr_root of left
			3: is right of the root
			1: is the root
		So, final result of traversal will be [4,5,2,3,1]
	*/

	exp := []int{4, 5, 2, 3, 1}
	for _, v := range vals {
		bt.InsertNodeLevelOrder(v, TreeNodeIntComparator)
	}

	actual, err := bt.GetPostOrder()
	if err != nil {
		t.Error("error while inorder traversal")
	}
	assert.Equal(t, exp, actual)
}
