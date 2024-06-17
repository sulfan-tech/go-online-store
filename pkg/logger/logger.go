package logger

import (
	"fmt"
	"io"
	"time"
)

type LogLevel int

type Writer interface {
	WriteString(s string) (n int, err error)
}

type customWriter struct {
	writer io.Writer
}

func NewWriter(w io.Writer) Writer {
	return &customWriter{writer: w}
}

func (w *customWriter) WriteString(s string) (n int, err error) {
	return fmt.Fprint(w.writer, s)
}

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

type Logger struct {
	stdout Writer // Output to stdout (terminal)
	tag    string // Tag for log messages
}

func NewLogger(output io.Writer, tag string) *Logger {
	return &Logger{
		stdout: NewWriter(output),
		tag:    tag,
	}
}

func (l *Logger) WithTag(tag string) *Logger {
	return &Logger{
		stdout: l.stdout,
		tag:    tag,
	}
}

func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

func (l *Logger) log(level LogLevel, msg string) {
	logMsg := fmt.Sprintf("%s [%s] %s %s\n", time.Now().Format(time.RFC3339), levelToString(level), l.tag, msg)

	// Output using the custom writer
	if l.stdout != nil {
		l.stdout.WriteString(logMsg)
	}
}

func levelToString(level LogLevel) string {
	switch level {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
