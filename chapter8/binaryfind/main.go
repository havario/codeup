package main

import (
	"fmt"
)

/*

arr := [6]int{1, 8, 10, 89, 1000, 1234} 8

二分法查找思路: 比如要查找的数是findVal
1. arr是一个有序数组, 并且从小到大排序
2. 先找到中间的下标 middle = (leftIndex+rightIndex)/2, 然后让中间下标的值和findVal比对

如果arr[middle] > findVal, 就应该向leftIndex -> (middle -1)
如果arr[middle] < findVal, 就应该向(middle +1) -> rightIndex
如果arr[middle] == findVal, 则找到

什么情况下, 就说明找不到 ?
if leftIndex > rightIndex {
	// 就说明找不到
	return
}

*/

func binaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	// 判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Printf("找不到\n")
		return
	}

	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		// 说明查找的数应该在 leftIndex -> middle-1 之间
		binaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 说明查找的数应该在 middle+1 -> rightIndex 之间
		binaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("下标为: %v\n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}

	// 测试
	binaryFind(&arr, 0, len(arr)-1, -1)
}
