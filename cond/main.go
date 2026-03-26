package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {

	var wg sync.WaitGroup
	var mu = sync.Mutex{}
	var cond = sync.NewCond(&mu)

	wg.Add(1)

	go func() {
		defer wg.Done()
		//TODO suspend goroutine until sharedRsc is populated
		for len(sharedRsc) == 0 {
			time.Sleep(1 * time.Millisecond)
		}

		fmt.Println(sharedRsc["rsc1"])
	}()

	cond.L.Lock()
	sharedRsc["rsc1"] = "foo"
	cond.Signal()
	cond.L.Unlock()
	wg.Wait()
}
