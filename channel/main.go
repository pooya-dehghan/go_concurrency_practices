package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case v1 := <-ch1:
			fmt.Printf("v1 %v \n", v1)
		case v2 := <-ch2:
			fmt.Printf("v2 %v \n", v2)
		}
	}

}
