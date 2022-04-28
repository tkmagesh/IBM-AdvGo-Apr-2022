package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	mutex     *sync.Mutex
	closed    bool
}

var ErrInvalidPoolSize = errors.New("invalid pool size. pool size has to be > 0")
var ErrPoolClosed = errors.New("pool is closed")

func New(factory func() (io.Closer, error), poolSize int) (*Pool, error) {
	if poolSize <= 0 {
		return nil, ErrInvalidPoolSize
	}
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, poolSize),
		mutex:     &sync.Mutex{},
		closed:    false,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	/*
		When a resource is "Acquired()"?
			the pool will check if it has any resources in the resource-pool
			if yes, return the resource from the resource-pool
			else create a new instance using the factory and return it
	*/

	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquire : From Pool")
		return r, nil
	default:
		fmt.Println("Acquire : From Factory")
		return p.factory()
	}

}

func (p *Pool) Release(resource io.Closer) error {
	/*
		When a resource is "Released()"?
			check if the resource-pool is free to accomodate the released resource
			if yes, keep the resource in the pool
			else "discard" it (call the 'close()' method on the resource)
	*/
	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case p.resources <- resource:
		fmt.Println("Release : Into the Pool")
		return nil
	default:
		fmt.Println("Release : Close & discard the resource")
		return resource.Close()
	}
}

func (p *Pool) Close() {
	/*
		When the pool is 'Close()'?
			//prevent any more acquisition of the resources
			//make sure all the resources are 'closed' and discarded

	*/
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
