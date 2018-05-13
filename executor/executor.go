package executor

type Func func()

type Executor interface {
	Close()
	Execute(Func)
}
