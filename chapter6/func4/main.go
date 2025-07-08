// Copyright (c) 2025 honeok <honeok@disroot.org>
//                           <i@honeok.com>

package main

import (
	"fmt"
)

// // 函数的变量是局部的，函数外不生效
// func test() {
// 	// n1是test函数的局部变量，只能在test函数中使用
// 	var n1 int = 10
// }

// 基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值
// func test02(n1 int) {
// 	n1 += 10
// 	fmt.Println("test02()n1 = ", n1)
// }

// n1是 *int类型
func test03(n1 *int) {
	*n1 += 10
	fmt.Println("test03()n1 = ", *n1)
}

// go不支持函数重载

func main() {
	// 不能使用n1，因为n1是test函数的局部变量
	// n1 := 20
	// test02(n1)
	// fmt.Println("main()n1 = ", n1)

	num := 20
	test03(&num)
	fmt.Println("main()n1 = ", num)
}
