package command

import "ivoeditor.com/core"

type Map struct {
	handlers map[string]Handler
	fallback Handler
}

func NewMap() *Map {
	return &Map{
		handlers: make(map[string]Handler),
	}
}

func (mp *Map) Set(name string, h Handler) {
	mp.handlers[name] = h
}

func (mp *Map) SetFallback(h Handler) {
	mp.fallback = h
}

func (mp *Map) get(cmd core.Command) (Handler, bool) {
	handler, ok := mp.handlers[cmd.Name]
	if !ok {
		handler = mp.fallback
	}
	return handler, handler != nil
}
