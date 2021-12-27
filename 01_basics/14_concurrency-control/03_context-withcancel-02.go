package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go process1(ctx)

	time.Sleep(5 * time.Second)

	cancel()

	time.Sleep(5 * time.Second)

	fmt.Println("main return")
}

func process1(ctx context.Context) {
	go process2(ctx) //传递context给其他goroutine

	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到外部停止信号，退出业务goroutine 1")
			return
		default:
			fmt.Println("业务goroutine 1 执行中...")
			time.Sleep(time.Second)
		}
	}
}

func process2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到外部停止信号，退出业务goroutine 2")
			return
		default:
			fmt.Println("业务goroutine 2 执行中...")
			time.Sleep(time.Second)
		}
	}
}
