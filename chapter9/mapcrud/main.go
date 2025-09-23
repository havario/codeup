/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

/*
map增删改查
*/

package main

import (
	"fmt"
)

func main() {
	citys := make(map[string]string)
	citys["1"] = "BeiJing"
	citys["2"] = "TianJing"
	citys["3"] = "ShangHai"
	fmt.Printf("%v\n", citys)

	/*
		map的修改
		因为3这个key存在, 因为下面就是修改
	*/
	citys["3"] = "ShangHai~"
	fmt.Printf("%v\n", citys)

	/*
		map的删除
		func delete(m map[Type]Typel, key Type)
		内建函数delete按照指定的键将元素从映射中删除, 若m为nil或无此元素delete不进行操作

		delete(map, "key") delete是一个内置函数, 如果key存在, 就删除该key-value, 如果key不存在, 不操作但是也不会报错
	*/
	delete(citys, "1")
	fmt.Printf("%v\n", citys)

	// 演示删除
	delete(citys, "4")
	fmt.Printf("%v\n", citys)

	// 查询
	val, ok := citys["2"]
	if ok {
		fmt.Printf("有key: %v\n", val)
	} else {
		fmt.Printf("没有key\n")
	}

	/*
		如果希望一次性删除所有的key
		1. 遍历所有的key, 遍历逐一删除
		2. 直接make一个新的空间
	*/
	citys = make(map[string]string)
	fmt.Printf("%v\n", citys)
}
