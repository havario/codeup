package main

import (
	"fmt"
)

func main() {
	// var 数组名 [数组大小]数据类型
	var intArr [3]int // 8个字节

	// 当定义数组后, 其实数组的各个元素有默认值0
	fmt.Printf("%v\n", intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Printf("%v\n", intArr)
	/*
		1. 数组的地址可以通过数组名获取 &intArr
		2. 数组的第一个元素地址就是数组的首地址
		3. 数组的各个元素的地址间隔是依据数组的类型决定, 比如int64->8 int32->4

		%p: 输出指针的值
	*/
	fmt.Printf("intArr的地址: %p intArr[0]地址: %p intArr[1]地址: %p\n",
		&intArr, &intArr[0], &intArr[1])

	// 从终端循环输入5个成绩, 保存到float64数组并输出
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值: ", i+1) // 元素下标为0就是第一个元素, %d把整数按十进制输出
		fmt.Scanln(&score[i])
	}

	// 遍历数组打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}

	// 四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Printf("numArr01: %v\n", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Printf("numArr02: %v\n", numArr02)

	var numArr03 = [...]int{8, 9, 10} // 这里的...是固定写法
	fmt.Printf("numArr03: %v\n", numArr03)

	numArr04 := [...]int{1: 800, 0: 900, 2: 999}
	fmt.Printf("numArr04: %v\n", numArr04)
}
