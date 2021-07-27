package main

import "fmt"

func f() (result int) {
	defer func() {
		result++
	}()
	return 100
}

func main() {
	fmt.Println(f())
}
