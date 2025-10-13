/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式2
	p2 := Person{"rose", 20}
	fmt.Println(p2)

	// 方式 3 和 4 都是结构体指针
	// 方式3
	var p3 *Person = new(Person)
	(*p3).Name = "rose"
	(*p3).Age = 21
	fmt.Println(*p3)

	// 方式4
	var person *Person = &Person{}
	/*
		因为person是一个指针, 因此标准访问字段的方法是
		(*person).Name = "scott"
		go的设计者为了程序员使用方便, 也可以 person.Name = "scott"
		go底层会对person.Name = "scott" 加(*person)
	*/
	(*person).Name = "scott"
	person.Name = "rose"

	(*person).Age = 88
	person.Age = 18
	fmt.Println(*person)
}
