package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerP(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("do something")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerP(ctx, &wg)
	}

	time.Sleep(10 * time.Second)
	cancel()

	wg.Wait()
}
