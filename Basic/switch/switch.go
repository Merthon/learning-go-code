package main

import (
	"fmt"
	"time"
)

/*
switch是多分支情况时条件语句
一个基本的switch,在同一个case语句中，可以使用逗号来分隔多个表达式。
不带表达式的switch是实现if/else逻辑的另外一种方式
类型开关是比较，类型而非值。
*/
func main()  {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's before noon")
	}

	whatAmI := func (i interface{})  {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}