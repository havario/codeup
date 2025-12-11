package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Server start.\n")

	// 监听
	listen,err := net.Listen("tcp","192.168.0.189:38438")
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	// 监听成功 等待客户端连接
	for {
		conn,err2 := listen.Accept()
		if err2 != nil {
			fmt.Printf("Waiting for client connection failed.: %v\n", err2)
		} else {
			// 连接成功
			fmt.Printf("Connection success: %v Client is: %v\n",conn,conn.RemoteAddr().String())
		}
	}
}
