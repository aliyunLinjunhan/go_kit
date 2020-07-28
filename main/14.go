package main

//func main() {
//	//matrix := [][]byte{
//	//	[]byte{'1','0','1','0','0'},
//	//	[]byte{'1','0','1','1','1'},
//	//	[]byte{'1','1','1','1','1'},
//	//	[]byte{'1','0','0','1','0'},
//	//}
//	var height int
//	var weight int
//	fmt.Scanf("%d %d\n", &height, &weight)
//
//	fmt.Println(maximalRectangle(matrix))
//}
//
//func maximalRectangle(matrix [][]byte) int {
//
//	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0])==0 {
//		return 0
//	}
//	var heights []int = make([]int, len(matrix[0]))
//	var max int = 0
//
//	for i:=0; i<len(matrix); i++ {
//		for j:=0; j<len(heights); j++ {
//			num, _ := strconv.Atoi(string(matrix[i][j]))
//			if num == 0 {
//				heights[j] = 0
//			} else {
//				heights[j] += num
//			}
//		}
//		max = Max(max, largestRectangleArea(heights))
//	}
//	return max
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


