package main

import "fmt"

//golang 指针
func main() {
	// 基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么?
	fmt.Println("i的地址=", &i)

	// 1. ptr是一个指针变量
	// 2. ptr的类型是 *int
	// 3. ptr本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)
}
