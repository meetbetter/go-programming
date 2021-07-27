package main

import (
	"fmt"
)

func main() {
	defer fmt.Println(" main 1111")
	defer fmt.Println("main 2222")
	f1()
	defer fmt.Println("main 3333")
	defer fmt.Println("main 4444")
}

func f1() {
	defer fmt.Println("f1 defer 1")
	defer fmt.Println("f1 defer 2")
}
