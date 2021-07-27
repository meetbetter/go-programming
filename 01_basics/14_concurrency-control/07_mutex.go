package main

import (
	"fmt"
	"sync"
)

// go run -race 07_mutex.go
func main() {
	var (
		wg sync.WaitGroup
		m  = make(map[string]int)
		mu sync.Mutex
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		defer mu.Unlock()
		mu.Lock()
		fmt.Println("goroutine 1 is write")
		m["zhangsan"] = 100
	}()

	go func() {
		defer wg.Done()
		defer mu.Unlock()
		mu.Lock()
		fmt.Println("goroutine 2 is write")
		m["zhangsan"] = 200
	}()

	wg.Wait()

	fmt.Printf("m:%v\n", m)
	fmt.Println("main return")
}
