package pkg2

import (
	"fmt"
)

func DoSomething() {
	fmt.Println("pkg2 DoSomething")
}

func init() {
	fmt.Println("pkg2 init is running")
}
