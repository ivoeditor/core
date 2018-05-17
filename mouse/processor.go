package mouse

import (
	"ivoeditor.com/core"
)

type Processor struct {
	ex    core.Executor
	mp    *Map
	pairs chan *processorPair
}

type processorPair struct {
	ctx core.Context
	mse core.Mouse
}

func NewProcessor(ex core.Executor) *Processor {
	p := Processor{
		ex:    ex,
		pairs: make(chan *processorPair),
	}

	go func() {
		for {
			p.process()
		}
	}()

	return &p
}

func (p *Processor) SetMap(mp *Map) {
	p.mp = mp
}

func (p *Processor) Process(ctx core.Context, mse core.Mouse) {
	p.pairs <- &processorPair{
		ctx: ctx,
		mse: mse,
	}
}

func (p *Processor) process() {
	pair := <-p.pairs

	handler, ok := p.mp.get(pair.mse)
	if !ok {
		pair.ctx.Logger().Infof("mouse: did not find mapping for %v", pair.mse)
		return
	}

	p.ex.Execute(executorFunc(handler, pair.ctx, pair.mse))
}

func executorFunc(h Handler, ctx core.Context, mse core.Mouse) func() {
	if h == nil {
		return func() {}
	}
	return func() {
		h(ctx, mse)
	}
}
