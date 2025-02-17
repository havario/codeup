// Copyright (C) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
	"unsafe"
)

// 演示go中整数类型的使用
func main() {
	var i int = 1
	fmt.Println("i=", i)

	// 测试int8的范围 -128 ~ 127
	// int16 int32 int64 类推
	var j int8 = 127
	fmt.Println("j=", j)

	// 测试uint8的范围 0 ~ 255
	var k uint8 = 255
	fmt.Println("k=", k)

	// int uint rune byte 使用
	var a int = 8900
	var b uint = 1
	var c byte = 255
	fmt.Println("a=", a, "b=", b, "c=", c)

	// 整型的使用细节
	var n1 = 100 // n1是什么类型
	// 如何查看某个变量的数据类型

	//fmt.Printf("") 格式化输出
	fmt.Printf("n1的类型是: %T\n", n1)

	// 如何在程序查看某个变量的占用字节大小和数据类型 (使用较多)
	var n2 int64 = 10
	// unsafe.Sizeof(n1) 是unsafe包的一个函数, 可以返回n1变量的占用字节数
	fmt.Printf("n2的数据类型: %T n2占用的字节数是: %d", n2, unsafe.Sizeof(n2)) 

	/* 	golang程序中整型变量在使用中, 遵守保小不保大的原则
		即在保证程序正常运行情况下, 尽量使用占用空间小的数据类型 */
	var age byte = 90
	fmt.Printf("age=", age)
}