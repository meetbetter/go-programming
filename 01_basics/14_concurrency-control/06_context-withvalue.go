package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "daydayup"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, key, "https://github.com/meetbetter/go-programming")

	defer cancel()

	go process(valCtx)

	time.Sleep(5 * time.Second)

	fmt.Println("main return")

}

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("timeout or cancel, process return:%v\n", ctx.Err())
			return
		default:
			fmt.Printf("业务process is running, context key:%s,value:%s\n", key, ctx.Value(key))
			time.Sleep(time.Second)
		}
	}
}
