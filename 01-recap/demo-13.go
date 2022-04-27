package main

import (
	"fmt"
	"time"
)

func main() {
	evenCh := func() chan int {
		evenCh := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				time.Sleep(100 * time.Millisecond)
				evenCh <- i * 2
			}
			close(evenCh)
		}()
		return evenCh
	}()

	oddCh := func() chan int {
		oddCh := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				time.Sleep(1000 * time.Millisecond)
				oddCh <- (i * 2) + 1
			}
			close(oddCh)
		}()
		return oddCh
	}()

	/*
		for i := 0; i < 10; i++ {
			evenNo := <-evenCh
			fmt.Println("Even -", evenNo)
			oddNo := <-oddCh
			fmt.Println("Odd -", oddNo)
		}
	*/

	for {
		select {
		case oddNo := <-oddCh:
			fmt.Println("Odd -", oddNo)
		case evenNo := <-evenCh:
			fmt.Println("Even -", evenNo)
		}
	}

}
