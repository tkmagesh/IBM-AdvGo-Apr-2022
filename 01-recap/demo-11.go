package main

import (
	"fmt"
	"time"
)

func main() {
	nos := genEvenNos()
	for {
		evenNo, success := <-nos
		if success {
			fmt.Println(evenNo)
			continue
		}
		break
	}
	fmt.Println("Done")
}

func genEvenNos() <-chan int {
	nos := make(chan int)
	count := 10
	go func(data chan<- int) {
		for i := 0; i < count; i++ {
			time.Sleep(500 * time.Millisecond)
			data <- i * 2
		}
		close(data)
	}(nos)
	return nos
}
