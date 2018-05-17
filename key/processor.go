package key

import (
	"math"
	"time"

	"ivoeditor.com/core"
)

type Processor struct {
	Timeout time.Duration

	ex core.Executor

	mode  string
	modes map[string]*Map

	pairs chan *processorPair
}

type processorPair struct {
	ctx core.Context
	key core.Key
}

func NewProcessor(ex core.Executor) *Processor {
	p := Processor{
		ex:    ex,
		modes: make(map[string]*Map),
		pairs: make(chan *processorPair),
	}

	go func() {
		for {
			p.process()
		}
	}()

	return &p
}

func (p *Processor) SetMode(mode string) {
	if _, ok := p.modes[mode]; !ok {
		panic("key: tried to set non-existent mode")
	}

	p.mode = mode
}

func (p *Processor) SetMap(mode string, mp *Map) {
	p.modes[mode] = mp
}

func (p *Processor) Process(ctx core.Context, key core.Key) {
	p.pairs <- &processorPair{
		ctx: ctx,
		key: key,
	}
}

func (p *Processor) process() {
	var (
		ctx     core.Context
		keys    []core.Key
		handler Handler
	)

	for {
		dur := p.Timeout
		if len(keys) == 0 {
			dur = time.Duration(math.MaxInt64)
		}

		select {
		case pair := <-p.pairs:
			ctx = pair.ctx
			keys = append(keys, pair.key)
		case <-time.After(dur):
			p.ex.Execute(executorFunc(handler, ctx, keys))
		}

		mp := p.modes[p.mode]

		var more, ok bool
		handler, more, ok = mp.get(keys)

		if !ok {
			ctx.Logger().Infof("key: did not find mapping for %v", keys)
			return
		}
		if more {
			continue
		}

		p.ex.Execute(executorFunc(handler, ctx, keys))
	}
}

func executorFunc(h Handler, ctx core.Context, keys []core.Key) func() {
	if h == nil {
		return func() {}
	}
	return func() {
		h(ctx, keys)
	}
}
