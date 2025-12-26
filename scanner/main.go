package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup

	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			address := fmt.Sprintf("192.168.0.1:%d", k)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%s open!\n", address)
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(startTime) / 1e9
	fmt.Printf("\n\n%d Seconds", elapsed)
}
