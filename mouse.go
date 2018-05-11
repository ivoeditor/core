package core

import termbox "github.com/nsf/termbox-go"

type Mouse struct {
	Action MouseAction
	Moved  bool
	Col    int
	Row    int
}

func newMouse(e termbox.Event) Mouse {
	var m Mouse

	m.Moved = e.Mod&termbox.ModMotion != 0
	m.Col = e.MouseX
	m.Row = e.MouseY

	switch e.Key {
	case termbox.MouseLeft:
		m.Action = MouseButtonLeft
	case termbox.MouseMiddle:
		m.Action = MouseButtonMiddle
	case termbox.MouseRight:
		m.Action = MouseButtonRight
	case termbox.MouseRelease:
		m.Action = MouseButtonRelease
	case termbox.MouseWheelUp:
		m.Action = MouseWheelUp
	case termbox.MouseWheelDown:
		m.Action = MouseWheelDown
	}

	return m
}

type MouseAction int

const (
	MouseButtonLeft MouseAction = iota
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonRelease
	MouseWheelUp
	MouseWheelDown
)

func (ma MouseAction) String() string {
	switch ma {
	case MouseButtonLeft:
		return "left"
	case MouseButtonMiddle:
		return "middle"
	case MouseButtonRight:
		return "right"
	case MouseButtonRelease:
		return "release"
	case MouseWheelUp:
		return "wheelUp"
	case MouseWheelDown:
		return "wheelDown"
	}
	return "invalid"
}
