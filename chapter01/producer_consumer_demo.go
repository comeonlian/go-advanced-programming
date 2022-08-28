package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(5, ch)

	go Consumer(ch)

	// time.Sleep(1 * time.Second)
	sig := make(chan os.Signal, 1) // Ctrl+C 退出
	signal.Notify(sig, syscall.SIGINT, syscall.SIGINT)
	fmt.Printf("quit (%v)\n", <-sig)
}
