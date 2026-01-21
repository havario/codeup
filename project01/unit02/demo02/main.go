package main

import "fmt"

// 全局变量
var (
	n7 = 100
	n8 = 9.7
)

func main() {
	// 变量使用方式
	// 第一种
	var num int = 18
	fmt.Println(num)

	// 第二种: 指定变量的类型 但不赋值 使用默认值
	var num2 int
	fmt.Println(num2)

	// 第三种: 省略变量的类型 让编译器根据值自行判断变量的类型
	num3 := 10
	fmt.Println(num3)

	fmt.Println("-------------")

	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	n4, name, n5 := 10, "jack", 7.8
	fmt.Println(n4, name, n5)

	n6, height := 6.9, 100.6
	fmt.Println(n6, height)

	fmt.Println(n7, n8)
}
