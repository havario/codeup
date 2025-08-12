package main

import (
	"fmt"
)

// 案例8的调用函数
// func test08(arr [3]int) {
// 	arr[0] = 88
// 	fmt.Printf("%v\n", arr)
// }

// 案例9的调用函数
func test09(arr *[3]int) {
	fmt.Printf("arr指针的地址: %p\n", &arr)
	(*arr)[0] = 88 //
}

func main() {
	/*
		1. 数组是多个相同类型数据的组合, 一个数组一旦声明/定义, 其长度是固定的不能动态变化
		2. 切片(略)
		3. 数组中的元素可以是任何数据类型包括值类型和引用类型, 但是不能混用, 长度是固定的, 不能动态变化否则越界

		4. 数组创建后, 如果没有赋值有默认值(零值)
			1> 数值 (整数系列, 浮点数系列) => 0
			2> 字符串 ==> ""
			3> bool类型 ==> false
	*/
	var arr01 [3]float32
	var arr02 [3]string
	var arr03 [3]bool
	fmt.Printf("arr01=%v arr02=%v arr03=%v\n", arr01, arr02, arr03)

	/*
		5. 使用数组的步骤
			1> 声明数组并开辟空间
			2> 给数组各个元素赋值(默认零值)
			3> 使用数组
		6. 数组的下标从0开始
	*/

	// 7. 数组下标必须在指定范围内使用, 否则报 panic 数组越界 比如var arr [5]int 则有效下标为0 - 4
	/*
		var arr04 [3]string // 0 - 2
		var index int = 3
		arr04[index] = "tom" // panic: runtime error: index out of range [3] with length 3
	*/

	// 8. Go的数组属值类型, 在默认情况下是值传递, 因此会进行值拷贝, 数组间不会相互影响
	// arr := [...]int{11, 22, 33}
	// test08(arr)
	// fmt.Printf("%v\n", arr)

	// 9. 如果在其他函数中, 去修改原来的数组, 可以使用指针传递(指针方式)
	arr := [...]int{11, 22, 33}
	fmt.Printf("数组首地址: %p\n", &arr)
	test09(&arr)
	fmt.Printf("main arr=%v\n", arr)

	// 10. 长度是数组类型的一部分, 在传递函数参数时 需要考虑数组的长度
}
