package mouse

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

func (p *Processor) Process(ctx core.Context, mse core.Mouse) {
	handler, ok := p.mp.get(mse)
	if !ok {
		ctx.Logger().Infof("mouse: did not find mapping for %v", mse)
		return
	}

	p.ex.Execute(executorFunc(handler, ctx, mse))
}

func executorFunc(h Handler, ctx core.Context, mse core.Mouse) func() {
	if h == nil {
		return func() {}
	}
	return func() {
		h(ctx, mse)
	}
}
