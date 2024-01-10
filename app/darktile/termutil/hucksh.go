package termutil

import (
	"image/color"

	i_termutil "github.com/liamg/darktile/internal/app/darktile/termutil"
)

type (
	Buffer            = i_termutil.Buffer
	Terminal          = i_termutil.Terminal
	Option            = i_termutil.Option
	Theme             = i_termutil.Theme
	ThemeFactory      = i_termutil.ThemeFactory
	WindowManipulator = i_termutil.WindowManipulator
	WindowState       = i_termutil.WindowState
	Cell              = i_termutil.Cell
	CellSlice         = i_termutil.CellSlice
	CellAttributes    = i_termutil.CellAttributes
	Line              = i_termutil.Line
	MeasuredRune      = i_termutil.MeasuredRune
)

const (
	StateUnknown   WindowState = i_termutil.StateUnknown
	StateMinimised             = i_termutil.StateMinimised
	StateNormal                = i_termutil.StateNormal
	StateMaximised             = i_termutil.StateMaximised
)

func New(options ...Option) *Terminal                  { return i_termutil.New(options...) }
func WithTheme(theme *Theme) Option                    { return i_termutil.WithTheme(theme) }
func WithWindowManipulator(m WindowManipulator) Option { return i_termutil.WithWindowManipulator(m) }
func WithLogFile(path string) Option                   { return i_termutil.WithLogFile(path) }
func NewThemeFactory() *ThemeFactory                   { return i_termutil.NewThemeFactory() }

func NewCell(r MeasuredRune, attr CellAttributes) Cell { return i_termutil.NewCell(r, attr) }

func NewCellAttributes(
	fg, bg color.Color,
	bold, italic, dim, underline, strikethrough, blink, inverse, hidden bool,
) CellAttributes {
	return i_termutil.NewCellAttributes(
		fg, bg,
		bold, italic, dim, underline, strikethrough, blink, inverse, hidden)
}
