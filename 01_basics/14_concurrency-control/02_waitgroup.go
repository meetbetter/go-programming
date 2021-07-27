package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("第%d号goroutine return\n", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("main return")
}
