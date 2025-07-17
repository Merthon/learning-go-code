/*
Go 拥有多种值类型，包括字符串、整型、浮点型、布尔型等。
字符串可以通过 + 连接。
*/
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.1/3.0=", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)

	fmt.Println(!true)
}
