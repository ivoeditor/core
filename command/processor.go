package command

import (
	"ivoeditor.com/core"
)

type Processor struct {
	ex core.Executor
	mp *Map
}

func NewProcessor(ex core.Executor) *Processor {
	p := Processor{
		ex: ex,
	}

	return &p
}

func (p *Processor) SetMap(mp *Map) {
	p.mp = mp
}

func (p *Processor) Process(ctx core.Context, cmd core.Command) {
	handler, ok := p.mp.get(cmd)
	if !ok {
		ctx.Logger().Infof("command: did not find mapping for %v", cmd)
		return
	}

	p.ex.Execute(executorFunc(handler, ctx, cmd))
}

func executorFunc(h Handler, ctx core.Context, cmd core.Command) func() {
	if h == nil {
		return func() {}
	}
	return func() {
		h(ctx, cmd)
	}
}
