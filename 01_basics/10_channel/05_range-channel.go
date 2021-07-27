package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		time.Sleep(time.Second * 5)
		close(ch) // 关闭channel
	}()

	for v := range ch { // 遍历channel，直到channel close
		fmt.Printf("get %d from channel\n", v)
	}

	fmt.Println("main end")
}
