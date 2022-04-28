package main

import (
	"fmt"
	"log"
	"math/rand"
	"pool-demo/db"
	"pool-demo/pool"
	"sync"
	"time"
)

func main() {
	/*
		create an instance of a Pool (with pool size and the factory function)

		What is a Resource?
			Any object that implments the "Close()" method (io.Closer interface)

		When a resource is "Acquired()"?
			the pool will check if it has any resources in the resource-pool
			if yes, return the resource from the resource-pool
			else create a new instance using the factory and return it

		When a resource is "Released()"?
			check if the resource-pool is free to accomodate the released resource
			if yes, keep the resource in the pool
			else "discard" it (call the 'close()' method on the resource)

		When the pool is 'Close()'?
			//prevent any more acquisition of the resources
			//make sure all the resources are 'closed' and discarded

		Impartant Note:
			When a resource is "Acquired" by a client, the same source should NOT be given to another client until it is released

		APIs
			-New(factory, poolSize)
			-Acquire() => resource
			-Release(resource)
			-Close()
	*/

	p, err := pool.New(db.DBConnectionFactory, 5 /* pool size */)

	if err != nil {
		log.Fatalln(err)
	}
	wg := &sync.WaitGroup{}
	clientCount := 10
	wg.Add(clientCount)
	for client := 0; client < clientCount; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()
	// 5 resources should have been discarded and 5 resources should have been maintained in the pool
	fmt.Println("Second batch of operations")
	var input string
	fmt.Scanln(&input)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(client int) {
			doWork(client, p) //Resources should be returned from the pool when acquired
			wg.Done()
		}(i)
	}
	wg.Wait()
	p.Close()
}

func doWork(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Release(conn)
	fmt.Printf("Worker : %d, Acquired %d:\n", id, conn.(*db.DBConnection).ID)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond) //simulating a time consuming operation
	fmt.Printf("Worker Done : %d, Releasing %d:\n", id, conn.(*db.DBConnection).ID)
}
