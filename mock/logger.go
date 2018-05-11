package mock

import "ivoeditor.com/core"

type logger struct{}

func NewLogger() core.Logger {
	return &logger{}
}

func (*logger) Info(v ...interface{})                  {}
func (*logger) Infof(format string, v ...interface{})  {}
func (*logger) Error(v ...interface{})                 {}
func (*logger) Errorf(format string, v ...interface{}) {}
