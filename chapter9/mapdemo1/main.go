/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

package main

import (
	"fmt"
)

/*
基本语法:
var 变量名 map[keyType]valueType

valueType通常为数字(整数 浮点数) string map struct
*/

func main() {
	// map的声明和注意事项
	var a map[string]string

	/*
		使用map前面需要先make, make的作用是给map分配数据空间
		前面参数是map数据类型, 后面是make分配的空间
	*/

	a = make(map[string]string, 10) // 10个kv键值对
	a["1"] = "Rosé"
	a["2"] = "Jennie"
	a["1"] = "Lisa"   // key不能重复, 如果重复以最后的定义为准
	a["3"] = "Jennie" // value可以重复
	fmt.Printf("%v\n", a)
}
