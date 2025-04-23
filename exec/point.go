package main

import "fmt"

func main() {
	// 打印内存地址
	var num int = 9
	fmt.Printf("num address is: %v\n", &num)

	// 定义一个指针修改其值
	var ptr *int
	ptr = &num
	*ptr = 10 // 修改为10会导致num 值变化
	fmt.Printf("num value is: %v\n", num)
}

// func main() {
// 	var a int = 300
// 	var b int = 400
// 	var ptr *int = &a
// 	*ptr = 100
// 	ptr = &b
// 	*ptr = 200
// 	fmt.Printf("a=%d,b=%d,*ptr=%d\n", a, b, *ptr)
// }
