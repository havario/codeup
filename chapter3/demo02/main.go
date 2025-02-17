//  变量的学习
package main

import "fmt"

func main()  {
	// golang变量的使用方式1
	// 第一种: 指定变量类型，声明后不赋值则使用默认值
	// int的默认值就是0
	var i int
	fmt.Println("i=",i)

	// 第二种: 根据值自行判断变量类型(类型推导)
	var num = 6.11
	fmt.Println("num=",num)

	// 第三种: 省略var, 注意 := 左侧的变量不应该是已经声明的,否则会导致编译错误
	// 下面的方式等价于 var name string name = "tom"
	// := 的: 不能省略,否则错误
	name := "tom"
	fmt.Println("name=",name)
}