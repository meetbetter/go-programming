package main

import "fmt"

func main() {
	defer fmt.Println("main entry")

	f1()

	f2()

	fmt.Println("main End")
}

func f1() {
	defer fmt.Println("f1 enrty")

	defer func() { // recover要配合defer使用才生效
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()

	panic("boom")

	fmt.Println("看不到我")
}

func f2() {
	fmt.Println("f2")
}
