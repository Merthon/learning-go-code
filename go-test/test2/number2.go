package main

import "fmt"

func main() {
	var score int
	fmt.Print("请输入成绩：")
	fmt.Scan(&score)

	if score < 0 || score > 100 {
		fmt.Println("成绩无效，请输入 0~100 的整数")
		return
	}

	switch {
	case score >= 90:
		fmt.Println("成绩优秀")
	case score >= 60:
		fmt.Println("成绩及格")
	default:
		fmt.Println("成绩不及格")
	}
}
