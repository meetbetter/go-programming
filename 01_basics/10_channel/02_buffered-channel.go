package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch) // 关闭channel
	}()

	for {
		v, ok := <-ch // 关闭的channel，ok为false
		if !ok {
			fmt.Println("channel is closed")
			break
		}
		fmt.Printf("get %d from channel\n", v)
	}

	fmt.Println("main end")

}
