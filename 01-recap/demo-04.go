package main

import (
	"fmt"
	"sync"
)

/* Communicate by sharing memory (NOT ADVICABLE) */
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("main started")
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = x + y
}
