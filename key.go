package core

import (
	"strings"

	termbox "github.com/nsf/termbox-go"
)

type Key struct {
	Code KeyCode
	Rune rune
	Mod  KeyMod
}

func newKey(ev termbox.Event) Key {
	var key Key

	if ev.Ch == 0 {
		switch ev.Key {
		case termbox.KeyF1:
			key.Code = KeyCodeF1
		case termbox.KeyF2:
			key.Code = KeyCodeF2
		case termbox.KeyF3:
			key.Code = KeyCodeF3
		case termbox.KeyF4:
			key.Code = KeyCodeF4
		case termbox.KeyF5:
			key.Code = KeyCodeF5
		case termbox.KeyF6:
			key.Code = KeyCodeF6
		case termbox.KeyF7:
			key.Code = KeyCodeF7
		case termbox.KeyF8:
			key.Code = KeyCodeF8
		case termbox.KeyF9:
			key.Code = KeyCodeF9
		case termbox.KeyF10:
			key.Code = KeyCodeF10
		case termbox.KeyF11:
			key.Code = KeyCodeF11
		case termbox.KeyF12:
			key.Code = KeyCodeF12
		case termbox.KeyInsert:
			key.Code = KeyCodeInsert
		case termbox.KeyDelete:
			key.Code = KeyCodeDelete
		case termbox.KeyHome:
			key.Code = KeyCodeHome
		case termbox.KeyEnd:
			key.Code = KeyCodeEnd
		case termbox.KeyPgup:
			key.Code = KeyCodePgup
		case termbox.KeyPgdn:
			key.Code = KeyCodePgdn
		case termbox.KeyArrowUp:
			key.Code = KeyCodeArrowUp
		case termbox.KeyArrowDown:
			key.Code = KeyCodeArrowDown
		case termbox.KeyArrowLeft:
			key.Code = KeyCodeArrowLeft
		case termbox.KeyArrowRight:
			key.Code = KeyCodeArrowRight
		case termbox.KeyCtrlSpace:
			key.Code = KeyCodeSpace
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlA:
			key.Rune = 'a'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlB:
			key.Rune = 'b'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlC:
			key.Rune = 'c'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlD:
			key.Rune = 'd'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlE:
			key.Rune = 'e'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlF:
			key.Rune = 'f'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlG:
			key.Rune = 'g'
			key.Mod = KeyModCtrl
		case termbox.KeyBackspace:
			key.Code = KeyCodeBackspace
		case termbox.KeyTab:
			key.Code = KeyCodeTab
		case termbox.KeyCtrlJ:
			key.Rune = 'j'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlK:
			key.Rune = 'k'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlL:
			key.Rune = 'l'
			key.Mod = KeyModCtrl
		case termbox.KeyEnter:
			key.Code = KeyCodeEnter
		case termbox.KeyCtrlN:
			key.Rune = 'n'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlO:
			key.Rune = 'o'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlP:
			key.Rune = 'p'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlQ:
			key.Rune = 'q'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlR:
			key.Rune = 'r'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlS:
			key.Rune = 's'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlT:
			key.Rune = 't'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlU:
			key.Rune = 'u'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlV:
			key.Rune = 'v'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlW:
			key.Rune = 'w'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlX:
			key.Rune = 'x'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlY:
			key.Rune = 'y'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrlZ:
			key.Rune = 'z'
			key.Mod = KeyModCtrl
		case termbox.KeyEsc:
			key.Code = KeyCodeEsc
		case termbox.KeyCtrl4:
			key.Rune = '4'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrl5:
			key.Rune = '5'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrl6:
			key.Rune = '6'
			key.Mod = KeyModCtrl
		case termbox.KeyCtrl7:
			key.Rune = '7'
			key.Mod = KeyModCtrl
		case termbox.KeySpace:
			key.Code = KeyCodeSpace
		case termbox.KeyBackspace2:
			key.Code = KeyCodeBackspace
		}
	} else {
		key.Rune = ev.Ch
	}

	if ev.Mod&termbox.ModAlt != 0 {
		key.Mod |= KeyModAlt
	}

	return key
}

type KeyCode int

const (
	KeyCodeRune KeyCode = iota
	KeyCodeF1
	KeyCodeF2
	KeyCodeF3
	KeyCodeF4
	KeyCodeF5
	KeyCodeF6
	KeyCodeF7
	KeyCodeF8
	KeyCodeF9
	KeyCodeF10
	KeyCodeF11
	KeyCodeF12
	KeyCodeInsert
	KeyCodeDelete
	KeyCodeHome
	KeyCodeEnd
	KeyCodePgup
	KeyCodePgdn
	KeyCodeArrowUp
	KeyCodeArrowDown
	KeyCodeArrowLeft
	KeyCodeArrowRight
	KeyCodeEsc
	KeyCodeEnter
	KeyCodeBackspace
	KeyCodeTab
	KeyCodeSpace
)

func (kc KeyCode) String() string {
	switch kc {
	case KeyCodeRune:
		return "rune"
	case KeyCodeF1:
		return "f1"
	case KeyCodeF2:
		return "f2"
	case KeyCodeF3:
		return "f3"
	case KeyCodeF4:
		return "f4"
	case KeyCodeF5:
		return "f5"
	case KeyCodeF6:
		return "f6"
	case KeyCodeF7:
		return "f7"
	case KeyCodeF8:
		return "f8"
	case KeyCodeF9:
		return "f9"
	case KeyCodeF10:
		return "f10"
	case KeyCodeF11:
		return "f11"
	case KeyCodeF12:
		return "f12"
	case KeyCodeInsert:
		return "insert"
	case KeyCodeDelete:
		return "delete"
	case KeyCodeHome:
		return "home"
	case KeyCodeEnd:
		return "end"
	case KeyCodePgup:
		return "pgup"
	case KeyCodePgdn:
		return "pgdn"
	case KeyCodeArrowUp:
		return "arrowUp"
	case KeyCodeArrowDown:
		return "arrowDown"
	case KeyCodeArrowLeft:
		return "arrowLeft"
	case KeyCodeArrowRight:
		return "arrowRight"
	case KeyCodeEsc:
		return "esc"
	case KeyCodeEnter:
		return "enter"
	case KeyCodeBackspace:
		return "backspace"
	case KeyCodeTab:
		return "tab"
	case KeyCodeSpace:
		return "space"
	}
	return "invalid"
}

type KeyMod int

const (
	KeyModNone KeyMod = 0
	KeyModCtrl KeyMod = 1 << (iota - 1)
	KeyModAlt
)

func (km KeyMod) String() string {
	if km == KeyModNone {
		return "none"
	}

	var mods []string
	if km&KeyModCtrl != 0 {
		mods = append(mods, "ctrl")
	}
	if km&KeyModAlt != 0 {
		mods = append(mods, "alt")
	}

	return strings.Join(mods, ", ")
}
