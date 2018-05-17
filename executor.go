package core

type Executor interface {
	Close()
	Execute(func())
}

type executor struct{}

func newExecutor() Executor {
	return &executor{}
}

func (ex *executor) Close()           {}
func (ex *executor) Execute(f func()) { f() }
