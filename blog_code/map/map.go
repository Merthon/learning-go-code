package main

import (
	"fmt"
)

func main() {
	/* 第一种声明map */
	var test1 map[string]string
	//在使用之前，需要先make，make的作用就是给map分配数据空间
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "go"
	test1["three"] = "java"

	fmt.Println(test1)

	/* 第二种声明 */
	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "go"
	test2["three"] = "java"

	fmt.Println(test2)

	/* 第三种声明 */
	test3 := map[string]string{
		"one":   "php",
		"two":   "go",
		"three": "java",
	}
	fmt.Println(test3)
}
