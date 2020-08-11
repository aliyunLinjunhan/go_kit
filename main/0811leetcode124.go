package main

//type TreeNode struct {
//    Val int
//    Left *TreeNode
//    Right *TreeNode
//}
//
//func maxPathSum(root *TreeNode) int {
//
//	var max int = ^(int(^uint(0)>>1))
//	maxSubTree(root, &max)
//	return max
//}
//
//func maxSubTree(node *TreeNode, max *int) int {
//
//	if node == nil {
//		return 0
//	}
//	left := Max124(maxSubTree(node.Left, max), 0)
//	right := Max124(maxSubTree(node.Right, max), 0)
//	if left + right + node.Val > *max {
//		*max = left + right + node.Val
//	}
//	return Max124(left, right) + node.Val
//}
//
//func Max124(x, y int) int {
//	if x > y {
//		return x
//	} else {
//		return y
//	}
//}
