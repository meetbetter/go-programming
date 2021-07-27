package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("收到停止信号，退出")
				return
			default:
				fmt.Println("goroutine业务运行中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	fmt.Println("通过channel通知goroutine停止")
	stop <- true

	time.Sleep(5 * time.Second)
	fmt.Println("main end")
}
