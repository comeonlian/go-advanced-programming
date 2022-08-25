package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var totalVal uint64 = 0

func workerProc(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		atomic.AddUint64(&totalVal, 1)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go workerProc(&wg)
	go workerProc(&wg)

	wg.Wait()

	fmt.Println(totalVal)
}
