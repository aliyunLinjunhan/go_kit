package main

import "fmt"

func main() {

	var str string
	fmt.Scanf("%s\n", &str)

	var leftNum int
	for i:=0; i<len(str); i++ {
		if str[i] != '(' && str[i] != ')' {
			fmt.Println("NO")
			return
		}
		if str[i] == '(' {
			leftNum++
		}
		if str[i] == ')' {
			if leftNum > 0 {
				leftNum--
			} else {
				fmt.Println("NO")
				return
			}
		}
	}
	if leftNum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
