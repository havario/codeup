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
	str1 := "hello北"                   // golang编码统一为utf-8, ascii码的字符(字母和数字) 占一个字节, 汉字占用3个字节
	fmt.Println("str len=", len(str1)) // 打印8 5+3

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

	fmt.Println("结果", "abc" == "Abc")

	// 返回子串在字符串第一次出现的index值, 如果没有返回 -1 strings.Index("NLT_abc", "abc") // 4
	index := strings.Index("NLT_abc", "abc")
	//                下标: 01234
	fmt.Printf("index=%v\n", index) // 4

	// 返回子串在字符串最后一次出现的index
	// 如没有返回-1: strings.LastIndex("go golang", "go")
	index = strings.LastIndex("go golang", "go") // 3
	fmt.Printf("index=%v\n", index)

	// 将指定的子串替换成另外一个子串: strings.Replace("go go hello", "go", "go语言", n)
	// n可以指定你希望替换几个，如果n=-1表示全部替换
	str1 = "go go hello"
	str2 = strings.Replace(str1, "go", "BeiJing", 1)
	fmt.Printf("str1=%v\n", str2)

	// 按照指定的某个字符为分割标识, 将一个字符串拆分成字符串数组 strArr := strings.Split("hello,world,ok", ",")
	strArr := strings.Split("hello,world,ok", ",") // 按,号进行拆分
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v", i, strArr[i])
		fmt.Printf("\n")
	}
	fmt.Printf("strArr=%v\n", strArr)

	// 将字符串的字母进行大小写的转换: strings.ToLower("Go")
	str1 = "golang Hello"
	str1 = strings.ToLower(str1)  // 全小写
	fmt.Printf("str1=%v\n", str1) // GOLANG HELLO
	str1 = strings.ToUpper(str1)  // 全大写
	fmt.Printf("str1=%v\n", str1) // golang hello

	// 将字符串左右两边的空格去掉 strings.TrimSpace(" tn a lone gopher ntrn ")
	str1 = strings.TrimSpace(" tn a lone gopher ntrn ")
	fmt.Printf("str1=%q\n", str1)

	// 将字符串左右两边指定的字符去掉 strings.Trim("! hello! ", " !")
	str1 = strings.Trim("! he!llo! ", " !") // ["hello"]  将左右两边 ! 和 "" 去掉, 因为写了一个空格和一个!号
	fmt.Printf("str1=%q\n", str1)
	// strings.TrimLeft() // 去左
	// strings.TrimRight() // 去右

	// 判断字符串是否以指定的字符串开头 strings.HasSuffix("ftp://192.168.10.1", "ftp")
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b) // true

	// 判断字符串是否以指定的字符串结束
	b = strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Printf("b=%v\n", b) // false
}
