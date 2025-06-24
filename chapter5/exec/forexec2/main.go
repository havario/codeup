// Copyright (c) 2025 honeok <honeok@disroot.org>

package main

import (
	"fmt"
)

func main() {
	// 统计3个班成绩情况, 每个班有5名学生
	// 求出各个班的平均分和所有班级的平均分 [学生的成绩从键盘输入]

	// 代码做活
	// 定义两个变量,表示班级的个数和班级的人数
	var classNum int = 3
	var stuNum int = 5

	var totalSum float64 = 0.0
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入%d班的第%d个学生的成绩: ", j, i)
			fmt.Scanln(&score)
			// 累计学生总分
			sum += score
		}
		fmt.Printf("第%d个班的平均分是%v\n", j, sum/float64(stuNum))
		// 将各个班的总成绩累计到totalSum
		totalSum += sum
	}
	fmt.Printf("各个班级的总成绩%v 各个班平均分是%v\n", totalSum, totalSum/float64(stuNum*classNum))
}
