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
	yellow       = 33
	lightGray    = 37
	lightRed     = 91
	lightMagenta = 95
	lightBlue    = 94
	lightCyan    = 96
)

func colorizer(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

type ColorLogs struct {
	w         slog.Handler
	addSource bool
}

func NewColorLogsHandler(w slog.Handler, addSource bool) *ColorLogs {
	return &ColorLogs{w: w, addSource: addSource}
}

func (cl *ColorLogs) Enabled(ctx context.Context, levels slog.Level) bool {
	return cl.w.Enabled(ctx, levels)
}

func (cl *ColorLogs) Handle(ctx context.Context, record slog.Record) error {
	var levelColor string
	switch record.Level {
	case slog.LevelDebug:
		levelColor = "\033[36m" // Cyan
	case slog.LevelInfo:
		levelColor = "\033[32m" // Green
	case slog.LevelWarn:
		levelColor = "\033[33m" // Yellow
	case slog.LevelError:
		levelColor = "\033[31m" // Red
	default:
		levelColor = "\033[0m" // Reset
	}

	var attrs []string
	record.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, attr.String())
		return true
	})

	source := &slog.Source{
		Function: "",
		File:     "",
		Line:     0,
	}

	if cl.addSource {
		frames := runtime.CallersFrames([]uintptr{record.PC})
		frame, _ := frames.Next()
		if frame.File != "" {
			source = &slog.Source{
				Function: frame.Function,
				File:     frame.File,
				Line:     frame.Line,
			}
		}

	}

	builder := strings.Builder{}
	for i, s := range attrs {
		builder.WriteString(colorizer(lightGray, s))
		if i != len(attrs)-1 {
			builder.WriteString("\n")
		}
	}

	srcBuilder := strings.Builder{}
	test := strings.SplitAfter(source.File, "/")
	if len(test) > 0 {
		srcBuilder.WriteString(test[len(test)-1])
	}

	timeStr := colorizer(yellow, record.Time.Format(timeFormat))
	keyWord := fmt.Sprintf("%s%s\033[0m", levelColor, record.Level.String())
	src := fmt.Sprintf("%s: %s", colorizer(lightBlue, "SRC"), colorizer(lightMagenta, srcBuilder.String()))
	line := fmt.Sprintf("%s: %s", colorizer(lightCyan, "|LINE"), colorizer(lightMagenta, strconv.Itoa(source.Line)))

	msg := record.Message
	formatted := fmt.Sprintf("%s %s %s %s%s \n%s\n", timeStr, keyWord, msg, src, line, builder.String())
	fmt.Fprint(os.Stdout, formatted)

	return nil
}

func (cl *ColorLogs) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ColorLogs{w: cl.w.WithAttrs(attrs)}
}

func (cl *ColorLogs) WithGroup(name string) slog.Handler {
	return &ColorLogs{w: cl.w.WithGroup(name)}
}
