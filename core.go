package core

import (
	termbox "github.com/nsf/termbox-go"
)

type Core struct {
	quit bool
	ex   Executor
	win  Window
	log  Logger
	ctx  *context
	cmd  *Command
}

func New() *Core {
	c := Core{
		ex:  newExecutor(),
		log: newLogger(),
	}

	c.ctx = &context{
		core: &c,
		buf:  newBuffer(0, 0),
	}

	return &c
}

func (c *Core) SetExecutor(ex Executor) {
	c.ex = ex
}

func (c *Core) SetWindow(win Window) {
	c.win = win
}

func (c *Core) SetLogger(log Logger) {
	c.log = log
}

func (c *Core) Run() {
	if err := termbox.Init(); err != nil {
		c.log.Errorf("core: could not initialize termbox: %v", err)
		return
	}
	defer termbox.Close()

	if c.win == nil {
		panic("core: window cannot be nil")
	}
	defer c.win.Close(c.newContext())

	defer c.ex.Close()

	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputNormal)

	for !c.quit {
		if c.cmd != nil {
			c.ex.Execute(func() { c.win.Command(c.newContext(), *c.cmd) })
			c.cmd = nil
		}

		data := make([]byte, 32)

		switch ev := termbox.PollRawEvent(data); ev.Type {
		case termbox.EventRaw:
			ev := termbox.ParseEvent(data[:ev.N])
			if ev.Type == termbox.EventNone {
				ev.Type = termbox.EventKey
				ev.Key = termbox.KeyEsc
			}

			switch ev.Type {
			case termbox.EventKey:
				c.ex.Execute(func() { c.win.Key(c.newContext(), newKey(ev)) })
			case termbox.EventMouse:
				c.ex.Execute(func() { c.win.Mouse(c.newContext(), newMouse(ev)) })
			}

		case termbox.EventResize:
			c.ex.Execute(func() { c.win.Resize(c.newContext()) })

		case termbox.EventInterrupt:
			break

		case termbox.EventError:
			c.log.Errorf("core: polled error termbox event: %v", ev.Err)

		default:
			c.log.Errorf("core: polled unknown termbox event")
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
