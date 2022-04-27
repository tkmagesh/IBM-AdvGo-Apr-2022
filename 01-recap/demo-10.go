package main

import (
	"fmt"
	"time"
)

func main() {
	nos := genEvenNos(10)
	for i := 0; i < 10; i++ {
		evenNo := <-nos
		fmt.Println(evenNo)
	}
	fmt.Println("Done")
}

func genEvenNos(count int) (nos chan int) {
	nos = make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			time.Sleep(500 * time.Millisecond)
			nos <- i * 2
		}
	}()
	return nos
}
