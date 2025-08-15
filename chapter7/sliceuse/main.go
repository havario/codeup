package main

import (
	"fmt"
)

func main() {
	// 切片的使用 --> make方法

	/*
		对于切片而言, 必须make使用

		1> 通过make方式创建切片可以指定切片的大小和容量
		2> 如果没有给切片各个元素赋值, 那么就会使用默认值 [int, float => 0  string => "" bool => false]
		3> 通过make方式创建的切片对应的数组是由make底层维护对外不可见, 即只能通过slice去访问各个元素
	*/
	var slice []float64 = make([]float64, 5, 10)
	fmt.Printf("slice的值: %v\n", slice)
	slice[1] = 10
	slice[3] = 20
	fmt.Printf("slice的值: %v\n", slice)
	fmt.Printf("slice的Size: %v\n", len(slice))
	fmt.Printf("slice的容量: %v\n", cap(slice))

	fmt.Println("====================")

	/*
		定义一个切片, 直接指定具体数组
	*/
	var strSlice []string = []string{"Lisa", "Jisoo", "Jennie", "Rosé"}
	fmt.Printf("strSlice的值: %v\n", strSlice)
	fmt.Printf("strSlice的Size: %v\n", len(strSlice))
	fmt.Printf("strSlice的容量: %v\n", cap(strSlice))
}
