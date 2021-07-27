package main

import "fmt"

func main() {
	a := 100

	switch { // 选择对应case执行，case都不成立则执行default，default可以不存在，
	case a < 0:
		fmt.Println("a is negative") // go中switch 执行完会自动break跳出switch,若想继续执行下面的case则手动fallthrough
		fallthrough
	case a > 0:
		fmt.Println("a is positive")
	default:
		fmt.Println("a is zero")
	}
}
