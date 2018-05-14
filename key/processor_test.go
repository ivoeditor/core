package key_test

import (
	"testing"
	"time"

	"ivoeditor.com/core"
	"ivoeditor.com/core/executor"
	"ivoeditor.com/core/key"
	"ivoeditor.com/core/mock"
)

func TestProcessor(t *testing.T) {
	type event struct {
		key core.Key
		dur time.Duration
	}

	tests := []struct {
		p    func(executor.Executor) (*key.Processor, *int)
		mode string
		evs  []event
		want int
	}{
		{
			p: func(ex executor.Executor) (*key.Processor, *int) {
				got := -1
				p := key.NewProcessor(ex)
				p.Timeout = 5 * time.Millisecond

				mp := key.NewMap()
				mp.Set([]core.Key{
					{Rune: 'a'},
					{Rune: 'b'},
					{Rune: 'c'},
				}, func(core.Context, []core.Key) { got = 0 })
				mp.Set([]core.Key{
					{Rune: 'a'},
					{Rune: 'b'},
					{Rune: 'c'},
					{Rune: 'd'},
				}, func(core.Context, []core.Key) { got = 1 })
				p.SetMap("", mp)

				return p, &got
			},
			mode: "",
			evs: []event{
				{key: core.Key{Rune: 'a'}, dur: 0 * time.Millisecond},
				{key: core.Key{Rune: 'b'}, dur: 0 * time.Millisecond},
				{key: core.Key{Rune: 'c'}, dur: 10 * time.Millisecond},
			},
			want: 0,
		},
		{
			p: func(ex executor.Executor) (*key.Processor, *int) {
				got := -1
				p := key.NewProcessor(executor.NewQueue())

				mp := key.NewMap()
				mp.SetFallback(func(core.Context, []core.Key) { got = 1 })
				p.SetMap("", mp)

				return p, &got
			},
			mode: "",
			evs: []event{
				{key: core.Key{Rune: 'a'}, dur: 0 * time.Millisecond},
			},
			want: 1,
		},
		{
			p: func(ex executor.Executor) (*key.Processor, *int) {
				got := -1
				p := key.NewProcessor(executor.NewQueue())

				mp := key.NewMap()
				mp.Set([]core.Key{
					{Rune: 'a'},
					{Rune: 'b'},
					{Rune: 'c'},
				}, func(core.Context, []core.Key) { got = 0 })
				mp.Set([]core.Key{
					{Rune: 'a'},
					{Rune: 'b'},
					{Rune: 'c'},
					{Rune: 'd'},
				}, func(core.Context, []core.Key) { got = 1 })
				p.SetMap("normal", mp)

				mp = key.NewMap()
				mp.Set([]core.Key{
					{Rune: 'x'},
					{Rune: 'y'},
					{Rune: 'z'},
				}, func(core.Context, []core.Key) { got = 0 })
				mp.Set([]core.Key{
					{Rune: 'x'},
					{Rune: 'y'},
				}, func(core.Context, []core.Key) { got = 1 })
				p.SetMap("insert", mp)

				return p, &got
			},
			mode: "insert",
			evs: []event{
				{key: core.Key{Rune: 'a'}, dur: 0 * time.Millisecond},
			},
			want: -1,
		},
	}

	for i, test := range tests {
		ex := executor.NewQueue()
		p, got := test.p(ex)

		p.SetMode(test.mode)
		for _, ev := range test.evs {
			p.Process(mock.NewContext(), ev.key)
			time.Sleep(ev.dur)
		}

		// Close might be called before handler is executed.
		time.Sleep(10 * time.Millisecond)
		ex.Close()

		if test.want != *got {
			t.Errorf("test %d: want %v, got %v", i, test.want, *got)
		}
	}
}
