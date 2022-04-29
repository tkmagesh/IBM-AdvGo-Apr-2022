package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(no int) {
				defer wg.Done()
				time.Sleep(time.Duration(no*500) * time.Millisecond)
				fmt.Println(no)
			}(i)
		}
		wg.Wait()
		fmt.Println("Done")
	*/

	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(no int) {
			time.Sleep(time.Duration(no*500) * time.Millisecond)
			fmt.Println(no)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Println("Done")
}
