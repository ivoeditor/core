package command_test

import (
	"testing"
	"time"

	"ivoeditor.com/core"
	"ivoeditor.com/core/command"
	"ivoeditor.com/core/executor"
	"ivoeditor.com/core/mock"
)

func TestProcessor(t *testing.T) {
	tests := []struct {
		p    func(executor.Executor) (*command.Processor, *int)
		cmd  core.Command
		want int
	}{
		{
			p: func(ex executor.Executor) (*command.Processor, *int) {
				got := -1
				p := command.NewProcessor(ex)

				mp := command.NewMap()
				mp.Set("Hello", func(core.Context, core.Command) { got = 0 })
				mp.Set("World", func(core.Context, core.Command) { got = 1 })
				p.SetMap(mp)

				return p, &got
			},
			cmd: core.Command{
				Name: "Hello",
			},
			want: 0,
		},
		{
			p: func(ex executor.Executor) (*command.Processor, *int) {
				got := -1
				p := command.NewProcessor(ex)

				mp := command.NewMap()
				mp.SetFallback(func(core.Context, core.Command) { got = 0 })
				p.SetMap(mp)

				return p, &got
			},
			cmd:  core.Command{},
			want: 0,
		},
	}

	for i, test := range tests {
		ex := executor.NewQueue()
		p, got := test.p(ex)

		p.Process(mock.NewContext(), test.cmd)

		// Close might be called before handler is processed.
		time.Sleep(10 * time.Millisecond)
		ex.Close()

		if test.want != *got {
			t.Errorf("test %d: want %v, got %v", i, test.want, *got)
		}
	}
}
