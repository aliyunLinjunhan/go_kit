package main

import "fmt"

// 给定一个字符串数组strs，再给定两个字符串str1和str2，返回在strs中str1和str2的最小距离，如果str1或str2为null，或不在strs中，返回-1。

// 输入包含有多行，第一输入一个整数n
//（
//1
//≤
//n
//≤
//1
//0
//5
//）
//（1≤n≤10
//5
// ），代表数组strs的长度，第二行有两个字符串分别代表str1和str2，接下来n行，每行一个字符串，代表数组strs (保证题目中出现的所有字符串长度均小于等于10)。

//func main() {
//
//	var length int
//	fmt.Scanf("%d\n", &length)
//
//	var str1, str2 string
//	fmt.Scanf("%s %s\n", &str1, &str2)
//	if str1 == "" || str2 == "" {
//		fmt.Println(-1)
//		return
//	}
//
//	var strs []string
//	for i:=0; i<length; i++ {
//		str := ""
//		fmt.Scanf("%s\n", &str)
//		strs = append(strs, str)
//	}
//
//	var posi1 []int
//	var posi2 []int
//	for i:=0; i<len(strs); i++ {
//		if strs[i] == str1 {
//			posi1 = append(posi1, i)
//		}
//		if strs[i] == str2 {
//			posi2 = append(posi2, i)
//		}
//	}
//	if len(posi1) == 0 || len(posi2) == 0 {
//		fmt.Println(-1)
//		return
//	}
//	var min = 278932332
//	for i:=0; i<len(posi1); i++ {
//		for j:=0; j<len(posi2); j++ {
//			if abs(posi1[i]-posi2[j]) < min {
//				min = abs(posi1[i]-posi2[j])
//			}
//		}
//	}
//	fmt.Println(min)
//}

func main() {

	var length int
	fmt.Scanf("%d\n", &length)

	var str1, str2 string
	fmt.Scanf("%s %s\n", &str1, &str2)
	if str1 == "" || str2 == "" {
		fmt.Println(-1)
		return
	}

	var strs []string
	for i:=0; i<length; i++ {
		str := ""
		fmt.Scanf("%s\n", &str)
		strs = append(strs, str)
	}
	var min = 278932332

	var last1 int = -1
	var last2 int = -1
	for i:=0; i<len(strs); i++ {
		if str1 == strs[i] {
			last1 = i
			if last2 == -1 {

			} else {
				if abs(i-last2) < min {
					min = abs(i-last2)
				}
			}
		}
		if str2 == strs[i] {
			last2 = i
			if last1 == -1 {

			} else {
				if abs(i-last1) < min {
					min = abs(i-last1)
				}
			}
		}
	}
	if last1 == -1 || last2 == -1 {
		fmt.Println(-1)
		return
	}
	fmt.Println(min)

}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
