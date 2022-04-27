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

			oddNo := <-oddCh
			fmt.Println("Odd -", oddNo)
		}
	*/
	evenDone := func() chan bool {
		done := make(chan bool)
		go func() {
			for evenNo := range evenCh {
				fmt.Println("Even -", evenNo)
			}
			done <- true
		}()
		return done
	}()

	oddDone := func() chan bool {
		done := make(chan bool)
		go func() {
			for oddNo := range oddCh {
				fmt.Println("Odd -", oddNo)
			}
			done <- true
		}()
		return done
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-oddDone:
			fmt.Println("Odd numbers done")
		case <-evenDone:
			fmt.Println("Even numbers done")
		}
	}

}
