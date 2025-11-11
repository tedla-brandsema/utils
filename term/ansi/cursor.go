package ansi

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

// Writer wraps an io.Writer and optionally handles:
// - converting '\n' → '\n\x1b[G'
// - raw mode enabling/restoring
type Writer struct {
	w        io.Writer
	enabled  bool
	rawFD    int
	oldState *term.State
}

func NewOutWriter() (*Writer, error) { return NewWriter(os.Stdout) }
func NewErrWriter() (*Writer, error) { return NewWriter(os.Stderr) }

// NewWriter wraps the given io.Writer.
// It enables newline reset and raw mode if writing to a TTY.
// Returns a WriteCloser for deferred cleanup.
func NewWriter(w io.Writer) (*Writer, error) {
	fd := -1
	if f, ok := w.(*os.File); ok {
		fd = int(f.Fd())
	}

	writer := &Writer{
		w:       w,
		enabled: fd >= 0 && term.IsTerminal(fd),
		rawFD:   fd,
	}

	if writer.enabled {
		oldState, err := term.MakeRaw(fd)
		if err != nil {
			return nil, err
		}
		writer.oldState = oldState
	}

	return writer, nil
}

// Close restores the terminal if raw mode was enabled.
func (w *Writer) Close() error {
	if w.oldState != nil && w.rawFD >= 0 {
		return term.Restore(w.rawFD, w.oldState)
	}
	return nil
}

// Write implements io.Writer, replacing '\n' with '\n\x1b[G' in TTY mode.
func (w *Writer) Write(p []byte) (int, error) {
	if !w.enabled {
		return w.w.Write(p)
	}

	out := make([]byte, 0, len(p)+len("\x1b[G")*bytes.Count(p, []byte{'\n'}))
	start := 0
	for i, b := range p {
		if b == '\n' {
			out = append(out, p[start:i]...)
			out = append(out, '\n', '\x1b', '[', 'G')
			start = i + 1
		}
	}
	out = append(out, p[start:]...)

	_, err := w.w.Write(out)
	return len(p), err // report input bytes consumed
}

// --- Cursor / line / screen helpers ---
// All helpers use WriterInstance; call SetWriter first.

func (w *Writer) MoveUpRows(rows int) {
	if rows > 0 {
		fmt.Fprintf(w.w, "\x1b[%dA", rows)
	}
}
func (w *Writer) MoveDownRows(rows int) {
	if rows > 0 {
		fmt.Fprintf(w.w, "\x1b[%dB", rows)
	}
}
func (w *Writer) MoveRightCols(cols int) {
	if cols > 0 {
		fmt.Fprintf(w.w, "\x1b[%dC", cols)
	}
}
func (w *Writer) MoveLeftCols(cols int) {
	if cols > 0 {
		fmt.Fprintf(w.w, "\x1b[%dD", cols)
	}
}

func (w *Writer) MoveUp()    { w.MoveUpRows(1) }
func (w *Writer) MoveDown()  { w.MoveDownRows(1) }
func (w *Writer) MoveRight() { w.MoveRightCols(1) }
func (w *Writer) MoveLeft()  { w.MoveLeftCols(1) }

func (w *Writer) CursorHome()            { fmt.Fprint(w.w, "\x1b[G") }
func (w *Writer) CursorPos(row, col int) { fmt.Fprintf(w.w, "\x1b[%d;%dH", row, col) }

func (w *Writer) ResetLine()        { fmt.Fprint(w.w, "\r\x1b[2K") }
func (w *Writer) ClearLine()        { fmt.Fprint(w.w, "\x1b[2K") }
func (w *Writer) ClearLineRight()   { fmt.Fprint(w.w, "\x1b[K") }
func (w *Writer) ClearScreen()      { fmt.Fprint(w.w, "\x1b[2J") }
func (w *Writer) ClearScreenBelow() { fmt.Fprint(w.w, "\x1b[J") }

func (w *Writer) HideCursor()    { fmt.Fprint(w.w, "\x1b[?25l") }
func (w *Writer) ShowCursor()    { fmt.Fprint(w.w, "\x1b[?25h") }
func (w *Writer) SaveCursor()    { fmt.Fprint(w.w, "\x1b[s") }
func (w *Writer) RestoreCursor() { fmt.Fprint(w.w, "\x1b[u") }

// Newline prints a newline and moves the cursor to column 0.
func (w *Writer) Newline() { fmt.Fprint(w.w, "\n\x1b[G") }

// Works reliably in xterm, iTerm2, and most Linux TTYs.
// Does not work in Windows Terminal or cmd.exe 
// It’s safe to call even if unsupported; it’ll just be ignored.
func (w *Writer) EnableBlinkingCursor()  { fmt.Fprint(w.w, "\x1b[?12h") }
func (w *Writer) DisableBlinkingCursor() { fmt.Fprint(w.w, "\x1b[?12l") }