package model

import (
	"example.com/m/src/helper"
	"gopkg.in/guregu/null.v4"
	"sync"
)




// TreeNode class
type TreeNode struct {
	key   string
	value int64
	left  null.String
	right null.String

	leftNode  *TreeNode
	rightNode *TreeNode
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) InitTree(treeData RequestBody) {

	for _, row := range treeData.Tree.Nodes {

		tree.InsertElement(row.ID, row.Value, row.Left, row.Right)
	}

}

func (tree *BinarySearchTree) MaxPathSum() int64 {
	if tree.rootNode == nil {
		return 0
	}
	maxValue := tree.rootNode.value

	// dfs return the max sum path that starts from root
	var dfs func(root *TreeNode) int64
	dfs = func(root *TreeNode) int64 {
		if root == nil {
			return 0
		}
		left := helper.MaxVal(0, dfs(root.leftNode))
		right := helper.MaxVal(0, dfs(root.rightNode))
		sum := root.value + left + right
		if sum > maxValue {
			maxValue = sum
		}

		return helper.MaxVal(left, right) + root.value
	}

	dfs(tree.rootNode)
	return maxValue
}

// InsertElement method
func (tree *BinarySearchTree) InsertElement(key string, value int64, left, right null.String) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, left, right, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}

}

// insertTreeNode method
func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {

	if newTreeNode.value < rootNode.value {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	}else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}

}

