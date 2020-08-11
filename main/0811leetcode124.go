type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	var max int = ^(int(^uint(0)>>1))
	maxSubTree(root, &max)
	return max
}

func maxSubTree(node *TreeNode, max *int) int {

	if node == nil {
		return 0
	}
	int left := maxSubTree()
}

func Max124(x, y int) {
	
}
