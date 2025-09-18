/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

package main

import (
	"fmt"
	"math/rand"
)

func BubbleSort(nums *[5]int) {
	temp := 0 // 临时变量用于交换值
	for i := 0; i < len(*nums)-1; i++ {
		for j := 0; j < len(*nums)-1-i; j++ {
			if (*nums)[j] > (*nums)[j+1] {
				// 交换值
				temp = (*nums)[j]
				(*nums)[j] = (*nums)[j+1]
				(*nums)[j+1] = temp
			}
		}
	}
}

func binaryFind(nums *[5]int, leftIndex int, rightIndex int, findVal int) {
	/*
		判断leftIndex是否大于rightIndex
		如果大于则说明已经越过middle
	*/
	if leftIndex > rightIndex {
		fmt.Printf("找不到下标\n")
		return
	}

	// 找中间值的下标
	middle := (leftIndex + rightIndex) / 2

	if (*nums)[middle] > findVal {
		// 说明查找的值在leftIndex -> middle-1 之间
		binaryFind(nums, leftIndex, middle-1, findVal)
	} else if (*nums)[middle] < findVal {
		// 说明查找的数应该在 middle+1 -> rightIndex 之间
		binaryFind(nums, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("下标为: %v\n", middle) // 找到
	}
}

func main() {
	var findVal int
	// 定义初始数组
	nums := [5]int{}

	fmt.Printf("输入一个100以内的值: ")
	fmt.Scanln(&findVal)

	// 生成5个100以内的随机数
	for i := 0; i < 5; i++ {
		nums[i] = rand.Intn(99) + 1
	}

	// 冒泡排序
	BubbleSort(&nums)

	// 二分查找
	binaryFind(&nums, 0, len(nums)-1, findVal)
}
