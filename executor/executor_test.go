package executor_test

import (
	"reflect"
	"testing"
	"time"

	"ivoeditor.com/core"
	"ivoeditor.com/core/executor"
)

func TestExecutor(t *testing.T) {
	tests := []struct {
		ex    core.Executor
		funcs []func()
		want  []int
	}{
		{
			ex: executor.NewQueue(),
			funcs: []func(){
				func() {},
				func() {},
				func() {},
				func() {},
				func() {},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			ex: executor.NewConcurrent(),
			funcs: []func(){
				func() { time.Sleep(4 * time.Millisecond) },
				func() { time.Sleep(2 * time.Millisecond) },
				func() { time.Sleep(8 * time.Millisecond) },
				func() { time.Sleep(10 * time.Millisecond) },
				func() { time.Sleep(6 * time.Millisecond) },
			},
			want: []int{1, 0, 4, 2, 3},
		},
	}

	for i, test := range tests {
		var got []int

		for i, f := range test.funcs {
			j, g := i, f
			test.ex.Execute(func() {
				g()
				got = append(got, j)
			})
		}

		test.ex.Close()

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("test %d: want %v, got %v", i, test.want, got)
		}
	}
}
