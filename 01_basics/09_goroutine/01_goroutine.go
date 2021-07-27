package main

import (
	"fmt"
	"time"
)

func main() {

	go func() { //多个协程并发执行
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine 1 is running")
			time.Sleep(time.Second * 1)
		}
		fmt.Println("goroutine exit")
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine 2 is running")
			time.Sleep(time.Second * 1)
		}
		fmt.Println("goroutine exit")
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("main is running")
	}

	fmt.Println("main exit")
}
