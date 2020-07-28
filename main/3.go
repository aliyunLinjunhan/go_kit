package main

import (
	"fmt"
	"bufio"
	"os"
)

// 给定一个数组arr，返回arr的最长无的重复子串的长度(无重复指的是所有字母都不相同)。
// 找出字符串的最长无重复字符串

// 解法1 牛客上过不了
//func main() {
//
//	var length int
//	fmt.Scanf("%d\n", &length)
//	var arr []byte
//	inputReader := bufio.NewReader(os.Stdin)
//	input, _ := inputReader.ReadString('\n')
//
//	for i:=0; i<len(input); i++ {
//		if input[i] != ' ' && input[i] != '\n'{
//			arr = append(arr, input[i])
//		}
//	}
//
//	record := make(map[byte]int)
//	max := ^(int(^uint(0) >> 1))
//	left := 0
//	right := 0
//	count := 0
//
//	for right < len(arr) && left < len(arr) {
//		//fmt.Println(count, record, left, right)
//		if count == 0 {
//			if v, ok := record[arr[right]]; !ok {
//				record[arr[right]] = 1
//				right++
//			} else {
//				record[arr[right]]++
//				if v > 0 {
//					count++
//					right++
//				} else {
//					right++
//				}
//			}
//		} else {
//			if record[arr[left]] > 1 {
//				record[arr[left]]--
//				if record[arr[left]] <= 1 {
//					count--
//				}
//			} else {
//				record[arr[left]] = 0
//			}
//			left++
//		}
//		if count == 0 && right - left > max {
//			max = right - left
//		}
//	}
//	if max == ^(int(^uint(0) >> 1)) {
//		max = 1
//	}
//	fmt.Println(max)
//}

// 解法2
func main() {

	var length int
	fmt.Scanf("%d\n", &length)
	var arr []byte
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	for i:=0; i<len(input); i++ {
		if input[i] != ' ' && input[i] != '\n'{
			arr = append(arr, input[i])
		}
	}

	record := make(map[byte]int)
	max := ^(int(^uint(0) >> 1))
	left := 0
	right := 0

	for right < len(arr) {
		if v, ok := record[arr[right]]; !ok {

		} else {
			if v >= left {
				left = v + 1
			}
		}
		record[arr[right]] = right
		right++
		if right - left > max {
			max = right - left
		}
	}
	fmt.Println(max)
}