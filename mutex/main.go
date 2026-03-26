package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(8)
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()

		balance += amount
		mu.Unlock()
	}

	withdrawal := func(amount int) {
		mu.Lock()

		balance -= amount
		mu.Unlock()
	}

	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()

	fmt.Printf("balance is %v", balance)
}
