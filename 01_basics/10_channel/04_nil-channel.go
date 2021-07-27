package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan struct{}

	go func() {
		ch <- struct{}{} //nil channel的读写永远阻塞
		fmt.Println("write channel successful")
	}()

	go func() {
		a := <-ch
		fmt.Println("get from channel:", a)
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("main end")

}
