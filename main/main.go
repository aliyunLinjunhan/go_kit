package main

import (
	"fmt"
	"time"
)

//
//import (
//	"fmt"
//	"strconv"
//)
//
////给定一个字符串str，如果str符合日常书写的整数形式，并且属于32位整数范围，返回str代表的整数值，否则返回0。
//func main() {
//
//	str := ""
//	fmt.Scanf("%s\n", &str)
//	var index bool = false
//
//	var ret int64
//	for i, v := range str {
//		if i == 0 {
//			if v == '-' {
//				index = true
//				continue
//			}
//			if v >= '1' && v <= '9' {
//				value, _ := strconv.Atoi(string(v))
//				ret = ret + int64(value)
//				continue
//			}
//			fmt.Println(0)
//			return
//		}
//		if i == 1 {
//			if index {
//				if v >= '1' && v <= '9' {
//					value, _ := strconv.Atoi(string(v))
//					ret = ret + int64(value)
//					continue
//				} else {
//					fmt.Println(0)
//					return
//				}
//			} else {
//				if v >= '0' && v <= '9' {
//					value, _ := strconv.Atoi(string(v))
//					ret = ret*10 + int64(value)
//					continue
//				} else {
//					fmt.Println(0)
//					return
//				}
//			}
//		}
//		if v >= '0' && v <= '9' {
//			value, _ := strconv.Atoi(string(v))
//			ret = ret*10 + int64(value)
//			continue
//		} else {
//			fmt.Println(0)
//			return
//		}
//
//	}
//	if index {
//		if ret < 2147483649 {
//			fmt.Println(-ret)
//		} else {
//			fmt.Println(0)
//		}
//	} else {
//		if ret < 2147483648 {
//			fmt.Println(ret)
//		} else {
//			fmt.Println(0)
//		}
//	}
//
//
//}

func main() {
	var user = ""
	defer fmt.Println("main defer there")
	go func() {
		defer func() {
			fmt.Println("defer here")
		}()
		defer func() {
			fmt.Println("defer here two")
		}()

		if user == "" {
			panic("should set user env.")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("get result \r\n")
}
