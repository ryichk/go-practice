package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 5*time.Second, errors.New("timeout caused by main"))
	defer cancel()
	taskA(ctx)
	taskB(ctx)
}

func taskA(ctx context.Context) {
	ctx, cancel := context.WithTimeoutCause(ctx, 1*time.Second, errors.New("timeout caused by taskA"))
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
	ctx, cancel := context.WithTimeoutCause(ctx, 1*time.Second, errors.New("timeout caused by taskB"))
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
		time.Sleep(3 * time.Second)
		close(out)
		fmt.Println("doSomething done")
	}()
	return out
}
