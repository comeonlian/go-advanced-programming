package main

import (
	"fmt"
	"sync"
	"time"
)

func workerMulti(wg *sync.WaitGroup, cancel chan int) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancel:
			return
		}
	}
}

func main() {
	cancel := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerMulti(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)

	wg.Wait()
}
