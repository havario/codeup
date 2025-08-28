package main

import "fmt"

func fibonacci(n int) []uint64 {
	// 声明切片
	fibonacciSlice := make([]uint64, n)
	// 第一个数和第二个数斐波那契为1
	fibonacciSlice[0] = 1
	fibonacciSlice[1] = 1
	for i := 2; i < n; i++ {
		fibonacciSlice[i] = fibonacciSlice[i-1] + fibonacciSlice[i-2]
	}

	return fibonacciSlice
}

func main() {
	/*
		1> 可以接受一个 n int
		2> 能够将斐波那契数列放在切片中
		3> 提示, 斐波那契数列形式
		arr[0] = 1; arr[1] = 1; arr[2] = 2; arr[3] = 3; arr[4] = 5; arr[5] = 8;

		思路:
		1. 声明一个函数 fibonacci(n int) ([]uint64)
		2. 编写 fibonacci(n int) 进行for循环存放斐波那契数列
		3. 单独提取 0 和 1
	*/
	fibonacci := fibonacci(10)
	fmt.Printf("fibonacci=%v\n", fibonacci)
}
