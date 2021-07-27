package pkg1

import "fmt"

// 全局常量和变量在init前执行
const Data1 = 100

func init() {
	fmt.Printf("pkg1 file02 init is running, Data1:%d\n", Data1)
}
