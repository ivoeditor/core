package mouse_test

import (
	"testing"
	"time"

	"ivoeditor.com/core"
	"ivoeditor.com/core/executor"
	"ivoeditor.com/core/mock"
	"ivoeditor.com/core/mouse"
)

func TestProcessor(t *testing.T) {
	tests := []struct {
		p    func(executor.Executor) (*mouse.Processor, *int)
		mse  core.Mouse
		want int
	}{
		{
			p: func(ex executor.Executor) (*mouse.Processor, *int) {
				got := -1
				p := mouse.NewProcessor(ex)

				mp := mouse.NewMap()
				mp.SetRegion(core.Mouse{
					Action: core.MouseButtonLeft,
					Col:    5,
					Row:    6,
				}, 12, 10, func(core.Context, core.Mouse) { got = 0 })
				mp.SetRegion(core.Mouse{
					Action: core.MouseButtonRight,
					Col:    5,
					Row:    6,
				}, 12, 10, func(core.Context, core.Mouse) { got = 1 })
				mp.SetRegion(core.Mouse{
					Action: core.MouseButtonLeft,
					Col:    15,
					Row:    26,
				}, 40, 100, func(core.Context, core.Mouse) { got = 2 })
				p.SetMap(mp)

				return p, &got
			},
			mse: core.Mouse{
				Action: core.MouseButtonLeft,
				Col:    7,
				Row:    7,
			},
			want: 0,
		},
		{
			p: func(ex executor.Executor) (*mouse.Processor, *int) {
				got := -1
				p := mouse.NewProcessor(ex)

				mp := mouse.NewMap()
				mp.SetFallback(func(core.Context, core.Mouse) { got = 0 })
				p.SetMap(mp)

				return p, &got
			},
			mse:  core.Mouse{},
			want: 0,
		},
		{
			p: func(ex executor.Executor) (*mouse.Processor, *int) {
				got := -1
				p := mouse.NewProcessor(ex)

				mp := mouse.NewMap()
				mp.Set(core.Mouse{
					Action: core.MouseWheelUp,
				}, func(core.Context, core.Mouse) { got = 0 })
				mp.SetFallback(func(core.Context, core.Mouse) { got = 1 })
				p.SetMap(mp)

				return p, &got
			},
			mse:  core.Mouse{Action: core.MouseWheelUp},
			want: 0,
		},
	}

	for i, test := range tests {
		ex := executor.NewQueue()
		p, got := test.p(ex)

		p.Process(mock.NewContext(), test.mse)

		// Close might be called before handler is processed.
		time.Sleep(10 * time.Millisecond)
		ex.Close()

		if test.want != *got {
			t.Errorf("test %d: want %v, got %v", i, test.want, *got)
		}
	}
}
