/*
Go中变量需要显式声明，并且在函数调用等情况下,编译器会检查其类型的正确性。
var 声明 1 个或者多个变量。
你可以一次性声明多个变量。
Go 会自动推断已经有初始值的变量的类型。
声明后却没有给出对应的初始值时，变量将会初始化为 零值 。 例如，int 的零值是 0。
:= 语法是声明并初始化变量的简写
*/
package main

import "fmt"

func main()  {
	var a = "chen"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

}

