// Copyright (c) 2025 honeok<i@honeok.com>

package main

import "fmt"

func main() {

	num1 := 100
	fmt.Printf("num1类型为: %T, num1的值: %v, num1的地址: %v\n", num1, num1, &num1)

	num2 := new(int) // *int
	*num2 = 611      // 改变其值
	fmt.Printf("num2类型为: %T, num2的值: %v, num2的地址: %v, num2指针指向的值是: %v\n", num2, num2, &num2, *num2)
}
