package main

import (
	"fmt"
	"github.com/comeonlian/go-advanced-programming/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	allMsg := p.Subscribe()
	golangMsg := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range allMsg {
			fmt.Println("all: ", msg)
		}
	}()

	go func() {
		for msg := range golangMsg {
			fmt.Println("golang: ", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
