package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	timeFormat   = "[2006-01-02 15:04:05]"
	reset        = "\033[0m"
	red          = 31
	green        = 32
	yellow       = 33
	cyan         = 36
	lightGray    = 37
	lightRed     = 91
	lightMagenta = 95
	lightBlue    = 94
	lightCyan    = 96
)

func colorizer(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

type Logger struct {
	w slog.Handler
	o *slog.HandlerOptions
}

func NewLoggerHandler(w slog.Handler, o *slog.HandlerOptions) *Logger {
	return &Logger{w: w, o: o}
}

func (cl *Logger) Enabled(ctx context.Context, levels slog.Level) bool {
	return cl.w.Enabled(ctx, levels)
}

func (cl *Logger) Handle(_ context.Context, record slog.Record) error {
	var keyWord string
	switch record.Level {
	case slog.LevelDebug:
		keyWord = colorizer(cyan, record.Level.String())
	case slog.LevelInfo:
		keyWord = colorizer(green, record.Level.String())
	case slog.LevelWarn:
		keyWord = colorizer(yellow, record.Level.String())
	case slog.LevelError:
		keyWord = colorizer(red, record.Level.String())
	}

	var attrs []string
	record.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, attr.String())
		return true
	})

	builder := strings.Builder{}
	for i, s := range attrs {
		builder.WriteString(colorizer(lightGray, s))
		if i != len(attrs)-1 {
			builder.WriteString("\n")
		}
	}

	timeStr := colorizer(yellow, record.Time.Format(timeFormat))
	src, line := addSource(cl.o.AddSource, runtime.CallersFrames([]uintptr{record.PC}))

	msg := record.Message
	formatted := fmt.Sprintf("%s %s %s %s %s \n%s\n", timeStr, keyWord, src, line, msg, builder.String())
	fmt.Fprint(os.Stdout, formatted)

	return nil
}

func (cl *Logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Logger{w: cl.w.WithAttrs(attrs)}
}

func (cl *Logger) WithGroup(name string) slog.Handler {
	return &Logger{w: cl.w.WithGroup(name)}
}

func addSource(source bool, frames *runtime.Frames) (string, string) {

	src := &slog.Source{
		Function: "",
		File:     "",
		Line:     0,
	}

	if source {
		frame, _ := frames.Next()
		if frame.File != "" {
			src = &slog.Source{
				Function: frame.Function,
				File:     frame.File,
				Line:     frame.Line,
			}
		}
		srcBuilder := strings.Builder{}
		test := strings.SplitAfter(src.File, "/")
		if len(test) > 0 {
			srcBuilder.WriteString(test[len(test)-1])
		}

		s := fmt.Sprintf("%s: %s", colorizer(lightBlue, "SRC"), colorizer(lightMagenta, srcBuilder.String()))
		line := fmt.Sprintf("%s: %s", colorizer(lightCyan, "| LINE"), colorizer(lightMagenta, strconv.Itoa(src.Line)))
		return s, line
	}
	return "", ""
}
