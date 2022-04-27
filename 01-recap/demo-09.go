package main

import (
	"fmt"
	"time"
)

/*
func main() {
	fmt.Println("main - started")
	done := make(chan struct{})
	go doSomething(done)
	<-done
	fmt.Println("main - completed")
}

func doSomething(done chan struct{}) {
	fmt.Println("doing something - started")
	time.Sleep(5 * time.Second)
	fmt.Println("doing something - completed")
	done <- struct{}{}
}
*/

func main() {
	fmt.Println("main - started")
	done := doSomething()
	<-done
	fmt.Println("main - completed")
}

func doSomething() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("doing something - started")
		time.Sleep(5 * time.Second)
		fmt.Println("doing something - completed")
		done <- struct{}{}
	}()
	return done
}
