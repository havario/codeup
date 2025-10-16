// Copyright (c) honeok 2025 <i@honeok.com>
// SPDX-License-Identifier: Apache-2.0

// 结构体的所有字段在内存中是连续分布的

package main

import (
	"fmt"
)

// 结构体
type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	// r1有4个整数int, 在内存中是连续分布的
	// 打印r1的地址
	fmt.Printf("r1.leftUp.x地址=%p r1.leftUp.y地址=%p r1.rightDown.x地址=%p r1.rightDown.y地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// r2有两个 *Point类型, 这两个*Point类型的本身地址也是连续的, 但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp 本身地址=%p r2.rightDown 本身地址=%p\n",
		&r2.leftUp, &r2.rightDown)

	// 他们指向的地址不一定是连续的, 要看运行时系统内存分布
	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指向地址=%p\n",
		r2.leftUp, r2.rightDown)
}
