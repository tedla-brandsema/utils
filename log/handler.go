package log

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

type colorWriter struct {
	out    io.Writer
	prefix string
}

func (cw *colorWriter) Write(p []byte) (int, error) {
	n1, err := io.WriteString(cw.out, cw.prefix)
	if err != nil {
		return n1, err
	}
	n2, err := cw.out.Write(p)
	return n1 + n2, err
}

// DevHandler formats log records with color-coded levels for cli output,
// using the provided io.Writer.
type DevHandler struct {
	base slog.Handler
	opts *slog.HandlerOptions
	cw   *colorWriter
}

func (h *DevHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.base.Enabled(ctx, level)
}

func (h *DevHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &DevHandler{
		base: h.base.WithAttrs(attrs),
		opts: h.opts,
		cw:   h.cw,
	}
}

func (h *DevHandler) WithGroup(name string) slog.Handler {
	return &DevHandler{
		base: h.base.WithGroup(name),
		opts: h.opts,
		cw:   h.cw,
	}
}

const format = "[15:05:05]"

// Handle formats and prints a log record at the appropriate log level.
func (h *DevHandler) Handle(_ context.Context, r slog.Record) error {
	level := LevelString(r.Level)

	timeStr := r.Time.Format(format)
	msg := Cyan(r.Message)

	h.cw.prefix = fmt.Sprintf("%s %s %s\n", timeStr, level, msg)

	padding := strings.Repeat(" ", len(format)+7)
	keyFormat := padding + "%s: "

	buf, tw := getWriterPair()
	defer putWriterPair(buf, tw)

	r.Attrs(func(a slog.Attr) bool {
		_, _ = fmt.Fprintf(tw, keyFormat+"\t%v\n", a.Key, a.Value)
		return true
	})

	if h.opts.AddSource {
		path := r.Source().File
		path = shortenPath(path)
		_, _ = fmt.Fprintf(tw, keyFormat+"\t%s:%d\n", "source", path, r.Source().Line)
	}
	_ = tw.Flush()

	if _, err := h.cw.Write(buf.Bytes()); err != nil {
		slog.Error("failed to write response", "err", err)
		return err
	}
	return nil
}

var (
	bufPool = Instance(func() *bytes.Buffer {
		return new(bytes.Buffer)
	})

	tabWriterPool = Instance(func() *tabwriter.Writer {
		return tabwriter.NewWriter(io.Discard, 0, 0, 1, ' ', 0)
	})
)

func getWriterPair() (*bytes.Buffer, *tabwriter.Writer) {
	buf := bufPool.Get()
	buf.Reset()

	tw := tabWriterPool.Get()
	tw.Init(buf, 0, 0, 1, ' ', 0)

	return buf, tw
}

func putWriterPair(buf *bytes.Buffer, tw *tabwriter.Writer) {
	_ = tw.Flush()
	tabWriterPool.Put(tw)
	buf.Reset()
	bufPool.Put(buf)
}

func shortenPath(path string) string {
	if cwd, err := os.Getwd(); err == nil {
		if rel, err := filepath.Rel(cwd, path); err == nil {
			return rel
		}
	}
	return path
}

// NewDevHandler initializes a DevHandler with optional HandlerOptions.
func NewDevHandler(out io.Writer, opts *slog.HandlerOptions) *DevHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	cw := &colorWriter{
		out: out,
	}

	return &DevHandler{
		base: slog.NewTextHandler(cw, opts),
		opts: opts,
		cw:   cw,
	}
}
