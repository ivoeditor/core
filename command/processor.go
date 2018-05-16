package command

import (
	"ivoeditor.com/core"
	"ivoeditor.com/core/executor"
)

type Processor struct {
	ex    executor.Executor
	mp    *Map
	pairs chan *processorPair
}

type processorPair struct {
	ctx core.Context
	cmd core.Command
}

func NewProcessor(ex executor.Executor) *Processor {
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

func (p *Processor) Process(ctx core.Context, cmd core.Command) {
	p.pairs <- &processorPair{
		ctx: ctx,
		cmd: cmd,
	}
}

func (p *Processor) process() {
	pair := <-p.pairs

	handler, ok := p.mp.get(pair.cmd)
	if !ok {
		pair.ctx.Logger().Infof("command: did not find mapping for %v", pair.cmd)
		return
	}

	p.ex.Execute(executorFunc(handler, pair.ctx, pair.cmd))
}

func executorFunc(h Handler, ctx core.Context, cmd core.Command) executor.Func {
	if h == nil {
		return func() {}
	}
	return func() {
		h(ctx, cmd)
	}
}
