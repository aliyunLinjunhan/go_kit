package main

import (
	"container/list"
	"fmt"
)

//func main() {
//
//	fmt.Println(largestRectangleArea([]int{0, 2,1,5,6,2,3, 0}))
//}
//
//func largestRectangleArea(heights []int) int {
//
//	if heights == nil || len(heights) == 0 {
//		return 0
//	}
//
//	heights = append([]int{0}, heights...)
//	heights = append(heights, 0)
//
//	stack := list.New()
//	var max int = 0
//
//	for i:=0; i<len(heights); i++ {
//		for stack.Len() != 0 && heights[i] < heights[stack.Back().Value.(int)] {
//			h := heights[stack.Back().Value.(int)]
//			stack.Remove(stack.Back())
//			max = Max(max, h * (i-stack.Back().Value.(int)-1))
//		}
//		stack.PushBack(i)
//	}
//	return max
//}
//
//func Max(x, y int) int {
//	if x > y {
//		return x
//	} else {
//		return y
//	}
//}