package main

import "fmt"

func main() {
	defer fmt.Println("main entry")
	foo()
	fmt.Println("main return")
}

func foo() {
	defer fmt.Println("foo entry")
	defer fmt.Println("foo defer2") // defer后进先出
	defer fmt.Println("foo defer3")

	fmt.Println("foo return")
}
