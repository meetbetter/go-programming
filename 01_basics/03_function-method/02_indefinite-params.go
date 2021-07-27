package main

import "fmt"

func foo(arr ...int) { // 不定参的实参个数可以为任意个
	fmt.Printf("%d, %v\n", len(arr), arr)
}

func main() {
	a, b, c := 1, 2, 3

	foo()
	foo(a, b, c)
	foo(a, b, c, 4, 5, 6)

	sli := make([]int, 10, 20)
	foo(sli...) // slice作为不定参时需要...
}
