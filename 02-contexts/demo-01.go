package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//root context
	rootCtx := context.Background()
	dataCtx := context.WithValue(rootCtx, "root", "main")
	fnCtx, cancel := context.WithCancel(dataCtx)
	defer cancel()
	go fn(fnCtx)

	var input string
	fmt.Scanln(&input)
	cancel()
	fmt.Scanln(&input)
	//child contexts
	/*
		context.WithCancel()
			=> returns a cancel function using which a "cancel" signal can be sent
		context.WithValue()

		context.WithDeadline()
		context.WithTimeout()
	*/
}

func fn(ctx context.Context) {
	fmt.Println("[@fn] data from context -> ", ctx.Value("root"))
	dataCtx := context.WithValue(ctx, "user", "magesh")
	f1Ctx, f1Cancel := context.WithCancel(dataCtx)
	defer f1Cancel()
	go f1(f1Ctx)
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fn done")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Print(".")
		}
	}
}

func f1(ctx context.Context) {
	fmt.Println("[@f1] data from context['root'] -> ", ctx.Value("root"))
	fmt.Println("[@f1] data from context['user'] -> ", ctx.Value("user"))
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f1 done")
			break LOOP
		default:
			time.Sleep(2000 * time.Millisecond)
			fmt.Print("f1")
		}
	}
}
