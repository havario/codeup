/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

/*
map遍历 for range
*/

package main

import (
	"fmt"
)

func main() {
	citys := make(map[string]string)
	citys["1"] = "北京"
	citys["2"] = "天津"
	citys["3"] = "上海"
	fmt.Printf("%v\n", citys)

	for k, v := range citys {
		fmt.Printf("key=%v value=%v\n", k, v)
	}

	fmt.Printf("citys Map 有 %v 对kv\n", len(citys))

	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["gender"] = "男"
	studentMap["stu01"]["address"] = "北京长安街"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Mary"
	studentMap["stu02"]["gender"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江"

	for k1, v1 := range studentMap {
		fmt.Printf("%v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
