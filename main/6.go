package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	var words []string
	words = strings.Split(str, " ")
	var ret string

	for i:=0; i<len(words); i++ {
		ret = ret + Reverse(words[i]) + " "
	}
	fmt.Println(ret[:len(ret)-1])
	return
}

func Reverse(str string) (ret string) {

	for i:=len(str)-1; i>=0; i-- {
		ret += string(str[i])
	}
	return
}
