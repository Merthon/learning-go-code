package main

import "fmt"

func main()  {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空的切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片是之前的两倍 */
	numbers1 := make([]int,len(numbers), (cap(numbers))*2)

	/* 将number的内容拷贝到number1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}