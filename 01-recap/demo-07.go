package main

import "fmt"

func main() {
	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	*/

	/*
		ch := make(chan int)
		no := <-ch
		ch <- 100
		fmt.Println("no =", no)
	*/
	/*
		ch := make(chan int)
		go func() {
			ch <- 100
		}()
		no := <-ch //initiating the 'receive' operation, but there is not data. So, the schedule will look for other scheduled goroutines and execute them
		fmt.Println("no =", no)
	*/

	ch := make(chan int, 1) // => even when a reveive operation is not initiated, the channel can hold '1' int value
	ch <- 100
	fmt.Println(len(ch))
	no := <-ch
	fmt.Println("no =", no)
}
