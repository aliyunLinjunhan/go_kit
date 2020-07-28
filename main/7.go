package main

import "fmt"

func main() {

	var str string
	var ret string

	var posi int
	fmt.Scanf("%d\n", &posi)
	fmt.Scanf("%s\n", &str)

	ret += str[posi:]
	ret += str[:posi]
	fmt.Println(ret)
	return

}
