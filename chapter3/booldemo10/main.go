// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
	"unsafe"
)

// 演示golang中bool类型使用
func main() {
	var b = false
	fmt.Println("b =", b)

	// 注意:
	// 1. bool类型占用的存储空间是1个字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
	// 2. bool类型只能取true或false
}