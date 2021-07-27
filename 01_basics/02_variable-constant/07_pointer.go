package main

import "fmt"

func main() {

	var a int = 10
	change1(a)
	fmt.Println("after change1:", a)
	change2(&a)
	fmt.Println("after change2:", a)
}

func change1(b int) {
	b = 100
}

func change2(b *int) {
	*b = 100
}
