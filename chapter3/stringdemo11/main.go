// Copyright (c) 2025 honeok <honeok@duck.com>

package main

import (
	"fmt"
)

func main() {
	// string的基本使用
	var address string = "\"I want to go\"\n \"瑞典斯德哥尔摩\""
	fmt.Println(address)
	fmt.Printf("cnm\nstring类型\n")

	// 字符串一旦赋值, 字符串就无法更改: 在go中字符串是不可变的
	// var str = "hello"
	// str[0] = 'a' // 这里就不能修改string的内容, 即go中的字符串是不可变的

	// 字符串的两种表示形式 双引号, 会识别转义字符 反引号
	// 以字符串的原生形式输出, 包括换行和特殊字符. 可以实现防止攻击
	// 输出原代码等效果
	str2 := "abc\nabc\n"
	fmt.Println(str2)

	// 使用反引号 ``
	str3 := `
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
	`
	fmt.Printf("反引号效果: %s\n", str3)

	// 字符串的拼接
	var str = "hello" + " world"
	str += " shabi!"
	fmt.Println(str)

	// 当一个拼接的操作很长时, 怎么办?
	// 可以分行, 需要将+ 号保留上一行即可
	var str4 = "hello" + " world!" +
	" hello" + " world!"
	
	fmt.Println(str4)

	// 查看数据类型的零值
	var a int // 0
	var b float32 // 0
	var c float64 // 0
	var d bool // false
	var name string // ""

	// %d: 十进制整数, 用于int类型
	// %f: 浮点数 默认6位小数, 用于float32和float64
	// %v: 通用占位符, 根据类型输出默认格式, 适用于任何类型

	fmt.Printf("a = %d, b = %f, c = %f, d = %v, name = %v\n", a, b, c, d, name)
}