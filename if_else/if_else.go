package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7是偶数")
	} else {
		fmt.Println("7是奇数")
	}

	if 8%4 == 0 {
		fmt.Println("8可以被四整除")
	}

	var num int = 9
	if num < 0 {
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "是一位数")
	} else {
		fmt.Println(num, "是多位数")
	}
}
