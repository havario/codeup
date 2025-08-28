package main

import (
	"fmt"
)

func main() {

	// string底层是一个byte数组, 因此string也可以进行切片处理
	str := "i@honeok.com"
	// 使用切片获取honeok.com
	slice := str[2:]
	fmt.Println("slice=", slice)

	/*
		string是不可变的, 也就说不能通过 str[0]='z' 方式来修改字符串
		如果需要修改字符串, 可以先将 string -> []byte 或者 []rune -> 修改 -> 重写转成 string
		i@honeok.com -> h@honeok.com
	*/
	arr1 := []byte(str)
	arr1[0] = 'h'
	str = string(arr1)
	fmt.Println("str=", str)

	/*
		细节: 转为byte数组后可以处理英文和数字, 但是不能处理中文
		原因: []byte安装字节处理, 而汉字是3个字节因此就会出现乱码
		解决: 将str转  []rune解决
	*/
	arr2 := []rune(str)
	arr2[0] = '贺'
	str = string(arr2)
	fmt.Println("str=", str)
}
