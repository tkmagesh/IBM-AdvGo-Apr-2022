package main

import (
	"fmt"
	"time"
)

func main() {
	nos := genEvenNos()
	for evenNo := range nos {
		fmt.Println(evenNo)
	}
	fmt.Println("Done")
}

func genEvenNos() (nos chan int) {
	nos = make(chan int)
	timeoutCh := timeout(10 * time.Second)
	go func() {
		i := 0
	LOOP:
		for {
			select {
			case nos <- i * 2:
				time.Sleep(500 * time.Millisecond)
				i++
			case <-timeoutCh:
				break LOOP
			}
		}
		close(nos)
	}()
	return nos
}

func timeout(t time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(t)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
