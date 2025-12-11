package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Printf("Client Start.\n")

	// 调用Dial函数
	conn,err := net.Dial("tcp","192.168.0.189:38438")
	if  err != nil {
		fmt.Printf("Client connection failed: %v \n",err)
		return
	}
	fmt.Printf("Connection success: %v\n",conn)
}
