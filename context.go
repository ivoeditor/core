package core

import (
	termbox "github.com/nsf/termbox-go"
)

type Context interface {
	Logger() Logger
	Quit()
	Command(Command)
	Buffer() *Buffer
	Render()
}

type context struct {
	core    *Core
	buf     *Buffer
	expired bool
}

func (ctx *context) Logger() Logger {
	return ctx.core.log
}

func (ctx *context) Quit() {
	ctx.core.quit = true
	go termbox.Interrupt()
}

func (ctx *context) Command(cmd Command) {
	ctx.core.cmd = &cmd
	go termbox.Interrupt()
}

func (ctx *context) Buffer() *Buffer {
	if ctx.expired {
		return nil
	}
	return ctx.buf
}

func (ctx *context) Render() {
	if ctx.expired {
		return
	}

	for row := 0; row < ctx.buf.Rows; row++ {
		for col := 0; col < ctx.buf.Cols; col++ {
			c := ctx.buf.Get(col, row)
			fg := termbox.Attribute(c.Fore)
			if c.Attr&CellAttrBold != 0 {
				fg |= termbox.AttrBold
			}
			if c.Attr&CellAttrUnderline != 0 {
				fg |= termbox.AttrUnderline
			}
			bg := termbox.Attribute(c.Back)
			termbox.SetCell(col, row, c.Rune, fg, bg)
		}
	}

	if err := termbox.Flush(); err != nil {
		ctx.core.log.Errorf("core: failed to flush buffer: %v", err)
	}
}
