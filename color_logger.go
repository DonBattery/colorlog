package colorlog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

type ColorLogger struct {
	Output    io.Writer
	conponent string
	debug     bool
}

func TurnOffColors() {
	color.NoColor = true
}

func TurnOnColors() {
	color.NoColor = false
}

// NewColorLogger creates a new color logger
func NewColorLogger(conponent string, debug bool) *ColorLogger {
	return &ColorLogger{
		Output:    os.Stdout,
		conponent: conponent,
		debug:     debug,
	}
}

func (c *ColorLogger) getDate() string {
	return time.Now().Format("15:04:05")
}

// Debug [DEBUG 17:47:02] Conponent: Debug message
// Does nothing if debug is false
func (c *ColorLogger) Debug(msg string) {
	if !c.debug {
		return
	}
	out := color.New(color.FgHiCyan, color.Bold).Sprint("DEBUG ", c.getDate())
	fmt.Fprintf(c.Output, "[%s] %s: %s\n", out, c.conponent, msg)
}

// Debugf formatted Debug message
func (c *ColorLogger) Debugf(format string, a ...interface{}) {
	c.Debug(fmt.Sprintf(format, a...))
}

// Info [INFO 17:47:02] Conponent: Info message
func (c *ColorLogger) Info(msg string) {
	out := color.New(color.FgHiBlue, color.Bold).Sprint("INFO  ", c.getDate())
	fmt.Fprintf(c.Output, "[%s] %s: %s\n", out, c.conponent, msg)
}

// Infof formatted Info message
func (c *ColorLogger) Infof(format string, a ...interface{}) {
	c.Info(fmt.Sprintf(format, a...))
}

// Warn [WARN 17:47:02] Conponent: Warning message
func (c *ColorLogger) Warn(msg string) {
	out := color.New(color.FgHiYellow, color.Bold).Sprint("WARN  ", c.getDate())
	fmt.Fprintf(c.Output, "[%s] %s: %s\n", out, c.conponent, msg)
}

// Warnf formatted Warn message
func (c *ColorLogger) Warnf(format string, a ...interface{}) {
	c.Warn(fmt.Sprintf(format, a...))
}

// Error [ERROR 17:47:02] Cononent: Error message
func (c *ColorLogger) Error(msg string) {
	out := color.New(color.FgHiRed, color.Bold).Sprint("ERROR ", c.getDate())
	fmt.Fprintf(c.Output, "[%s] %s: %s\n", out, c.conponent, msg)
}

// Errorf formatted Error message
func (c *ColorLogger) Errorf(format string, a ...interface{}) {
	c.Error(fmt.Sprintf(format, a...))
}

// Fatal [FATAL 17:18:02] Conponent: Fatal error message
// exits the program with status code 1
func (c *ColorLogger) Fatal(msg string) {
	out := color.New(color.FgHiYellow, color.BgRed, color.Bold).Sprint("FATAL ", c.getDate())
	fmt.Fprintf(c.Output, "[%s] %s: %s\n", out, c.conponent, msg)
	os.Exit(1)
}

// Fatalf formatted Fatal error message
func (c *ColorLogger) Fatalf(format string, a ...interface{}) {
	c.Fatal(fmt.Sprintf(format, a...))
}
