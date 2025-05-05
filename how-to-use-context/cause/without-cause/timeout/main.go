package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	taskA(ctx)
	taskB(ctx)
}

func taskA(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	fmt.Println("start taskA...")

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-doSomething():
		fmt.Println("taskA done")
	}
}

func taskB(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	fmt.Println("start taskB...")

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-doSomething():
		fmt.Println("taskB done")
	}
}

func doSomething() <-chan struct{} {
	out := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(out)
		fmt.Println("doSomething done")
	}()
	return out
}
