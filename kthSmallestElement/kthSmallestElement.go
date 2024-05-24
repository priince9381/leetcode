package kthsmallestelement

import (
	"fmt"
	"sort"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeIteration(node *TreeNode, treevalue []int) []int {
	if node == nil {
		return treevalue
	}
	// Visit the current node
	treevalue = append(treevalue, node.Val)
	// Traverse the left subtree
	treevalue = treeIteration(node.Left, treevalue)
	// Traverse the right subtree
	treevalue = treeIteration(node.Right, treevalue)
	return treevalue
}

func kthSmallest(root *TreeNode, k int) int {
	t := treeIteration(root, []int{})
	sort.Ints(t)
	return t[k-1]
}

func TestKthSmallestElement() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthSmallest(root, 1))
}
