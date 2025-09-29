/*
Copyright (c) 2025 honeok <i@honeok.com>
SPDX-License-Identifier: MIT
*/

/*
1> 使用 map[string]map[string]sting 的map类型
2> key: 表示用户名是唯一的, 不可以重复
3> 如果某个用户名存在, 就将其密码修改 "888888" 如果不存在就增加这个用户信息: 包括昵称nickname和密码pwd
4> 编写—个函数 modifyUser(users map[string]map[string]string, name string) 完成上述功能
*/

package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string, nickname string) {

	// 判断users中是否存在name
	if users[name] != nil {
		// 用户存在
		users[name]["pwd"] = "888888"
	} else {
		// 用户不存在
		users[name] = make(map[string]string, 2)
		users[name]["nickname"] = nickname
		users[name]["pwd"] = "888888"
	}
}

func main() {
	users := make(map[string]map[string]string, 10)
	users["Jisoo"] = make(map[string]string, 2)

	users["Jisoo"]["pwd"] = "999999"
	users["Jisoo"]["nickname"] = "Kim"

	modifyUser(users, "Rosé", "Rosé")
	modifyUser(users, "Jennie", "Jennie")
	modifyUser(users, "Jisoo", "Jisoo")

	fmt.Printf("%v\n", users)
}
