/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

// map的使用

package main

import (
	"fmt"
)

func main() {
	// 第一种
	var a map[string]string
	a = make(map[string]string, 10)
	a["1"] = "Rosé"
	a["2"] = "Jennie"
	a["1"] = "Lisa"
	a["3"] = "Jennie"
	fmt.Printf("%v\n", a)

	// 第二种
	citys := make(map[string]string)
	citys["1"] = "BeiJing"
	citys["2"] = "TianJing"
	citys["3"] = "ShangHai"
	fmt.Printf("%v\n", citys)

	// 第三种
	heros := map[string]string{
		"hero1": "武松",
		"hero2": "武大郎",
		"hero3": "潘金莲",
	}
	fmt.Printf("%v\n", heros)

	/*
		课堂练习: 演示一个key-value的value是map的案例
		比如: 我们要存放3个学生信息, 每个学生有name和gender信息
		思路: map[string]map[string]string
	*/
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["gender"] = "男"
	studentMap["stu01"]["address"] = "北京长安街"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Mary"
	studentMap["stu02"]["gender"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江"

	fmt.Printf("%v\n", studentMap["stu01"])
	fmt.Printf("%v\n", studentMap["stu02"])
}
