package core

type Buffer struct {
	Cols int
	Rows int
	cc   []*Cell
}

func newBuffer(cols, rows int) *Buffer {
	return &Buffer{
		Cols: cols,
		Rows: rows,
		cc:   make([]*Cell, cols*rows),
	}
}

func (buf *Buffer) Set(col, row int, c *Cell) {
	buf.check(col, row)
	buf.cc[col+row*buf.Cols] = c
}

func (buf *Buffer) Get(col, row int) *Cell {
	buf.check(col, row)
	return buf.cc[col+row*buf.Cols]
}

func (buf *Buffer) check(col, row int) {
	if col >= buf.Cols || col < 0 || row >= buf.Rows || row < 0 {
		panic("runtime error: index out of bounds")
	}
}

type Cell struct {
	Rune rune
	Fore CellColor
	Back CellColor
	Attr CellAttr
}

type CellColor int

const (
	CellColorDefault CellColor = iota
	CellColorBlack
	CellColorRed
	CellColorGreen
	CellColorYellow
	CellColorBlue
	CellColorMagenta
	CellColorCyan
	CellColorWhite
)

func (cc CellColor) String() string {
	switch cc {
	case CellColorDefault:
		return "default"
	case CellColorBlack:
		return "black"
	case CellColorRed:
		return "red"
	case CellColorGreen:
		return "green"
	case CellColorYellow:
		return "yellow"
	case CellColorBlue:
		return "blue"
	case CellColorMagenta:
		return "magenta"
	case CellColorCyan:
		return "cyan"
	case CellColorWhite:
		return "white"
	}
	return "invalid"
}

type CellAttr int

const (
	CellAttrNone CellAttr = 1 << iota
	CellAttrBold
	CellAttrUnderline
)

func (ca CellAttr) String() string {
	switch ca {
	case CellAttrNone:
		return "none"
	case CellAttrBold:
		return "bold"
	case CellAttrUnderline:
		return "underline"
	}
	return "invalid"
}
