package main

import "fmt"

func main() {

	var str string
	fmt.Scanf("%s\n", &str)

	dp := getDp(str)
	fmt.Println(getString(str, dp))
}

func getDp(str string) [][]int {

	var ret [][]int = make([][]int, len(str))

	for i:=0; i<len(ret); i++ {
		ret[i] = make([]int, len(str))
	}
	for i:=1; i<len(str); i++ {
		if str[i-1] == str[i] {
			ret[i-1][i] = 0
		} else {
			ret[i-1][i] = 1
		}
	}
	for i:=2; i<len(str); i++ {
		for j := i-2; j>=0; j-- {
			if str[i] == str[j] {
				ret[j][i] = ret[j+1][i-1]
			} else {
				if ret[j+1][i] < ret[j][i-1] {
					ret[j][i] = ret[j+1][i] + 1
				} else {
					ret[j][i] = ret[j][i-1] + 1
				}
			}
		}
	}
	return ret
}

func getString(source string, dp [][]int) string {

	var res []byte = make([]byte, len(source) + dp[0][len(source)-1])
	sl, sr := 0, len(source)-1
	rl, rr := 0, len(res) - 1
	for sl <= sr {
		if source[sl] == source[sr] {
			res[rl] = source[sl]
			res[rr] = source[sr]
			rl++
			sl++
			rr--
			sr--
		} else {
			if dp[sl][sr-1] < dp[sl+1][sr] {
				res[rl] = source[sr]
				res[rr] = source[sr]
				rl++
				rr--
				sr--
			} else {
				res[rl] = source[sl]
				res[rr] = source[sl]
				rl++
				rr--
				sl++
			}
		}
	}
	return string(res)
}
