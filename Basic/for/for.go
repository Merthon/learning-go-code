package main

import "fmt"

/*
for 是 Go 中唯一的循环结构。
最基础的方式，单个循环条件。
经典的初始/条件/后续 for 循环。
不带条件的 for 循环将一直重复执行， 直到在循环体内使用了 break 或者 return 跳出循环。
你也可以使用 continue 直接进入下一次循环。
*/

func main()  {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

	for {
        fmt.Println("loop")
        break
    }

	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
	
}
