package main

import (
	"fmt"
	"time"
)

func main() {
	nos := genEvenNos()
	for {
		evenNo, success := <-nos
		if !success {
			break
		}
		fmt.Println(evenNo)
	}
	fmt.Println("Done")
}

func genEvenNos() (nos chan int) {
	nos = make(chan int)
	count := 10
	go func() {
		for i := 0; i < count; i++ {
			time.Sleep(500 * time.Millisecond)
			nos <- i * 2
		}
		close(nos)
	}()
	return nos
}
