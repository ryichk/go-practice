package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("ctx cleanup done")
	})
	defer stop()

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case <-ticker.C:
				fmt.Println("tick")
			}
		}
	}()

	// do something

	cancel()
}
