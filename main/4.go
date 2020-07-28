package main

import "fmt"

// 找到指定类型的新类型字符串

func main() {

	var length int
	var posi int
	fmt.Scanf("%d %d\n", &length, &posi)

	var str string
	fmt.Scanf("%s\n", &str)

	var index int = 0

	for {
		if str[index] >= 'a' && str[index] <= 'z' {
			if posi == index {
				fmt.Println(string(str[index]))
				return
			} else {
				index++
				continue
			}
		}
		if str[index] >= 'A' && str[index] <= 'Z' {
			if posi == index || posi == index+1{
				fmt.Println(string(str[index: index+2]))
				return
			} else {
				index = index + 2
				continue
			}
		}
	}
}
