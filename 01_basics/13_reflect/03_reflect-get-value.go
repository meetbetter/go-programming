package main

import (
	"fmt"
	"reflect"
)

func getValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	default:
		fmt.Println("unknown type kind")
	}
}
func main() {
	var a float32 = 1.234
	var b int64 = 18
	var c int = 20
	getValue(a) // type is float32, value is 1.234
	getValue(b) // type is int64, value is 18
	getValue(c)
}
