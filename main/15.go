package main

func main() {

}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	var max int = 0
	maxSubTree(root, &max)
	return max
}

func maxSubTree(node *TreeNode, max *int) int {

	if node == nil {
		return 0
	}
	left := Max(maxSubTree(node.Left, max), 0)
	right := Max(maxSubTree(node.Right, max), 0)
	*max = Max(*max, Max(left, right)+node.Val)
	return Max(left, right) + node.Val
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}