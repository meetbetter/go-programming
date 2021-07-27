package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	commonChannel(ch)

	time.Sleep(time.Second)

	go readChannel(ch)

	go writeChannel(ch)

	time.Sleep(time.Second)

}

func commonChannel(ch chan int) { // 可读可写channel
	go func() {
		ch <- 10
		fmt.Printf("commonChannel write %d to channel \n", 10)
	}()
	fmt.Printf("commonChannel get from channel %d\n", <-ch)
}

func readChannel(ch <-chan int) { //只读channel
	fmt.Printf("readChannel read %d from channel\n", <-ch)
}

func writeChannel(ch chan<- int) { //只写channel
	ch <- 100
	fmt.Printf("writeChannel write %d to channel\n", 100)
}
