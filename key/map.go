package key

import "ivoeditor.com/core"

type Map struct {
	*mapNode
	fallback Handler
}

type mapNode struct {
	children map[core.Key]*mapNode
	handler  Handler
}

func NewMap() *Map {
	return &Map{
		mapNode: newMapNode(),
	}
}

func newMapNode() *mapNode {
	return &mapNode{
		children: make(map[core.Key]*mapNode),
	}
}

func (mp *Map) Set(keys []core.Key, h Handler) {
	node := mp.mapNode
	for _, key := range keys {
		child, ok := node.children[key]
		if !ok {
			child = newMapNode()
			node.children[key] = child
		}

		node = child
	}

	node.handler = h
}

func (mp *Map) SetFallback(h Handler) {
	mp.fallback = h
}

func (mp *Map) get(keys []core.Key) (Handler, bool, bool) {
	node := mp.mapNode
	for _, key := range keys {
		var ok bool
		node, ok = node.children[key]
		if !ok {
			if mp.fallback != nil {
				return mp.fallback, false, true
			}
			return nil, false, false
		}
	}

	return node.handler, len(node.children) > 0, true
}
