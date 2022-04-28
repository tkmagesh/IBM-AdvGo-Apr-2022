package pool

import "io"

type Pool struct {
	/*  */
}

func New(factory func() (io.Closer, error), poolSize int) (*Pool, error) {
	/*  */
}

func (p *Pool) Acquire() (io.Closer, error) {
	/*
		When a resource is "Acquired()"?
			the pool will check if it has any resources in the resource-pool
			if yes, return the resource from the resource-pool
			else create a new instance using the factory and return it
	*/
}

func (p *Pool) Release(resource io.Closer) error {
	/*
		When a resource is "Released()"?
			check if the resource-pool is free to accomodate the released resource
			if yes, keep the resource in the pool
			else "discard" it (call the 'close()' method on the resource)
	*/
}

func (p *Pool) Close() {
	/*
		When the pool is 'Close()'?
			//prevent any more acquisition of the resources
			//make sure all the resources are 'closed' and discarded

	*/
}
