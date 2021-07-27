package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go process(ctx)

	// time.Sleep(2 * time.Second)
	// cancel()

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
