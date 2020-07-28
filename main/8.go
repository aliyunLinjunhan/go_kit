package main

import "fmt"

func main() {

	var str string
	fmt.Scanf("%s\n", &str)

	var ret string

	for i:=0; i<len(str); i++ {
		if str[i] == '*' {
			ret = "*" + ret
			continue
		}
		ret += string(str[i])
	}
	fmt.Println(ret)
}
