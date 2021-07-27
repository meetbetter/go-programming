package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// go run -race 08_atomic.go
func main() {
	var (
		wg  sync.WaitGroup
		num int64 = 0
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			atomic.AddInt64(&num, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			atomic.AddInt64(&num, 1)
		}
	}()

	wg.Wait()

	fmt.Printf("num:%d\n", num)

	fmt.Println("main return")
}
