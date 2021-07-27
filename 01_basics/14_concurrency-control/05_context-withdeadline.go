package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(time.Second * 3)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go process(ctx)

	time.Sleep(10 * time.Second)

	fmt.Println("main return")

}

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("timeout or cancel, process return:%v\n", ctx.Err())
			return
		default:
			fmt.Println("业务process is running...")
			time.Sleep(time.Second)
		}
	}
}
