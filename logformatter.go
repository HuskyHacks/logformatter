package logformatter

import (
	"fmt"

	"github.com/fatih/color"
)

// LogType represents the type of log message.
type LogType int

const (
	// Warning log type
	Warning LogType = iota
	// Info log type
	Info
	// Success log type
	Success
	// Error log type
	Error
	// Warning, no color
	WarningNoColor
	// Info, no color
	InfoNoColor
	// Success, no color
	SuccessNoColor
	// Error, no color
	ErrorNoColor
)

var logTypeColors = map[LogType]*color.Color{
	Warning:        color.New(color.FgYellow),
	Info:           color.New(color.FgBlue),
	Success:        color.New(color.FgGreen),
	Error:          color.New(color.FgRed),
	WarningNoColor: color.New(color.Reset),
	InfoNoColor:    color.New(color.Reset),
	SuccessNoColor: color.New(color.Reset),
	ErrorNoColor:   color.New(color.Reset),
}

var logTypeGlyphs = map[LogType]string{
	Warning: "[!] ",
	Info:    "[*] ",
	Success: "[✔] ",
	Error:   "[✘] ",
}

// Log prints a formatted log message with color and glyph based on the log type.
func Log(logType LogType, message string) {
	color := logTypeColors[logType]
	glyph := logTypeGlyphs[logType]
	formattedMessage := fmt.Sprintf("%s%s", color.Sprint(glyph), color.Sprint(message))
	fmt.Println(formattedMessage)
}
