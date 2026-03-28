package main

import (
	"context"
	"fmt"
)

func generator(ctx context.Context) <-chan int {
	out := make(chan int)
	intge := []int{1, 2, 3, 4, 5, 6, 7, 8}

	go func() {
		defer close(out)
		for _, v := range intge {
			select {
			case out <- v:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := generator(ctx)

	for v := range ch {
		fmt.Printf("value %v \n", v)
		if v == 5 {
			cancel()
		}
	}

}
