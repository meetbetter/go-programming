package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- true
	}()

	a := <-ch
	fmt.Println("get from channel:", a)

	fmt.Println("main end")

}
