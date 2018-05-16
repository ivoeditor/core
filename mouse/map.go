package mouse

import "ivoeditor.com/core"

type Map struct {
	items    []*mapItem
	regions  []*mapRegion
	fallback Handler
}

type mapItem struct {
	mse     core.Mouse
	handler Handler
}

type mapRegion struct {
	*mapItem
	cols int
	rows int
}

func NewMap() *Map {
	return &Map{}
}

func (mp *Map) Set(mse core.Mouse, h Handler) {
	mp.items = append(mp.items, &mapItem{
		mse:     mse,
		handler: h,
	})
}

func (mp *Map) SetRegion(mse core.Mouse, cols, rows int, h Handler) {
	mp.regions = append(mp.regions, &mapRegion{
		mapItem: &mapItem{
			mse:     mse,
			handler: h,
		},
		cols: cols,
		rows: rows,
	})
}

func (mp *Map) SetFallback(h Handler) {
	mp.fallback = h
}

func (mp *Map) get(mse core.Mouse) (Handler, bool) {
	for _, item := range mp.items {
		if item.mse == mse {
			return item.handler, true
		}
	}
	for _, reg := range mp.regions {
		if mse.Col >= reg.mse.Col &&
			mse.Col < reg.mse.Col+reg.cols &&
			mse.Row >= reg.mse.Row &&
			mse.Row < reg.mse.Row+reg.rows {
			return reg.handler, true
		}
	}
	return mp.fallback, mp.fallback != nil
}
