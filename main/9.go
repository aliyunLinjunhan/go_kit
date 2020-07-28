package main

import "fmt"

// 爆破
//func isEffective(str string) bool {
//
//	var leftNum int
//	for i:=0; i<len(str); i++ {
//		if str[i] != '(' && str[i] != ')' {
//			return false
//		}
//		if str[i] == '(' {
//			leftNum++
//		}
//		if str[i] == ')' {
//			if leftNum > 0 {
//				leftNum--
//			} else {
//				return false
//			}
//		}
//	}
//	if leftNum == 0 {
//		return true
//	} else {
//		return false
//	}
//}

func main() {

	var str string
	fmt.Scanf("%s\n", &str)
	var pre int
	var max int

	var dp []int = make([]int, len(str))

	for i:=1; i<len(dp); i++ {
		pre = i - dp[i-1] - 1
		if str[i] == ')' {
			if pre >=0 && str[pre] == '(' {
				dp[i] = dp[i-1] + 2
				if pre > 0 {
					dp[i] += dp[pre-1]
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	fmt.Println(max)
}
