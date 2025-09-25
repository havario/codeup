/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

/*
map切片

切片的数据类型如果是map, 则我们称为slice of map
map切片: 这样使用则map个数就可以动态变化
*/

package main

import (
	"fmt"
)

func main() {
	/*
	   要求: 使用一个map来记录monster的信息name 和 age
	   也就是说一个monster对应一个map, 并且妖怪的个数可以动态的增加 => map切片
	*/

	// 声明一个map切片
	monsters := make([]map[string]string, 2) // 准备放两个妖怪

	// 增加妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "250"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "红孩儿"
		monsters[1]["age"] = "251"
	}

	// 使用切片的append函数, 动态增加monsters
	// 先定义monsters信息
	newMonster := map[string]string{
		"name": "狐狸精",
		"age":  "300",
	}
	monsters = append(monsters, newMonster)

	for _, v := range monsters {
		fmt.Printf("%v\n", v)
	}
}
