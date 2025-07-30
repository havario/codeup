package main

import (
	"fmt"
)

// 打印99乘法表
func printMulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println("")
	}
}

func main() {

	// 从终端输入一个整数表示要打印的乘法表对应的数
	var num int

	fmt.Printf("请输入九九乘法表对应的数: ")
	fmt.Scanln(&num)

	printMulti(num)
}
