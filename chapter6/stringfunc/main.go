// Copyright (c) 2025 honeok <i@honeok.com>
//
// 字符串中常用的系统函数

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度, 按字节
	str1 := "hello北"                    // golang编码统一为utf-8, ascii码的字符(字母和数字) 占一个字节, 汉字占用3个字节
	fmt.Println("str len= ", len(str1)) // 打印8 5+3

	// 字符串遍历, 同时处理有中文的问题 r := []rune(str)
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数: n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果", n)
	}

	// 整数转字符串
	str1 = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str1, str1)

	// 字符串转 []byte var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// byte转字符串 str = string([]byte{97, 98, 99})
	str1 = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str1) // 转为ascii码

	// 10进制转 2 8 16进制 str = strconv.FormatInt(123, 2) , 返回对应的字符串
	str1 = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的2进制:%v\n", str1)

	str1 = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制:%v\n", str1)

	// 查找子串是否在指定的字符串中: strings.Contains("seafood","foo") // 返回bool值 true or false
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b)

	// 统计一个字符串有几个指定的子串 strings.Contains("ceheese", "e") // 4
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	// 不区分大小写的字符串比较(==是区分字母大小写的) fmt.Println(strings.EqualFold("abc", "Abc"))
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b) // true

	// 返回子串在字符串第一次出现的index值, 如果没有返回 -1 strings.Index("NLT_abc", "abc") // 4
	index := strings.Index("NLT_abc", "abc")
	//                下标: 01234
	fmt.Printf("index=%v\n", index) // 4
}
