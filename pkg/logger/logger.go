package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type ColorLogs struct {
	w slog.Handler
}

func NewColorLogsHandler(w slog.Handler) *ColorLogs {
	return &ColorLogs{w: w}
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

	//source := record.
	//if source.File != "" {
	//	fmt.Printf("Source: %s:%d\n", source.File, source.Line)
	//}

	//func(groups []string, a slog.Attr) slog.Attr {
	//	if a.Key == slog.SourceKey {
	//		s := a.Value.Any().(*slog.Source)
	//		s.File = path.Base(s.File)
	//	}
	//	return a
	//}
	var attrs []slog.Attr
	record.Attrs(func(attr slog.Attr) bool {
		fmt.Println(attr)
		attrs = append(attrs, attr)
		return true
	})

	fmt.Println(attrs)

	timeStr := record.Time.Format("[2006-01-02 15:04:05]")
	keyWord := fmt.Sprintf("%s%s\033[0m", levelColor, record.Level.String()) // Цвет только для уровня

	msg := record.Message
	formatted := fmt.Sprintf("%s %s %s\n", timeStr, keyWord, msg)
	fmt.Fprint(os.Stdout, formatted)

	return nil
}

func (cl *ColorLogs) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ColorLogs{w: cl.w.WithAttrs(attrs)}
}

func (cl *ColorLogs) WithGroup(name string) slog.Handler {
	return &ColorLogs{w: cl.w.WithGroup(name)}
}
