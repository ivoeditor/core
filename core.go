package core

import (
	termbox "github.com/nsf/termbox-go"
)

type Core struct {
	quit bool
	win  Window
	log  Logger
	ctx  *context
	cmd  *Command
}

func New(win Window, log Logger) *Core {
	if win == nil {
		panic("core: win is nil")
	}
	if log == nil {
		log = newLogger()
	}

	return &Core{
		log: log,
		win: win,
	}
}

func (c *Core) Run() {
	if err := termbox.Init(); err != nil {
		c.log.Errorf("termbox: could not initialize: %v", err)
		return
	}
	defer termbox.Close()

	defer c.win.Close(c.newContext())

	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputNormal)

	for !c.quit {
		if c.cmd != nil {
			c.win.Command(c.newContext(), *c.cmd)
			c.cmd = nil
		}

		data := make([]byte, 32)

		switch e := termbox.PollRawEvent(data); e.Type {
		case termbox.EventRaw:
			data := data[:e.N]
			e := termbox.ParseEvent(data)
			if e.Type == termbox.EventNone {
				e.Type = termbox.EventKey
				e.Key = termbox.KeyEsc
			}

			switch e.Type {
			case termbox.EventKey:
				c.win.Key(c.newContext(), newKey(e))
			case termbox.EventMouse:
				c.win.Mouse(c.newContext(), newMouse(e))
			}

		case termbox.EventResize:
			c.win.Resize(c.newContext())

		case termbox.EventInterrupt:
			break

		case termbox.EventError:
			c.log.Errorf("termbox: polled error event: %v", e.Err)

		default:
			c.log.Errorf("termbox: polled unknown event")
		}
	}
}

func (c *Core) newContext() *context {
	c.ctx.expired = true

	c.ctx = &context{
		core: c,
		buf:  c.ctx.buf,
	}

	cols, rows := termbox.Size()
	if c.ctx.buf.Cols != cols || c.ctx.buf.Rows != rows {
		buf := newBuffer(cols, rows)
		for row := 0; row < c.ctx.buf.Rows && row < buf.Rows; row++ {
			for col := 0; col < c.ctx.buf.Cols && col < buf.Cols; col++ {
				c := c.ctx.buf.Get(col, row)
				buf.Set(col, row, c)
			}
		}

		c.ctx.buf = buf
	}

	return c.ctx
}
