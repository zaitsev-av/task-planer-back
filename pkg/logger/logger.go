package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const (
	timeFormat   = "[2006-01-02 15:04:05]"
	reset        = "\033[0m"
	yellow       = 33
	lightRed     = 91
	lightMagenta = 95
)

func colorizer(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

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
	var attrs = make(map[string]string)
	record.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.String()
		return true
	})

	attrsAsBytes, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return fmt.Errorf("error when marshaling attrs: %w", err)
	}
	var attributes string

	outAttrs := strings.Builder{}
	outAttrs.WriteString(string(attrsAsBytes))

	timeStr := colorizer(yellow, record.Time.Format(timeFormat))
	keyWord := fmt.Sprintf("%s%s\033[0m", levelColor, record.Level.String())
	attributes = colorizer(lightMagenta, outAttrs.String())

	msg := record.Message
	formatted := fmt.Sprintf("%s %s %s\n %s\n", timeStr, keyWord, msg, attributes)
	fmt.Fprint(os.Stdout, formatted)

	return nil
}

func (cl *ColorLogs) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ColorLogs{w: cl.w.WithAttrs(attrs)}
}

func (cl *ColorLogs) WithGroup(name string) slog.Handler {
	return &ColorLogs{w: cl.w.WithGroup(name)}
}
