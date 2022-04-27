package main

import (
	"fmt"
	"time"
)

func main() {
	done := func() chan bool {
		doneCh := make(chan bool)
		go func() {
			var input string
			fmt.Scanln(&input)
			doneCh <- true
		}()
		return doneCh
	}()

	nos := genEvenNos(done)
	for evenNo := range nos {
		fmt.Println(evenNo)
	}
	fmt.Println("Done")
}

func genEvenNos(done chan bool) (nos chan int) {
	nos = make(chan int)
	go func() {
		i := 0
	LOOP:
		for {
			select {
			case nos <- i * 2:
				time.Sleep(500 * time.Millisecond)
				i++
			case <-done:
				break LOOP
			}
		}
		close(nos)
	}()
	return nos
}
