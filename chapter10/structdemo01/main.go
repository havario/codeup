/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

package main

import (
	"fmt"
)

/*
张老太养了两只猫猫:
一只名字叫小白今年3岁——白色
还有一只叫小花今年100岁——花色
请编写一个程序, 当用户输入小猫的名字时, 就显示该猫的名字年龄、颜色
如果用户输入的小猫名错误, 则显示张老太没有这只猫猫。
*/

/*
定义cat结构体, 将cat的字段放入结构体
在go里结构体struct是值类型
*/

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

/*
结构体是自定义的数据类型, 代表一类事物
结构体变量 (实例) 是具体的、实际的, 代表一个具体变量
*/

func main() {
	// 创建Cat变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "fish"
	fmt.Printf("cat1地址为: %p\n", &cat1)
	fmt.Printf("cat1 = %v\n", cat1)
}
