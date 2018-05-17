package core

import (
	"log"
	"os"
)

type Logger interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

type logger struct {
	info *log.Logger
	err  *log.Logger
}

func newLogger() *logger {
	return &logger{
		info: log.New(os.Stderr, "INFO  ", log.LstdFlags),
		err:  log.New(os.Stderr, "ERROR ", log.LstdFlags),
	}
}

func (l *logger) Info(v ...interface{})                  { l.info.Print(v) }
func (l *logger) Infof(format string, v ...interface{})  { l.info.Printf(format, v) }
func (l *logger) Error(v ...interface{})                 { l.err.Print(v) }
func (l *logger) Errorf(format string, v ...interface{}) { l.err.Printf(format, v) }
