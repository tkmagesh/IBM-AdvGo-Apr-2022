package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go fn(i) //fn is scheduled to be executed in future
	}
	wg.Wait()
	fmt.Println("main completed")
}

func fn(i int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("fn -", i)
	return
}
