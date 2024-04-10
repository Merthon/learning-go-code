package main

import "fmt"

func main() {
	/*定义局部变量*/
	var a int = 100
	b := 200
	number := 0

	number = max(a, b)
	fmt.Println("最大值是", number)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}
