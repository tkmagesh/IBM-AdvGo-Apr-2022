package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	MAXPROCS, e := strconv.Atoi(os.Args[1])
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Println("existing GOMAXPROCS = ", runtime.GOMAXPROCS(MAXPROCS))
	noOfRoutines, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	for i := 1; i <= noOfRoutines; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	var input string
	fmt.Scanln(&input)
}

func fn(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn -", i)
	/* var input string
	fmt.Scanln(&input) */
}
