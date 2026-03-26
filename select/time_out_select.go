package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	select {
	case v := <-ch:
		fmt.Printf("v %v", v)
	case <-time.After(1 * time.Second):
		fmt.Printf("time out")
	}

}
