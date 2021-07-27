package main

import "fmt"

func f1() {
	defer fmt.Println("f1 entry")
	panic("boom") // panic后会执行当前goroutine上的defer
	fmt.Println("f1 end")
}
func main() {
	defer fmt.Println("main Entry")
	f1()
	fmt.Println("main End")
}
