package main

import "fmt"

func main() {
	// var声明1个或者多个变量
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	//:= 语法是声明并初始化变量的简写
	f := "short" //var f string = "short"

	fmt.Println(f)
}
