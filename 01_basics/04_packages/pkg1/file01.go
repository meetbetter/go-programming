package pkg1

import "fmt"

func DoSomething() {
	fmt.Println("pkg1 DoSomething")
}

func init() {
	fmt.Printf("pkg1 file01 init is running, Data1:%d\n", Data1)
}
