package main

import "fmt"

/* 用户输入一个整数，判断奇偶性 */
func main() {
	var a int
	fmt.Print("请输入一个数")
	fmt.Scan(&a)

	if a % 2 == 0 {
		fmt.Println("这个数是偶数")
	}else {
		fmt.Println("这个数是奇数")
	}
}