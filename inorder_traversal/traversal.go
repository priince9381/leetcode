package inordertraversal

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Traversal []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return Traversal
	}
	inorderTraversal(root.Left)
	Traversal = append(Traversal, root.Val)
	inorderTraversal(root.Right)
	return Traversal
}

func TestInorderTraversal() {
	root := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}
	fmt.Println(inorderTraversal(root))
}
