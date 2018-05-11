package mock

import (
	"ivoeditor.com/core"
)

type context struct{}

func NewContext() core.Context {
	return &context{}
}

func (*context) Logger() core.Logger  { return NewLogger() }
func (*context) Quit()                {}
func (*context) Command(core.Command) {}
func (*context) Buffer() *core.Buffer { return nil }
func (*context) Render()              {}
