// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
)

func main() {
	// 在golang中++ 和 --只能独立使用
	// var i int = 8
	// var a int
	// a = i++ // i++ 和 i-- 只能独立使用
	// a = i--

	// if i++ > 0 {
	// 	println("ok")
	// }

	var i int = 1
	i++
	fmt.Println("i=", i)
}
