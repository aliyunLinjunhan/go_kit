package main

import "fmt"

// 给定字符串str1和str2，求str1的字串中含有str2所有字符的最小字符串长度。

//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//
//字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。


func main() {
	str1 := ""
	str2 := ""
	fmt.Scanf("%s\n", &str1)
	fmt.Scanf("%s\n", &str2)

	if len(str1) == 0 || len(str2) == 0 ||len(str1) < len(str2) {
		fmt.Println(0)
		return
	}

	left := 0
	right := 0
	match := len(str2)
	min := int(^uint(0) >> 1)

	ret := []int{}
	record := make(map[byte]int)
	for i:=0; i<len(str2); i++ {
		if _, ok := record[str2[i]]; ok {
			record[str2[i]]++
		} else {
			record[str2[i]] = 1
		}
	}

	for right < len(str1) {
		if match > 0 {
			if v, ok := record[str1[right]]; ok {
				if v > 0 {
					match--
				}
				record[str1[right]]--
			} else {
				record[str1[right]] = -1
			}
			right++
		} else {
			if v, ok := record[str1[left]]; ok {
				if v >= 0 {
					match++
				}
				record[str1[left]]++
			}
			left++

		}
		if match == 0 && right - left < min {
			min = right - left
		}
		if match == 0 {
			boo := true
			for _, v := range record {
				if v != 0 {
					boo = false
				}
			}
			if boo {
				ret = append(ret, left)
			}
		}
	}
	if min < int(^uint(0)>>1) {
		fmt.Println(min)
	} else {
		fmt.Println(0)
	}
	fmt.Println(ret)
}