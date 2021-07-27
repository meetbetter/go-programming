package main

import "fmt"

func f1() {
	fmt.Println("f1 entry")
	panic("boom")
	fmt.Println("f1 end")
}
func main() {
	fmt.Println("main Entry")
	f1()
	fmt.Println("main End")
}
