/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

/*
map使用细节
*/

package main

import (
	"fmt"
)

// 定义学生结构体
type Stu struct {
	Name      string
	Age       int
	Telephone string
}

func modify(map1 map[int]int) {
	map1[10] = 900
}

func main() {
	/*
		1> map是引用类型, 遵守引用类型传递的机制, 在一个函数接收map, 修改后，会直接修改原来的map
		2> map的容量达到后, 再想map增加元素会自动扩容, 并不会发生panic, 也就是说map能动态的增长键值对 (key-value)
	*/
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1) // 如果是900代表引用传递, 如果仍然是1则是值传递
	fmt.Printf("%v\n", map1)

	/*
		3> map的value也经常使用 `struct类型`, 更适合管理复杂的数据 (比前面value是一个map更好)

		map的key为学生的学号, 是唯一的
		map的value为结构体, 包含学生的名字 年龄 地址
	*/

	students := make(map[string]Stu)

	// 创建学生
	stu1 := Stu{Name: "tom", Age: 18, Telephone: "苹果"}
	stu2 := Stu{Name: "marry", Age: 28, Telephone: "安卓"}

	students["1"] = stu1
	students["2"] = stu2

	// 遍历学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是: %v\n", k)
		fmt.Printf("学生的名字是: %v\n", v.Name)
		fmt.Printf("学生的年龄是: %v\n", v.Age)
		fmt.Printf("学生的手机是: %v\n", v.Telephone)
		fmt.Println()
	}
}
