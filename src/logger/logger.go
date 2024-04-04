package logger

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var timestampStyle = lipgloss.NewStyle().
	Align(lipgloss.Right).
	Foreground(lipgloss.ANSIColor(7)).
	Faint(true).MarginRight(1)

var levelStyle = lipgloss.NewStyle().Width(5)

var statusStyle = lipgloss.NewStyle().Width(3).MarginLeft(1).MarginRight(1).Bold(true)

var msgStyle = lipgloss.NewStyle().
	Faint(true).
	PaddingLeft(1).
	Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(lipgloss.Color("8"))

var pathStyle = lipgloss.NewStyle().MarginLeft(1)

type Level int8

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 8
	LevelError Level = 16
)

type LoggerOpts struct {
	TimeFormat string
}

type Logger struct {
	w    io.Writer
	b    bytes.Buffer
	opts LoggerOpts
}

func NewLogger(writer io.Writer, opts LoggerOpts) *Logger {
	return &Logger{
		w:    writer,
		b:    *bytes.NewBufferString(""),
		opts: opts,
	}
}

type LogOpts struct {
	Level  Level
	Status uint16
	Method string
	Path   string
}

func (l *Logger) Log(msg string, opts LogOpts) {

	var (
		ts, level, status, styledMsg string
	)

	ts = timestampStyle.Render(time.Now().Format(l.opts.TimeFormat))

	if opts.Status <= 0 {
		status = statusStyle.Render("---")
	} else if opts.Status < 200 {
		status = statusStyle.Foreground(lipgloss.Color("8")).Render(fmt.Sprintf("%d", opts.Status))
	} else if opts.Status < 300 {
		status = statusStyle.Foreground(lipgloss.Color("2")).Render(fmt.Sprintf("%d", opts.Status))
	} else if opts.Status < 400 {
		status = statusStyle.Foreground(lipgloss.Color("5")).Render(fmt.Sprintf("%d", opts.Status))
	} else if opts.Status < 500 {
		status = statusStyle.Foreground(lipgloss.Color("3")).Render(fmt.Sprintf("%d", opts.Status))
	} else {
		status = statusStyle.Foreground(lipgloss.Color("1")).Render(fmt.Sprintf("%d", opts.Status))
	}

	if opts.Level < LevelInfo {
		level = levelStyle.Foreground(lipgloss.Color("4")).Render("DEBUG")
	} else if opts.Level < LevelWarn {
		level = levelStyle.Foreground(lipgloss.Color("8")).Render("INFO")
	} else if opts.Level < LevelError {
		level = levelStyle.Foreground(lipgloss.Color("3")).Render("WARN")
	} else {
		level = levelStyle.Foreground(lipgloss.Color("1")).Render("ERROR")
	}

	styledMsg = msgStyle.Render(msg)

	statusLine := lipgloss.JoinHorizontal(lipgloss.Top, level, status, opts.Method, pathStyle.Render(opts.Path))
	fullMsg := lipgloss.JoinVertical(lipgloss.Left, statusLine, styledMsg)
	if msg == "" {
		fullMsg = statusLine
	}

	l.b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, ts, fullMsg))
	l.b.WriteString("\n")

	l.b.WriteTo(l.w)
}
