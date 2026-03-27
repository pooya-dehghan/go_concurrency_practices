package main

import "fmt"

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

func main() {
	for v := range square(generator(1, 2, 3, 4)) {
		fmt.Printf("%v \n", v)
	}
}
