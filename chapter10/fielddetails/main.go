/*
Copyright (c) 2025 honeok ♡ Rosé <i@honeok.com>
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
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // map
}

type Monster struct {
	Name string
	Age  int
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

	// 使用slice 先make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println(p1.slice)

	// 使用map 同样先make
	p1.map1 = make(map[string]string, 5)
	p1.map1["beauty1"] = "Rosé"
	p1.map1["beauty2"] = "Jennie"
	fmt.Println(p1.map1)

	// 不同结构体变量的字段是独立互不影响, 一个结构体变量字段的更改不影响另外一个, 结构体是值类型
	var monster1 Monster
	monster1.Name = "牛魔酬宾"
	monster1.Age = 500

	monster2 := monster1
	monster2.Name = "喷嚏精"                // 修改monster2的Name验证结构体是值类型 默认为值拷贝
	fmt.Println("monster1 = ", monster1) // monster1 =  {牛魔酬宾 500}
	fmt.Println("monster2 = ", monster2) // monster2 =  {喷嚏怪 500}
}
