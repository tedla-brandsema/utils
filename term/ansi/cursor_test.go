package ansi

import (
	"bytes"
	"os"
	"testing"

	"golang.org/x/term"
)

func TestWrite_NewlineReplacement(t *testing.T) {
	buf := &bytes.Buffer{}
	w := &Writer{w: buf, enabled: true}

	input := []byte("line1\nline2\nend")
	n, err := w.Write(input)
	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}
	if n != len(input) {
		t.Errorf("Write() returned n=%d; want %d", n, len(input))
	}

	got := buf.String()
	want := "line1\n\x1b[Gline2\n\x1b[Gend"
	if got != want {
		t.Errorf("Write() output = %q; want %q", got, want)
	}
}

func TestWrite_Disabled(t *testing.T) {
	buf := &bytes.Buffer{}
	w := &Writer{w: buf, enabled: false}

	input := []byte("abc\ndef")
	n, err := w.Write(input)
	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}
	if n != len(input) {
		t.Errorf("Write() returned n=%d; want %d", n, len(input))
	}

	got := buf.String()
	want := "abc\ndef"
	if got != want {
		t.Errorf("Write() output = %q; want %q", got, want)
	}
}

func TestNewline(t *testing.T) {
	buf := &bytes.Buffer{}
	w := &Writer{w: buf, enabled: true}

	w.Newline()
	got := buf.String()
	want := "\n\x1b[G"
	if got != want {
		t.Errorf("Newline() = %q; want %q", got, want)
	}
}

func TestResetLine(t *testing.T) {
	buf := &bytes.Buffer{}
	w := &Writer{w: buf, enabled: true}

	w.ResetLine()
	got := buf.String()
	want := "\x1b[1A\x1b[2K" // MoveUp + ClearLine
	if got != want {
		t.Errorf("ResetLine() = %q; want %q", got, want)
	}
}

func TestCursorHelpers(t *testing.T) {
	buf := &bytes.Buffer{}
	w := &Writer{w: buf, enabled: true}

	w.MoveUpRows(2)
	w.MoveDownRows(3)
	w.MoveLeftCols(4)
	w.MoveRightCols(5)

	w.MoveUp()
	w.MoveDown()
	w.MoveLeft()
	w.MoveRight()

	w.CursorHome()
	w.CursorPos(7, 8)
	w.ClearLine()
	w.ClearLineRight()
	w.ClearScreen()
	w.ClearScreenBelow()
	w.HideCursor()
	w.ShowCursor()
	w.SaveCursor()
	w.RestoreCursor()

	got := buf.String()
	want := "\x1b[2A\x1b[3B\x1b[4D\x1b[5C" +
		"\x1b[1A\x1b[1B\x1b[1D\x1b[1C" +
		"\x1b[G\x1b[7;8H\x1b[2K\x1b[K\x1b[2J\x1b[J\x1b[?25l\x1b[?25h\x1b[s\x1b[u"

	if got != want {
		t.Errorf("Cursor helpers output = %q; want %q", got, want)
	}
}

func TestOverwritePreviousLine(t *testing.T) {
	buf := &bytes.Buffer{}
	w := &Writer{w: buf, enabled: true}

	// Write first line
	w.Write([]byte("Processing..."))
	w.Newline() // move to next line (cursor at col 0)

	// Overwrite the previous line
	w.ResetLine()
	w.Write([]byte("Done!       ")) // add spaces to fully clear old text
	w.Newline()

	got := buf.String()
	want := "Processing...\n\x1b[G" +      // first line + automatic newline reset
		"\x1b[1A\x1b[2KDone!       \n\x1b[G" // ResetLine + new text + newline reset

	if got != want {
		t.Errorf("Overwrite previous line output = %q; want %q", got, want)
	}
}

func TestNewWriter_WithNonTTY(t *testing.T) {
	buf := &bytes.Buffer{}
	w, err := NewWriter(buf)
	if err != nil {
		t.Fatalf("NewWriter() returned unexpected error: %v", err)
	}

	if w == nil {
		t.Fatal("NewWriter() returned nil Writer")
	}
	if w.enabled {
		t.Error("NewWriter() enabled = true; want false for non-TTY writer")
	}
	if w.rawFD != -1 {
		t.Errorf("NewWriter() rawFD = %d; want -1", w.rawFD)
	}
	if w.oldState != nil {
		t.Error("NewWriter() oldState != nil; want nil")
	}
}

func TestNewWriter_WithStdout_AndFakeTTY(t *testing.T) {
	// We can’t force term.IsTerminal to return true without a TTY,
	// but we can simulate the setup and ensure no panic or side effects.
	w, err := NewWriter(os.Stdout)
	if err != nil {
		// On CI or non-TTY environments, this may fail if MakeRaw is called.
		// That’s acceptable, so we only assert sanity.
		t.Logf("NewWriter(os.Stdout) returned expected error on non-TTY: %v", err)
		return
	}

	// If running interactively in a terminal, ensure proper setup.
	if w.enabled && w.rawFD < 0 {
		t.Error("expected valid rawFD when enabled == true")
	}
	if !w.enabled && w.oldState != nil {
		t.Error("oldState should be nil when not enabled")
	}
	_ = w.Close()
}

func TestNewOutAndErrWriter(t *testing.T) {
	out, err := NewOutWriter()
	if err != nil && term.IsTerminal(int(os.Stdout.Fd())) {
		t.Errorf("NewOutWriter() unexpected error on TTY: %v", err)
	}
	if out == nil {
		t.Fatal("NewOutWriter() returned nil")
	}

	errw, err := NewErrWriter()
	if err != nil && term.IsTerminal(int(os.Stderr.Fd())) {
		t.Errorf("NewErrWriter() unexpected error on TTY: %v", err)
	}
	if errw == nil {
		t.Fatal("NewErrWriter() returned nil")
	}

	// Both should have valid io.Writer references
	if out.w == nil || errw.w == nil {
		t.Error("expected Writer.w to be set for both NewOutWriter and NewErrWriter")
	}

	// Close should always be safe
	if cerr := out.Close(); cerr != nil {
		t.Errorf("Close() returned unexpected error: %v", cerr)
	}
	if cerr := errw.Close(); cerr != nil {
		t.Errorf("Close() returned unexpected error: %v", cerr)
	}
}