package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go process(ctx)

	time.Sleep(10 * time.Second)

	cancel()

	fmt.Println("main return")
}

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到外部停止信号，退出业务goroutine")
			return
		default:
			fmt.Println("业务goroutine执行中...")
			time.Sleep(time.Second)
		}
	}
}
