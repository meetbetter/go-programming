package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(process)
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("main end")
}

func process() {
	fmt.Println("process is running")
}
