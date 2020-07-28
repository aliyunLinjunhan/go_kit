package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var str string
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = str[:len(str)-1]

	fmt.Println(strings.ReplaceAll(str, " ", "%20"))
}
