// Copyright 2025 honeok <i@honeok.com>

package main

import (
	"fmt"
)

func main() {

	// 使用常规for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	// slice := arr[1:4] // 20 30 40
	slice := arr[:] // 20 30 40 (0-结尾)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v \n", i, slice[i])
	}

	fmt.Println()

	// 使用for--range遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	fmt.Println()

	/*
		切片可以继续切片
		slice已经是一个切片
	*/
	slice2 := slice[2:4] // 30 40
	slice2[0] = 100      // 同理slice[0]也会变化, 因为slice和slice2指向同一个数据空间, 因此slice2变化slice也会变化
	fmt.Println("slice2=", slice2)
}
