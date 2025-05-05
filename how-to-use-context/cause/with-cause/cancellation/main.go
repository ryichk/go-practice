package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	wg.Add(1)
	go task(ctx)

	// do something

	cancel(errors.New("canceled by main func"))
	wg.Wait()
}

func task(ctx context.Context) {
	defer wg.Done()

	ctx, cancel := context.WithCancelCause(ctx)
	wg.Add(1)
	go subTask(ctx)

	// do something

	cancel(errors.New("canceled by task func"))
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
		time.Sleep(time.Second)
		close(out)
		fmt.Println("doSomething done")
	}()
	return out
}
