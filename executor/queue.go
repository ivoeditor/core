package executor

import (
	"sync"

	"github.com/eapache/channels"
)

const QueueInfiniteSize = 0

type Queue struct {
	funcsIn  chan<- interface{}
	funcsOut <-chan interface{}
	closing  bool
	close    sync.WaitGroup
}

func NewQueue() *Queue {
	return NewQueueSize(QueueInfiniteSize)
}

func NewQueueSize(size int) *Queue {
	if size < 0 {
		panic("queue: invalid size given")
	}

	q := Queue{}

	if size == QueueInfiniteSize {
		ch := channels.NewInfiniteChannel()
		q.funcsIn = ch.In()
		q.funcsOut = ch.Out()
	} else {
		ch := make(chan interface{}, size)
		q.funcsIn = ch
		q.funcsOut = ch
	}

	go q.execute()

	return &q
}

func (q *Queue) Close() {
	q.closing = true
	q.close.Wait()
}

func (q *Queue) Execute(f Func) {
	if q.closing {
		return
	}

	q.close.Add(1)
	q.funcsIn <- f
}

func (q *Queue) execute() {
	for {
		(<-q.funcsOut).(Func)()
		q.close.Done()
	}
}
