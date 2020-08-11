package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	length := len(preorder)
	return buildSubTree(preorder, inorder, 0, length-1, 0, length-1)
}

func buildSubTree(preorder []int, inorder []int, pl, pr, il, ir int) *TreeNode {

	if pr <= pl || ir <= il {
		return nil
	}
	node := &TreeNode{Val: preorder[pl]}
	index := searchIndex(inorder, preorder[pl])
	node.Left = buildSubTree(preorder, inorder, pl+1, pl+index, il, index)
	node.Right = buildSubTree(preorder, inorder, pl+index+1, pr, index+1, ir)
	return node
}

func searchIndex(arr []int, num int) int{

	// 默认数组不重复的情况下找某个数的下标
	for i:=0; i<len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}
