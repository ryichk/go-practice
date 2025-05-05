package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go task(ctx)

	// do something

	cancel()
	wg.Wait()
}

func task(ctx context.Context) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go subTask(ctx)

	// do something

	cancel()
}

func subTask(ctx context.Context) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-doSomething():
		fmt.Println("subtask done")
	}
}

func doSomething() <-chan struct{} {
	out := make(chan struct{})
	go func() {
		defer close(out)
		time.Sleep(time.Second)
		fmt.Println("doSomething done")
	}()
	return out
}
