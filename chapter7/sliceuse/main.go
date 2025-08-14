package main

import (
	"fmt"
)

func main() {
	// 切片的使用 --> make方法

	/*
		对于切片而言, 必须make使用
	*/
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Printf("%v\n", slice)
	fmt.Printf("slice的Size: %v\n", len(slice))
	fmt.Printf("slice的容量: %v\n", cap(slice))
}
