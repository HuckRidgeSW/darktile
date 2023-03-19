package termutil

import (
	"image/color"
	"io"
	"os"
	"sync"
)

// Run2 is a stripped-down version of Run that leaves it up to the caller to
// start a process that interacts with the given pty.
//
// FIXME: Run2 is a shitty name.
func (t *Terminal) Run2(updateChan chan struct{}, pty *os.File, rows uint16, cols uint16) error {
	t.updateChan = updateChan
	t.pty = pty

	if err := t.SetSize(rows, cols); err != nil {
		return err
	}

	var processWg sync.WaitGroup
	processWg.Add(1)
	go func() {
		t.process()
		processWg.Done()
	}()

	t.running = true

	t.windowManipulator.SetTitle("darktile")

	_, _ = io.Copy(t, t.pty)
	close(t.closeChan)
	processWg.Wait()
	return nil
}

func (buffer *Buffer) Lines() []Line    { return buffer.lines }
func (line *Line) Cells() []Cell        { return line.cells }
func (line *Line) Append(cells ...Cell) { line.append(cells...) }
func (line *Line) Wrapped() bool        { return line.wrapped }
func (line *Line) SetWrapped(w bool)    { line.wrapped = w }

func (line *Line) SetCell(n int, cell Cell) {
	line.cells[n] = cell
}

func (line *Line) Truncate(n int) {
	if n < len(line.cells) {
		line.cells = line.cells[:n]
	}
}

func (line *Line) Copy() Line {
	l := Line{
		wrapped: line.wrapped,
		cells:   make([]Cell, len(line.cells)),
	}
	copy(l.cells, line.cells)
	return l
}

func (line *Line) StringNoNulls() string {
	runes := []rune{}
	for _, cell := range line.cells {
		r := cell.r.Rune
		if r == 0 {
			r = ' '
		}
		runes = append(runes, r)
	}
	l := len(runes)
	for l > 0 && runes[l-1] == ' ' {
		l--
	}
	runes = runes[:l]
	return string(runes)
}

func (buffer *Buffer) DefaultCell(applyEffects bool) Cell {
	return buffer.defaultCell(applyEffects)
}

func NewCell(r MeasuredRune, attr CellAttributes) Cell {
	return Cell{
		r:    r,
		attr: attr,
	}
}

func NewCellAttributes(
	fg, bg color.Color,
	bold, italic, dim, underline, strikethrough, blink, inverse, hidden bool,
) CellAttributes {
	return CellAttributes{
		fgColour:      fg,
		bgColour:      bg,
		bold:          bold,
		italic:        italic,
		dim:           dim,
		underline:     underline,
		strikethrough: strikethrough,
		blink:         blink,
		inverse:       inverse,
		hidden:        hidden,
	}
}

func (c Cell) Equal(c2 Cell) bool {
	return c.r == c2.r &&
		c.attr.Equal(c2.attr)
}

func (ca CellAttributes) Equal(cb CellAttributes) bool {
	return EqColors(ca.fgColour, cb.fgColour) &&
		EqColors(ca.bgColour, cb.bgColour) &&
		ca.bold == cb.bold &&
		ca.italic == cb.italic &&
		ca.dim == cb.dim &&
		ca.underline == cb.underline &&
		ca.strikethrough == cb.strikethrough &&
		ca.blink == cb.blink &&
		ca.inverse == cb.inverse &&
		ca.hidden == cb.hidden
}

func EqColors(c1, c2 color.Color) bool {
	// darktile uses color.RGBA, gio uses color.NRGBA.  If they're the same
	// type, just compare them directly.

	if c1 == nil || c2 == nil {
		return c1 == c2
	}

	// color.RGBA is the predominant (only?) kind of color used in Darktile.
	{
		c1, ok1 := c1.(color.RGBA)
		c2, ok2 := c2.(color.RGBA)
		if ok1 && ok2 {
			return c1 == c2
		}
	}

	// color.RGBA is the predominant (only?) kind of color used in Gio.
	{
		c1, ok1 := c1.(color.NRGBA)
		c2, ok2 := c2.(color.NRGBA)
		if ok1 && ok2 {
			return c1 == c2
		}
	}

	// compare them by component
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return r1 == r2 &&
		g1 == g2 &&
		b1 == b2 &&
		a1 == a2
}
