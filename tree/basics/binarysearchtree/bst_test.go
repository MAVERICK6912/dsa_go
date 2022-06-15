package tree

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/stretchr/testify/assert"
)

func TestAddNodes(t *testing.T) {
	vals := []int{5, 9, 3, 4, 7, 1, 6}

	var bst *BinaryTree[int]
	bst = bst.New(TreeNodeIntComparator)

	err := bst.AddNodes(vals...)
	if err != nil {
		t.Error("error while adding nodes to bst. error was: ", err.Error())
	}
	/*
					5
				  /   \
				 3     9
				/ \   /
			   1   4 7
			        /
				   6
		InOrder: {1,3,4,5,6,7,9}
	*/

	expected := []int{1, 3, 4, 5, 6, 7, 9}

	actual, err := bst.GetInorder()
	if err != nil {
		t.Error("error while getting inorder of bst. error was: ", err.Error())
	}
	assert.Equal(t, expected, actual)
}

func TestSearch(t *testing.T) {
	vals := []int{5, 9, 3, 4, 7, 1, 6}

	var bst *BinaryTree[int]
	bst = bst.New(TreeNodeIntComparator)

	err := bst.AddNodes(vals...)
	if err != nil {
		t.Error("error while adding nodes to bst. error was: ", err.Error())
	}

	// search for an existing node
	exist, actual, err := bst.Search(6)
	if err != nil {
		t.Error("error while searching node in bst. error was: ", err.Error())
	}

	assert.Equal(t, true, exist)
	assert.Equal(t, actual.data, 6)

	// search for node that doesn't exist in bst

	exist, _, err = bst.Search(69)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
	assert.Equal(t, false, exist)
}

func TestRemove(t *testing.T) {
	vals := []int{5, 9, 3, 4, 7, 1, 6}

	var bst *BinaryTree[int]
	bst = bst.New(TreeNodeIntComparator)

	err := bst.AddNodes(vals...)
	if err != nil {
		t.Error("error while adding nodes to bst. error was: ", err.Error())
	}
	expectedInOrder := []int{1, 3, 4, 5, 7, 9}
	// remove a node which exists
	err = bst.Remove(6)
	assert.Nil(t, err)
	actual, err := bst.GetInorder()
	if err != nil {
		t.Error("error while getting inorder of bst. error was: ", err.Error())
	}
	assert.Equal(t, expectedInOrder, actual)

	// remove root of bst
	err = bst.Remove(bst.root.data)
	assert.Nil(t, err)
	expectedInOrder = []int{1, 3, 4, 7, 9}
	actual, err = bst.GetInorder()
	if err != nil {
		t.Error("error while getting inorder of bst. error was: ", err.Error())
	}
	assert.Equal(t, expectedInOrder, actual)
}
