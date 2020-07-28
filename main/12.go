package main

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	index := make(map[byte]int)
	for i:=0; i<len(s1); i++ {
		if _, ok := index[s1[i]]; ok {
			index[s1[i]]++
		} else {
			index[s1[i]] = 1
		}
	}
	fmt.Println(index)
	for i:=0; i<len(s2); i++ {
		if _, ok := index[s2[i]]; ok {
			index[s2[i]]--
			if index[s2[i]] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
}
