/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

package main

import (
	"fmt"
)

/*
如果结构体的字段类型为: 指针 slice map 的零值都是nil 即还没有分配空间
如果需要使用这样的字段, 需要先make才能使用
*/

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int  // 指针
	slice  []int // 切片
	map1   map[string]string
}

func main() {
	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok")
	}
	if p1.slice == nil {
		fmt.Println("ok")
	}
	if p1.map1 == nil {
		fmt.Println("ok")
	}

	// 使用slice
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println(p1)
}
