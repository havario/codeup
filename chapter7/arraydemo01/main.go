package main

import "fmt"

func main() {
	/*
		一个养鸡场有6只鸡, 它们的体重分别是3kg 5kg 1kg 3.4kg 2kg 50kg 请问这六只鸡的总体重是多少? 平均体重是多少?
		请你编一个程序 => 数组
	*/

	// 使用数组的方式来解决问题

	// 1. 定义一个数组
	var hens [6]float64
	// 2. 给数组每个元素赋值
	hens[0] = 3.0 // hens数组的第一个元素赋值 hens[0]
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	// 3. 遍历数组求出总体重
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	// 4. 求出平均体重
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)
}
