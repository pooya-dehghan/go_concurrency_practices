package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)

	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	publisher := func(c <-chan int) {
		for vc := range c {
			out <- vc
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go publisher(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch := generator(1, 2, 3, 4)

	sqCh1 := square(ch)
	sqCh2 := square(ch)

	mCh := merge(sqCh1, sqCh2)

	for v := range mCh {
		fmt.Printf("%v \n", v)
	}
}
