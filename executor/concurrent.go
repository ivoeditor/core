package executor

import (
	"sync"
)

type Concurrent struct {
	closing bool
	close   sync.WaitGroup
}

func NewConcurrent() *Concurrent {
	return &Concurrent{}
}

func (c *Concurrent) Close() {
	c.closing = true
	c.close.Wait()
}

func (c *Concurrent) Execute(f func()) {
	if c.closing {
		return
	}

	c.close.Add(1)
	go func() {
		f()
		c.close.Done()
	}()
}
