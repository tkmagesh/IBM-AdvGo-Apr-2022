package main

import (
	"fmt"
)

func main() {

	fmt.Println("main started")
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch //receive operation
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result //send operation
}
